package twirp

import (
	"context"
	"github.com/pepeunlimited/microservice-kit/mail"
	"github.com/pepeunlimited/users/internal/pkg/ent"
	"github.com/pepeunlimited/users/internal/pkg/mysql/ticket"
	userrepo "github.com/pepeunlimited/users/internal/pkg/mysql/user"
	"github.com/pepeunlimited/users/internal/server/errorz"
	"github.com/pepeunlimited/users/internal/server/validator"
	"github.com/pepeunlimited/users/pkg/rpc/user"
)

type UserServer struct {
	tickets      ticket.TicketRepository
	users        userrepo.UserRepository
	validator    validator.UserServerValidator
	smtpUsername string
	smtpPassword string
	smtpProvider mail.Provider
}

func (server UserServer) SetProfilePicture(ctx context.Context, params *user.SetProfilePictureParams) (*user.ProfilePicture, error) {
	if err := server.validator.SetProfilePicture(params); err != nil {
		return nil, err
	}
	_, err := server.users.GetUserById(ctx, int(params.UserId))
	if err != nil {
		return nil, errorz.User(err)
	}
	err = server.users.SetProfilePictureID(ctx, int(params.UserId), params.ProfilePictureId)
	if err != nil {
		return nil, errorz.User(err)
	}
	return &user.ProfilePicture{ProfilePictureId: params.ProfilePictureId}, nil
}

func (server UserServer) DeleteProfilePicture(ctx context.Context, params *user.DeleteProfilePictureParams) (*user.ProfilePicture, error) {
	err := server.validator.DeleteProfilePicture(params)
	if err != nil {
		return nil, err
	}
	fromDB, err := server.users.GetUserById(ctx, int(params.UserId))
	if err != nil {
		return nil, errorz.User(err)
	}
	if fromDB.ProfilePictureID == nil {
		return &user.ProfilePicture{}, nil
	}
	if err := server.users.DeleteProfilePictureID(ctx, int(params.UserId)); err != nil {
		return nil, errorz.User(err)
	}
	profilePictureId := *fromDB.ProfilePictureID
	return &user.ProfilePicture{ProfilePictureId: profilePictureId}, nil
}

func (server UserServer) CreateUser(ctx context.Context, params *user.CreateUserParams) (*user.User, error) {
	if err := server.validator.CreateUser(params); err != nil {
		return nil, err
	}
	fromDB, role, err := server.users.CreateUser(ctx, params.Username, params.Email, params.Password)
	if err  != nil {
		return nil, errorz.User(err)
	}
	return ToUser(fromDB, []*ent.Role{role}), nil
}

func (server UserServer) GetUser(ctx context.Context, params *user.GetUserParams) (*user.User, error) {
	err := server.validator.GetUser(params)
	if err != nil {
		return nil, err
	}
	user, roles, err := server.users.GetUserRolesByUserId(ctx, int(params.UserId))
	if err != nil {
		return nil, errorz.User(err)
	}
	return ToUser(user, roles), nil
}

func NewUserServer(client *ent.Client,
	smtpUsername string,
	smtpPassword string,
	provider mail.Provider) UserServer {
	return UserServer{
		users:        userrepo.New(client),
		tickets:      ticket.New(client),
		validator:    validator.NewUserServerValidator(),
		smtpPassword: smtpPassword,
		smtpUsername: smtpUsername,
		smtpProvider: provider,
	}
}