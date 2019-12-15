package server

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pepeunlimited/microservice-kit/cryptoz"
	"github.com/pepeunlimited/microservice-kit/rpcz"
	"github.com/pepeunlimited/users/internal/app/app1/ent"
	"github.com/pepeunlimited/users/internal/app/app1/repository"
	"github.com/pepeunlimited/users/internal/app/app1/validator"
	"github.com/pepeunlimited/users/rpc"
	"github.com/twitchtv/twirp"
	"golang.org/x/crypto/bcrypt"
)

type UserServer struct {
	tickets repository.TicketRepository
	users repository.UserRepository
	crypto  cryptoz.Crypto
	validator validator.UserServerValidator
}

func (server UserServer) VerifySignIn(ctx context.Context, params *rpc.VerifySignInParams) (*rpc.User, error) {
	err := server.validator.SignIn(params)
	if err != nil {
		return nil, err
	}
	user, err := server.users.GetUserByUsername(ctx, params.Username)
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
		Role:     user.Role,
	}, nil
}

func (server UserServer) UpdatePassword(context.Context, *rpc.UpdatePasswordParams) (*rpc.UpdatePasswordResponse, error) {
	panic("implement me")
}

func (server UserServer) ForgotPassword(context.Context, *rpc.ForgotPasswordParams) (*empty.Empty, error) {
	panic("implement me")
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
	user, err := server.users.CreateUser(ctx, params.Username, params.Email, params.Password, repository.User)
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
		Role:   			  user.Role,
	}, nil
}

func (server UserServer) isUserError(err error) error {
	if err == repository.ErrUserNotExist {
		return twirp.NotFoundError("user not exist").WithMeta(rpcz.Reason, rpc.UserNotFound)
	}
	if err == repository.ErrUserLocked {
		return twirp.NewError(twirp.PermissionDenied ,"user is locked").WithMeta(rpcz.Reason, rpc.UserIsLocked)
	}
	if err == repository.ErrUserBanned {
		return twirp.NewError(twirp.PermissionDenied ,"user is banned").WithMeta(rpcz.Reason, rpc.UserIsBanned)
	}
	// unknown
	return twirp.NewError(twirp.Internal ,"unknown error: "+err.Error())
}

func (server UserServer) GetUser(ctx context.Context, params *rpc.GetUserParams) (*rpc.User, error) {
	userId, err := rpcz.GetUserId(ctx)
	if err != nil {
		return nil, twirp.InternalError("can't access userId from ctx err: "+err.Error())
	}
	user, err := server.users.GetUserById(ctx, int(userId))
	if err != nil {
		return nil, server.isUserError(err)
	}
	return &rpc.User{
		Id:       int64(user.ID),
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
	}, nil
}




func NewUserServer(client *ent.Client) UserServer {
	return UserServer{
		users: 		repository.NewUserRepository(client),
		tickets: 	repository.NewTicketRepository(client),
		crypto:		cryptoz.NewCrypto(),
		validator: 	validator.NewUserServerValidator(),
	}
}