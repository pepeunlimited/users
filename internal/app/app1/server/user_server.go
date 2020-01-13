package server

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/pepeunlimited/authorization-twirp/rpcauthorization"
	"github.com/pepeunlimited/files/rpcspaces"
	"github.com/pepeunlimited/microservice-kit/cryptoz"
	"github.com/pepeunlimited/microservice-kit/mail"
	"github.com/pepeunlimited/microservice-kit/rpcz"
	"github.com/pepeunlimited/users/internal/app/app1/ent"
	"github.com/pepeunlimited/users/internal/app/app1/ticketrepo"
	"github.com/pepeunlimited/users/internal/app/app1/userrepo"
	"github.com/pepeunlimited/users/internal/app/app1/validator"
	"github.com/pepeunlimited/users/rpcusers"
	"github.com/twitchtv/twirp"
	"golang.org/x/crypto/bcrypt"
	"log"
	"path"
	"time"
)

type UserServer struct {
	tickets       ticketrepo.TicketRepository
	users         userrepo.UserRepository
	crypto        cryptoz.Crypto
	validator     validator.UserServerValidator
	authorization rpcauthorization.AuthorizationService
	smtpUsername  string
	smtpPassword  string
	smtpProvider  mail.Provider
	spaces 		  rpcspaces.SpacesService
}

func (server UserServer) SetProfilePicture(ctx context.Context, params *rpcusers.SetProfilePictureParams) (*rpcusers.ProfilePicture, error) {
	if err := server.validator.SetProfilePicture(params); err != nil {
		return nil, err
	}
	user, err := rpcauthorization.IsSignedIn(ctx, server.authorization)
	if err != nil {
		return nil, err
	}
	// validate access to the file
	file, err := server.spaces.GetFile(ctx, &rpcspaces.GetFileParams{
		FileId: &wrappers.Int64Value{
			Value: params.ProfilePictureId,
		},
	})
	if err != nil {
		return nil, err
	}

	if file.UserId != user.UserId {
		return nil, twirp.InvalidArgumentError("profile_picture_id", "can't access other uploader fileID").WithMeta(rpcz.Reason, rpcusers.ProfilePictureAccessDenied);
	}

	err = server.users.SetProfilePictureID(ctx, int(user.UserId), params.ProfilePictureId)
	if err != nil {
		return nil, server.isUserError(err)
	}

	return &rpcusers.ProfilePicture{ProfilePictureId: params.ProfilePictureId}, nil
}

func (server UserServer) DeleteProfilePicture(ctx context.Context, params *rpcusers.DeleteProfilePictureParams) (*rpcusers.ProfilePicture, error) {
	user, err := rpcauthorization.IsSignedIn(ctx, server.authorization)
	if err != nil {
		return nil, err
	}
	fromDB, err := server.users.GetUserById(ctx, int(user.UserId))
	if err != nil {
		return nil, server.isUserError(err)
	}
	if fromDB.ProfilePictureID == nil {
		return &rpcusers.ProfilePicture{}, nil
	}
	if err := server.users.DeleteProfilePictureID(ctx, int(user.UserId)); err != nil {
		return nil, server.isUserError(err)
	}
	profilePictureId := *fromDB.ProfilePictureID
	return &rpcusers.ProfilePicture{ProfilePictureId: profilePictureId}, nil
}

func (server UserServer) VerifySignIn(ctx context.Context, params *rpcusers.VerifySignInParams) (*rpcusers.User, error) {
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

	return &rpcusers.User{
		Id:               int64(user.ID),
		Username:         user.Username,
		Email:            user.Email,
		Roles:            rolesToString(roles),
		ProfilePictureId:  userId,
	}, nil
}

func (server UserServer) UpdatePassword(ctx context.Context, params *rpcusers.UpdatePasswordParams) (*rpcusers.UpdatePasswordResponse, error) {
	verified, err := rpcauthorization.IsSignedIn(ctx, server.authorization)
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
	return &rpcusers.UpdatePasswordResponse{}, nil
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

func (server UserServer) ForgotPassword(ctx context.Context, params *rpcusers.ForgotPasswordParams) (*empty.Empty, error) {
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
			return nil, twirp.NewError(twirp.AlreadyExists, "ticket already exist").WithMeta(rpcz.Reason, rpcusers.TicketExist)
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

func (server UserServer) VerifyResetPassword(ctx context.Context, params *rpcusers.VerifyPasswordParams) (*rpcusers.VerifyPasswordResponse, error) {
	if err := server.validator.VerifyResetPassword(params); err != nil {
		return nil, err
	}
	_, _, err := server.tickets.GetTicketUserByToken(ctx, params.TicketToken)
	if err != nil {
		return nil, server.isTicketError(err)
	}
	return &rpcusers.VerifyPasswordResponse{}, nil
}

func (server UserServer) ResetPassword(ctx context.Context, params *rpcusers.ResetPasswordParams) (*rpcusers.ResetPasswordResponse, error) {
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
	return &rpcusers.ResetPasswordResponse{}, nil
}

func (server UserServer) CreateUser(ctx context.Context, params *rpcusers.CreateUserParams) (*rpcusers.User, error) {
	if err := server.validator.CreateUser(params); err != nil {
		return nil, err
	}
	user, role, err := server.users.CreateUser(ctx, params.Username, params.Email, params.Password)
	if err  != nil {
		switch err {
		case userrepo.ErrUsernameExist:
			return nil, twirp.NewError(twirp.AlreadyExists, err.Error()).WithMeta(rpcz.Reason, rpcusers.UsernameExist)
		case userrepo.ErrEmailExist:
			return nil, twirp.NewError(twirp.AlreadyExists, err.Error()).WithMeta(rpcz.Reason, rpcusers.EmailExist)
		}
		return nil, twirp.NewError(twirp.Aborted, err.Error())
	}
	return &rpcusers.User{
		Id:                   int64(user.ID),
		Username:             user.Username,
		Email:                user.Email,
		Roles:   			  []string{role.Role},
	}, nil
}

func (server UserServer) isUserError(err error) error {
	switch err {
	case userrepo.ErrUserNotExist:
		return twirp.NotFoundError("user not exist").WithMeta(rpcz.Reason, rpcusers.UserNotFound)
	case userrepo.ErrUserLocked:
		return twirp.NewError(twirp.PermissionDenied ,"user is locked").WithMeta(rpcz.Reason, rpcusers.UserIsLocked)
	case userrepo.ErrUserBanned:
		return twirp.NewError(twirp.PermissionDenied ,"user is banned").WithMeta(rpcz.Reason, rpcusers.UserIsBanned)
	}
	log.Print("user-service: unknown isUserError: "+err.Error())
	//unknown
	return twirp.NewError(twirp.Internal ,"unknown error: "+err.Error())
}

func (server UserServer) isCryptoError(err error) error {
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return twirp.NewError(twirp.InvalidArgument, err.Error()).WithMeta(rpcz.Reason, rpcusers.InvalidCredentials)
	}
	return twirp.InternalError("user-service: unknown isCryptoError: "+err.Error())
}

func (server UserServer) isTicketError(err error) error {
	switch err {
	case ticketrepo.ErrTicketNotExist:
		return twirp.NewError(twirp.NotFound, "ticket not found").WithMeta(rpcz.Reason, rpcusers.TicketNotFound)
	case ticketrepo.ErrTicketExpired:
		return twirp.NewError(twirp.InvalidArgument, "token expired").WithMeta(rpcz.Reason, rpcusers.TicketExpired)
	}
	log.Print("user-service: unknown isTicketError: "+err.Error())
	// unknown
	return twirp.InternalErrorWith(err)
}

func (server UserServer) GetUser(ctx context.Context, params *rpcusers.GetUserParams) (*rpcusers.User, error) {
	resp, err := rpcauthorization.IsSignedIn(ctx, server.authorization)
	if err != nil {
		return nil, err
	}
	user, err := server.users.GetUserByEmail(ctx, resp.Email)
	if err != nil {
		return nil, server.isUserError(err)
	}

	userId := &wrappers.Int64Value{}
	if user.ProfilePictureID != nil {
		userId.Value = *user.ProfilePictureID
	}

	return &rpcusers.User{
		Id:               resp.UserId,
		Username:         resp.Username,
		Email:            resp.Email,
		Roles:            resp.Roles,
		ProfilePictureId:  userId,
	}, nil
}

func NewUserServer(client *ent.Client, authorization rpcauthorization.AuthorizationService, smtpUsername string, smtpPassword string, provider mail.Provider, spaces rpcspaces.SpacesService) UserServer {
	return UserServer{
		users:         userrepo.NewUserRepository(client),
		tickets:       ticketrepo.NewTicketRepository(client),
		crypto:        cryptoz.NewCrypto(),
		validator:     validator.NewUserServerValidator(),
		authorization: authorization,
		smtpPassword:  smtpPassword,
		smtpUsername:  smtpUsername,
		smtpProvider:  provider,
		spaces: 	   spaces,
	}
}