package twirp

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/pepeunlimited/microservice-kit/cryptoz"
	"github.com/pepeunlimited/microservice-kit/mail"
	"github.com/pepeunlimited/users/internal/pkg/ent"
	"github.com/pepeunlimited/users/internal/pkg/mysql/ticket"
	userrepo "github.com/pepeunlimited/users/internal/pkg/mysql/user"
	"github.com/pepeunlimited/users/internal/server/errorz"
	"github.com/pepeunlimited/users/internal/server/validator"
	"github.com/pepeunlimited/users/pkg/rpc/credentials"
	"github.com/twitchtv/twirp"
	"log"
	"path"
	"time"
)

type CredentialsServer struct {
	tickets      ticket.TicketRepository
	users        userrepo.UserRepository
	crypto       cryptoz.Crypto
	validator    validator.UserServerValidator
	smtpUsername string
	smtpPassword string
	smtpProvider mail.Provider
}

func (server CredentialsServer) VerifySignIn(ctx context.Context, params *credentials.VerifySignInParams) (*credentials.VerifySignInResponse, error) {
	err := server.validator.VerifySignIn(params)
	if err != nil {
		return nil, err
	}
	user, roles, err := server.users.GetUserRolesByUsername(ctx, params.Username)
	if err != nil {
		return nil, errorz.User(err)
	}
	if err := server.crypto.Check(user.Password, params.Password); err != nil {
		return nil, errorz.Crypto(err)
	}

	userId := &wrappers.Int64Value{}
	if user.ProfilePictureID != nil {
		userId.Value = *user.ProfilePictureID
	}
	return &credentials.VerifySignInResponse{
		Id:       int64(user.ID),
		Username: user.Username,
		Email:    user.Email,
		Roles:    rolesToString(roles),
	}, nil
}

func (server CredentialsServer) UpdatePassword(ctx context.Context, params *credentials.UpdatePasswordParams) (*empty.Empty, error) {
	err := server.validator.UpdatePassword(params)
	if err != nil {
		return nil, err
	}
	user,_, err := server.users.GetUserRolesByUserId(ctx, int(params.UserId))
	if err != nil {
		return nil, errorz.User(err)
	}
	if err := server.crypto.Check(user.Password, params.CurrentPassword); err != nil {
		return nil, errorz.Crypto(err)
	}
	_, err = server.users.UpdatePassword(ctx, int(params.UserId), params.CurrentPassword, params.NewPassword)
	if err != nil {
		return nil, errorz.User(err)
	}
	return &empty.Empty{}, nil
}

func (server CredentialsServer) findUserByUsernameOrEmail(ctx context.Context, username *wrappers.StringValue, email *wrappers.StringValue) (*ent.User, error) {
	if email == nil  {
		user, err := server.users.GetUserByUsername(ctx, username.Value)
		if err != nil {
			return nil, errorz.User(err)
		}
		return user, nil
	}
	user, err := server.users.GetUserByEmail(ctx, email.Value)
	if err != nil {
		return nil, errorz.User(err)
	}
	return user, nil
}

func (server CredentialsServer) ForgotPassword(ctx context.Context, params *credentials.ForgotPasswordParams) (*empty.Empty, error) {
	if err :=  server.validator.ValidForgotPassword(params); err != nil {
		return nil, err
	}

	fromDB, err := server.findUserByUsernameOrEmail(ctx, params.Username, params.Email)
	if err != nil {
		return nil, err
	}

	ticket, err := server.tickets.CreateTicket(ctx, time.Now().UTC().Add(1*time.Hour), fromDB.ID)
	if err != nil {
		return nil, errorz.Ticket(err)
	}

	baseURL := mail.CreateBaseURL("http://api.dev.pepeunlimited.com")
	baseURL.Path = path.Join(baseURL.Path, "reset_password")
	baseURL.Path = path.Join(baseURL.Path, ticket.Token)

	client := mail.NewBuilder(
		server.smtpUsername,
		server.smtpPassword).
		From(mail.PepeUnlimited, "Pepe Unlimited Oy").
		To([]string{fromDB.Email}).
		Subject("Reset Password").
		Content(baseURL.String()).
		Build(server.smtpProvider)

	err = client.Send()
	if err != nil {
		log.Print("users-service: unknown error during mail send on forgot password: "+err.Error())
		server.tickets.UseTicket(ctx, ticket.Token)
		return nil, twirp.NewError(twirp.Aborted, mail.MailSendFailed)
	}
	return &empty.Empty{}, nil
}

func (server CredentialsServer) VerifyResetPassword(ctx context.Context, params *credentials.VerifyPasswordParams) (*empty.Empty, error) {
	if err := server.validator.VerifyResetPassword(params); err != nil {
		return nil, err
	}
	_, _, err := server.tickets.GetTicketUserByToken(ctx, params.TicketToken)
	if err != nil {
		return nil, errorz.Ticket(err)
	}
	return &empty.Empty{}, nil
}

func (server CredentialsServer) ResetPassword(ctx context.Context, params *credentials.ResetPasswordParams) (*empty.Empty, error) {
	if err := server.validator.ResetPassword(params); err != nil {
		return nil, err
	}
	ticket,user, err := server.tickets.GetTicketUserByToken(ctx, params.TicketToken)
	if err != nil {
		return nil, errorz.Ticket(err)
	}
	_, err = server.users.ResetPassword(ctx, user.ID, params.Password)
	if err != nil {
		return nil, errorz.User(err)
	}
	server.tickets.UseTicket(ctx, ticket.Token)
	return &empty.Empty{}, nil
}

func NewCredentialsServer(client *ent.Client,
	smtpUsername string,
	smtpPassword string,
	provider mail.Provider) CredentialsServer {
	return CredentialsServer{
		users:        userrepo.New(client),
		tickets:      ticket.New(client),
		crypto:       cryptoz.NewCrypto(),
		validator:    validator.NewUserServerValidator(),
		smtpPassword: smtpPassword,
		smtpUsername: smtpUsername,
		smtpProvider: provider,
	}
}