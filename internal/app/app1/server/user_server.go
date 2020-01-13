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
		return nil, isUserError(err)
	}

	return &rpcusers.ProfilePicture{ProfilePictureId: params.ProfilePictureId}, nil
}

func (server UserServer) DeleteProfilePicture(ctx context.Context, params *empty.Empty) (*rpcusers.ProfilePicture, error) {
	user, err := rpcauthorization.IsSignedIn(ctx, server.authorization)
	if err != nil {
		return nil, err
	}
	fromDB, err := server.users.GetUserById(ctx, int(user.UserId))
	if err != nil {
		return nil, isUserError(err)
	}
	if fromDB.ProfilePictureID == nil {
		return &rpcusers.ProfilePicture{}, nil
	}
	if err := server.users.DeleteProfilePictureID(ctx, int(user.UserId)); err != nil {
		return nil, isUserError(err)
	}
	profilePictureId := *fromDB.ProfilePictureID
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

func (server UserServer) GetUser(ctx context.Context, empty *empty.Empty) (*rpcusers.User, error) {
	resp, err := rpcauthorization.IsSignedIn(ctx, server.authorization)
	if err != nil {
		return nil, err
	}
	user, err := server.users.GetUserByEmail(ctx, resp.Email)
	if err != nil {
		return nil, isUserError(err)
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

func NewUserServer(client *ent.Client,
	authorization rpcauthorization.AuthorizationService,
	smtpUsername string,
	smtpPassword string,
	provider mail.Provider,
	spaces rpcspaces.SpacesService) UserServer {
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