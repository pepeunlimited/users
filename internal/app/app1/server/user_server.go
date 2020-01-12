package server

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
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
	tickets 			repository.TicketRepository
	users 				repository.UserRepository
	crypto  			cryptoz.Crypto
	validator 			validator.UserServerValidator
	authService 		rpc2.AuthorizationService
	smtpUsername 		string
	smtpPassword 		string
	smtpProvider 		mail.Provider
}

func (server UserServer) SetProfilePicture(ctx context.Context, params *rpc.SetProfilePictureParams) (*rpc.ProfilePicture, error) {
	if err := server.validator.SetProfilePicture(params); err != nil {
		return nil, err
	}

	token, err := rpcz.GetAuthorizationWithoutPrefix(ctx)
	if err != nil {
		return nil, twirp.RequiredArgumentError("authorization")
	}
	// verify the token from the authorization service: blacklist and expired..
	user, err := server.authService.VerifyAccessToken(ctx, &rpc2.VerifyAccessTokenParams{AccessToken:token})
	if err != nil {
		return nil, err
	}

	err = server.users.SetProfilePictureID(ctx, int(user.UserId), params.ProfilePictureId)
	if err != nil {
		return nil, server.isUserError(err)
	}

	return &rpc.ProfilePicture{ProfilePictureId: params.ProfilePictureId}, nil
}

func (server UserServer) DeleteProfilePicture(ctx context.Context, params *rpc.DeleteProfilePictureParams) (*rpc.ProfilePicture, error) {
	token, err := rpcz.GetAuthorizationWithoutPrefix(ctx)
	if err != nil {
		return nil, twirp.RequiredArgumentError("authorization")
	}
	// verify the token from the authorization service: blacklist and expired..
	user, err := server.authService.VerifyAccessToken(ctx, &rpc2.VerifyAccessTokenParams{AccessToken:token})
	if err != nil {
		return nil, err
	}
	fromDB, err := server.users.GetUserById(ctx, int(user.UserId))
	if err != nil {
		return nil, server.isUserError(err)
	}
	if fromDB.ProfilePictureID == nil {
		return &rpc.ProfilePicture{}, nil
	}
	if err := server.users.DeleteProfilePictureID(ctx, int(user.UserId)); err != nil {
		return nil, server.isUserError(err)
	}
	profilePictureId := *fromDB.ProfilePictureID
	return &rpc.ProfilePicture{ProfilePictureId: profilePictureId}, nil
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
		return nil, server.isCryptoError(err)
	}

	userId := &wrappers.Int64Value{}
	if user.ProfilePictureID != nil {
		userId.Value = *user.ProfilePictureID
	}

	return &rpc.User{
		Id:               int64(user.ID),
		Username:         user.Username,
		Email:            user.Email,
		Roles:            rolesToString(roles),
		ProfilePictureId:  userId,
	}, nil
}

func (server UserServer) UpdatePassword(ctx context.Context, params *rpc.UpdatePasswordParams) (*rpc.UpdatePasswordResponse, error) {
	token, err := rpcz.GetAuthorizationWithoutPrefix(ctx)
	if err != nil {
		return nil, twirp.RequiredArgumentError("authorization")
	}
	// verify the token from the authorization service: blacklist and expired..
	verified, err := server.authService.VerifyAccessToken(ctx, &rpc2.VerifyAccessTokenParams{AccessToken:token})
	if err != nil {
		return nil, err
	}
	user,_, err := server.users.GetUserRolesByUsername(ctx, verified.Username)
	if err != nil {
		return nil, server.isUserError(err)
	}
	if err := server.crypto.Check(user.Password, params.CurrentPassword); err != nil {
		return nil, server.isCryptoError(err)
	}
	_, err = server.users.UpdatePassword(ctx, int(verified.UserId), params.CurrentPassword, params.NewPassword)
	if err != nil {
		return nil, server.isUserError(err)
	}
	return &rpc.UpdatePasswordResponse{}, nil
}


func (server UserServer) findUserByUsernameOrEmail(ctx context.Context, username *wrappers.StringValue, email *wrappers.StringValue) (*ent.User, error) {
	if email == nil  {
		user, err := server.users.GetUserByUsername(ctx, username.Value)
		if err != nil {
			return nil, server.isUserError(err)
		}
		return user, nil
	}
	user, err := server.users.GetUserByEmail(ctx, email.Value)
	if err != nil {
		return nil, server.isUserError(err)
	}
	return user, nil
}

func (server UserServer) ForgotPassword(ctx context.Context, params *rpc.ForgotPasswordParams) (*empty.Empty, error) {
	if err :=  server.validator.ValidForgotPassword(params); err != nil {
		return nil, err
	}

	user, err := server.findUserByUsernameOrEmail(ctx, params.Username, params.Email)
	if err != nil {
		return nil, err
	}

	ticket, err := server.tickets.CreateTicket(ctx, time.Now().UTC().Add(1*time.Hour), user.ID)
	if err != nil {
		if ent.IsConstraintError(err) {
			return nil, twirp.NewError(twirp.AlreadyExists, "ticket already exist").WithMeta(rpcz.Reason, rpc.TicketExist)
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

func (server UserServer) VerifyResetPassword(ctx context.Context, params *rpc.VerifyPasswordParams) (*rpc.VerifyPasswordResponse, error) {
	if err := server.validator.VerifyResetPassword(params); err != nil {
		return nil, err
	}
	_, _, err := server.tickets.GetTicketUserByToken(ctx, params.TicketToken)
	if err != nil {
		return nil, server.isTicketError(err)
	}
	return &rpc.VerifyPasswordResponse{}, nil
}

func (server UserServer) ResetPassword(ctx context.Context, params *rpc.ResetPasswordParams) (*rpc.ResetPasswordResponse, error) {
	if err := server.validator.ResetPassword(params); err != nil {
		return nil, err
	}
	ticket,user, err := server.tickets.GetTicketUserByToken(ctx, params.TicketToken)
	if err != nil {
		return nil, server.isTicketError(err)
	}
	_, err = server.users.ResetPassword(ctx, user.ID, params.Password)
	if err != nil {
		return nil, server.isUserError(err)
	}
	server.tickets.UseTicket(ctx, ticket.Token)
	return &rpc.ResetPasswordResponse{}, nil
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
	log.Print("user-service: unknown isUserError: "+err.Error())
	//unknown
	return twirp.NewError(twirp.Internal ,"unknown error: "+err.Error())
}

func (server UserServer) isCryptoError(err error) error {
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return twirp.NewError(twirp.InvalidArgument, err.Error()).WithMeta(rpcz.Reason, rpc.InvalidCredentials)
	}
	return twirp.InternalError("user-service: unknown isCryptoError: "+err.Error())
}

func (server UserServer) isTicketError(err error) error {
	switch err {
	case repository.ErrTicketNotExist:
		return twirp.NewError(twirp.NotFound, "ticket not found").WithMeta(rpcz.Reason, rpc.TicketNotFound)
	case repository.ErrTicketExpired:
		return twirp.NewError(twirp.InvalidArgument, "token expired").WithMeta(rpcz.Reason, rpc.TicketExpired)
	}
	log.Print("user-service: unknown isTicketError: "+err.Error())
	// unknown
	return twirp.InternalErrorWith(err)
}

func (server UserServer) GetUser(ctx context.Context, params *rpc.GetUserParams) (*rpc.User, error) {
	token, err := rpcz.GetAuthorizationWithoutPrefix(ctx)
	if err != nil {
		return nil, twirp.RequiredArgumentError("authorization")
	}
	// verify the token from the authorization service: blacklist and expired..
	resp, err := server.authService.VerifyAccessToken(ctx, &rpc2.VerifyAccessTokenParams{AccessToken:token})
	if err != nil {
		return nil, err
	}
	user, err := server.users.GetUserByEmail(ctx, resp.Email)
	if err != nil {
		return nil, err
	}

	userId := &wrappers.Int64Value{}
	if user.ProfilePictureID != nil {
		userId.Value = *user.ProfilePictureID
	}

	return &rpc.User{
		Id:               resp.UserId,
		Username:         resp.Username,
		Email:            resp.Email,
		Roles:            resp.Roles,
		ProfilePictureId:  userId,
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