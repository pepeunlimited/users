package server

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pepeunlimited/microservice-kit/rpcz"
	"github.com/pepeunlimited/users/internal/app/app1/ent"
	"github.com/pepeunlimited/users/internal/app/app1/repository"
	"github.com/pepeunlimited/users/internal/app/app1/validator"
	rpc2 "github.com/pepeunlimited/users/rpc"
	"github.com/twitchtv/twirp"
)

type UserServer struct {
	tickets repository.TicketRepository
	users repository.UserRepository
	validator validator.UserServerValidator
}

func (server UserServer) UpdatePassword(context.Context, *rpc2.UpdatePasswordParams) (*rpc2.UpdatePasswordResponse, error) {
	panic("implement me")
}

func (server UserServer) ForgotPassword(context.Context, *rpc2.ForgotPasswordParams) (*empty.Empty, error) {
	panic("implement me")
}

func (server UserServer) VerifyResetPassword(context.Context, *rpc2.VerifyPasswordParams) (*rpc2.VerifyPasswordResponse, error) {
	panic("implement me")
}

func (server UserServer) ResetPassword(context.Context, *rpc2.ResetPasswordParams) (*rpc2.ResetPasswordResponse, error) {
	panic("implement me")
}

func (server UserServer) CreateUser(ctx context.Context, params *rpc2.CreateUserParams) (*rpc2.User, error) {
	if err := server.validator.CreateUser(params); err != nil {
		return nil, err
	}
	user, err := server.users.CreateUser(ctx, params.Username, params.Email, params.Password, repository.User)
	if err  != nil {
		switch err {
		case repository.ErrUsernameExist:
			return nil, twirp.NewError(twirp.AlreadyExists, err.Error()).WithMeta(rpcz.Unique, "username")
		case repository.ErrEmailExist:
			return nil, twirp.NewError(twirp.AlreadyExists, err.Error()).WithMeta(rpcz.Unique, "email")
		}
		return nil, twirp.NewError(twirp.Aborted, err.Error())
	}
	return &rpc2.User{
		Id:                   int64(user.ID),
		Username:             user.Username,
		Email:                user.Email,
		Role:   			  user.Role,
	}, nil
}

func (server UserServer) GetUser(ctx context.Context, params *rpc2.GetUserParams) (*rpc2.User, error) {
	userId, err := rpcz.GetUserId(ctx)
	if err != nil {
		return nil, twirp.InternalError("can't access userId from ctx err: "+err.Error())
	}
	user, err := server.users.GetUserById(ctx, int(userId))
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, twirp.NotFoundError("user not exist").WithMeta(rpcz.NotFound, "user")
		}
		return nil, err
	}
	return &rpc2.User{
		Id:       int64(user.ID),
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
	}, nil
}

func NewUserServer(client *ent.Client) UserServer {
	return UserServer{
		users: repository.NewUserRepository(client),
		tickets: repository.NewTicketRepository(client),
		validator: validator.NewUserServerValidator(),
	}
}