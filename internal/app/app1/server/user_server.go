package server

import (
	"context"
	"github.com/pepeunlimited/microservice-kit/mail"
	"github.com/pepeunlimited/microservice-kit/rpcz"
	"github.com/pepeunlimited/users/internal/app/app1/ent"
	"github.com/pepeunlimited/users/internal/app/app1/ticketrepo"
	"github.com/pepeunlimited/users/internal/app/app1/userrepo"
	"github.com/pepeunlimited/users/internal/app/app1/validator"
	"github.com/pepeunlimited/users/rpcusers"
	"github.com/twitchtv/twirp"
)

type UserServer struct {
	tickets       ticketrepo.TicketRepository
	users         userrepo.UserRepository
	validator     validator.UserServerValidator
	smtpUsername  string
	smtpPassword  string
	smtpProvider  mail.Provider
}

func (server UserServer) SetProfilePicture(ctx context.Context, params *rpcusers.SetProfilePictureParams) (*rpcusers.ProfilePicture, error) {
	if err := server.validator.SetProfilePicture(params); err != nil {
		return nil, err
	}
	_, err := server.users.GetUserById(ctx, int(params.UserId))
	if err != nil {
		return nil, isUserError(err)
	}
	err = server.users.SetProfilePictureID(ctx, int(params.UserId), params.ProfilePictureId)
	if err != nil {
		return nil, isUserError(err)
	}
	return &rpcusers.ProfilePicture{ProfilePictureId: params.ProfilePictureId}, nil
}

func (server UserServer) DeleteProfilePicture(ctx context.Context, params *rpcusers.DeleteProfilePictureParams) (*rpcusers.ProfilePicture, error) {
	err := server.validator.DeleteProfilePicture(params)
	if err != nil {
		return nil, err
	}
	user, err := server.users.GetUserById(ctx, int(params.UserId))
	if err != nil {
		return nil, isUserError(err)
	}
	if user.ProfilePictureID == nil {
		return &rpcusers.ProfilePicture{}, nil
	}
	if err := server.users.DeleteProfilePictureID(ctx, int(params.UserId)); err != nil {
		return nil, isUserError(err)
	}
	profilePictureId := *user.ProfilePictureID
	return &rpcusers.ProfilePicture{ProfilePictureId: profilePictureId}, nil
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

func (server UserServer) GetUser(ctx context.Context, params *rpcusers.GetUserParams) (*rpcusers.User, error) {
	err := server.validator.GetUser(params)
	if err != nil {
		return nil, err
	}
	user, roles, err := server.users.GetUserRolesByUserId(ctx, int(params.UserId))
	if err != nil {
		return nil, isUserError(err)
	}
	ppID := int64(0)
	if user.ProfilePictureID != nil {
		ppID = *user.ProfilePictureID
	}
	return &rpcusers.User{
		Id:               int64(user.ID),
		Username:         user.Username,
		Email:            user.Email,
		Roles:            rolesToString(roles),
		ProfilePictureId:  ppID,
	}, nil
}

func NewUserServer(client *ent.Client,
	smtpUsername string,
	smtpPassword string,
	provider mail.Provider) UserServer {
	return UserServer{
		users:         userrepo.NewUserRepository(client),
		tickets:       ticketrepo.NewTicketRepository(client),
		validator:     validator.NewUserServerValidator(),
		smtpPassword:  smtpPassword,
		smtpUsername:  smtpUsername,
		smtpProvider:  provider,
	}
}