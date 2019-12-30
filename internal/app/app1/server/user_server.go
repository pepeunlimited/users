package server

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	rpc2 "github.com/pepeunlimited/authorization-twirp/rpc"
	"github.com/pepeunlimited/microservice-kit/cryptoz"
	"github.com/pepeunlimited/microservice-kit/mail"
	"github.com/pepeunlimited/microservice-kit/rpcz"
	"github.com/pepeunlimited/users/internal/app/app1/ent"
	"github.com/pepeunlimited/users/internal/app/app1/repository"
	"github.com/pepeunlimited/users/internal/app/app1/validator"
	"github.com/pepeunlimited/users/rpc"
	"github.com/twitchtv/twirp"
	"golang.org/x/crypto/bcrypt"
	"log"
	"path"
	"time"
)

type UserServer struct {
	tickets repository.TicketRepository
	users repository.UserRepository
	crypto  cryptoz.Crypto
	validator validator.UserServerValidator
	authService rpc2.AuthorizationService


	smtpUsername 		string
	smtpPassword 		string
	smtpProvider 		mail.Provider
}

func (server UserServer) VerifySignIn(ctx context.Context, params *rpc.VerifySignInParams) (*rpc.User, error) {
	err := server.validator.VerifySignIn(params)
	if err != nil {
		return nil, err
	}
	user, roles, err := server.users.GetUserRolesByUsername(ctx, params.Username)
	if err != nil {
		return nil, server.isUserError(err)
	}
	if err := server.crypto.Check(user.Password, params.Password); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return nil, twirp.NewError(twirp.Unauthenticated, err.Error()).WithMeta(rpcz.Reason, rpc.Credentials)
		}
		return nil, twirp.InternalError("unknown error during sign-in err: "+err.Error())
	}
	return &rpc.User{
		Id:       int64(user.ID),
		Username: user.Username,
		Email:    user.Email,
		Roles:    rolesToString(roles),
	}, nil
}

func (server UserServer) UpdatePassword(context.Context, *rpc.UpdatePasswordParams) (*rpc.UpdatePasswordResponse, error) {
	panic("implement me")
}

func (server UserServer) ForgotPassword(ctx context.Context, params *rpc.ForgotPasswordParams) (*empty.Empty, error) {
	if err :=  server.validator.ValidForgotPassword(params); err != nil {
		return nil, err
	}
	user, err := server.users.GetUserByUsername(ctx, params.Username)
	if err != nil {
		if err == repository.ErrUserNotExist {
			return nil, twirp.NewError(twirp.NotFound, "user not exist").WithMeta(rpcz.Reason, rpc.UserNotFound)
		} else {
			log.Print("users-service: unknown error during get user on forgot password: "+err.Error())
			return nil, twirp.InternalErrorWith(err)
		}
	}
	ticket, err := server.tickets.CreateTicket(ctx, time.Now().UTC().Add(1*time.Hour), user.ID)
	if err != nil {
		if ent.IsConstraintFailure(err) {
			return nil, twirp.NewError(twirp.Aborted, "ticket token already exist").WithMeta(rpcz.Reason, rpc.TicketTokenExist)
		}
		log.Print("users-service: unknown error during create ticket on forgot password: "+err.Error())
		return nil, twirp.InternalErrorWith(err)
	}

	baseURL := mail.CreateBaseURL("http://api.dev.pepeunlimited.com")
	baseURL.Path = path.Join(baseURL.Path, "reset_password")
	baseURL.Path = path.Join(baseURL.Path, ticket.Token)

	client := mail.NewBuilder(
		server.smtpUsername,
		server.smtpPassword).
		From(mail.PepeUnlimited, "Pepe Unlimited Oy").
		To([]string{user.Email}).
		Subject("Reset Password").
		Content(baseURL.String()).
		Build(server.smtpProvider)

	err = client.Send()
	if err != nil {
		log.Print("users-service: unknown error during mail send on forgot password: "+err.Error())
		server.tickets.UseTicket(ctx, ticket.Token)
		return nil, twirp.NewError(twirp.Aborted, "failed to send mail for user").WithMeta(rpcz.Reason, mail.MailSendFailed)
	}
	return &empty.Empty{}, nil
}

func (server UserServer) VerifyResetPassword(context.Context, *rpc.VerifyPasswordParams) (*rpc.VerifyPasswordResponse, error) {
	panic("implement me")
}

func (server UserServer) ResetPassword(context.Context, *rpc.ResetPasswordParams) (*rpc.ResetPasswordResponse, error) {
	panic("implement me")
}

func (server UserServer) CreateUser(ctx context.Context, params *rpc.CreateUserParams) (*rpc.User, error) {
	if err := server.validator.CreateUser(params); err != nil {
		return nil, err
	}
	user, role, err := server.users.CreateUser(ctx, params.Username, params.Email, params.Password)
	if err  != nil {
		switch err {
		case repository.ErrUsernameExist:
			return nil, twirp.NewError(twirp.AlreadyExists, err.Error()).WithMeta(rpcz.Reason, rpc.UsernameExist)
		case repository.ErrEmailExist:
			return nil, twirp.NewError(twirp.AlreadyExists, err.Error()).WithMeta(rpcz.Reason, rpc.EmailExist)
		}
		return nil, twirp.NewError(twirp.Aborted, err.Error())
	}
	return &rpc.User{
		Id:                   int64(user.ID),
		Username:             user.Username,
		Email:                user.Email,
		Roles:   			  []string{role.Role},
	}, nil
}

func (server UserServer) isUserError(err error) error {
	switch err {
	case repository.ErrUserNotExist:
		return twirp.NotFoundError("user not exist").WithMeta(rpcz.Reason, rpc.UserNotFound)
	case repository.ErrUserLocked:
		return twirp.NewError(twirp.PermissionDenied ,"user is locked").WithMeta(rpcz.Reason, rpc.UserIsLocked)
	case repository.ErrUserBanned:
		return twirp.NewError(twirp.PermissionDenied ,"user is banned").WithMeta(rpcz.Reason, rpc.UserIsBanned)
	}
	//unknown
	return twirp.NewError(twirp.Internal ,"unknown error: "+err.Error())
}

func (server UserServer) GetUser(ctx context.Context, params *rpc.GetUserParams) (*rpc.User, error) {
	token, err := rpcz.GetAuthorizationWithoutPrefix(ctx)
	if err != nil {
		return nil, twirp.RequiredArgumentError("Authorization")
	}
	// verify the token from the authorization service: blacklist and expired..
	resp, err := server.authService.Verify(ctx, &rpc2.VerifyParams{Token:token})
	if err != nil {
		return nil, err
	}
	return &rpc.User{
		Id:       resp.UserId,
		Username: resp.Username,
		Email:    resp.Email,
		Roles:    resp.Roles,
	}, nil
}

func NewUserServer(client *ent.Client, authService rpc2.AuthorizationService, smtpUsername string, smtpPassword string, provider mail.Provider) UserServer {
	return UserServer{
		users: 			repository.NewUserRepository(client),
		tickets: 		repository.NewTicketRepository(client),
		crypto:			cryptoz.NewCrypto(),
		validator: 		validator.NewUserServerValidator(),
		authService: 	authService,
		smtpPassword:	smtpPassword,
		smtpUsername:	smtpUsername,
		smtpProvider:	provider,
	}
}