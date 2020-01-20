package rpcusers

import (
	"context"
	"github.com/pepeunlimited/microservice-kit/errorz"
)

func (u *UsersMock) SetProfilePicture(ctx context.Context,  params *SetProfilePictureParams) (*ProfilePicture, error) {
	if u.Errors.IsEmpty() {
		return &ProfilePicture{
			ProfilePictureId: params.ProfilePictureId,
		}, nil
	}
	return nil, u.Errors.Pop()
}

func (u *UsersMock) CreateUser(context.Context, *CreateUserParams) (*User, error) {
	if u.Errors.IsEmpty() {
		return u.user(), nil
	}
	return nil, u.Errors.Pop()
}

func (u *UsersMock) GetUser(context.Context, *GetUserParams) (*User, error) {
	if u.Errors.IsEmpty() {
		return u.user(), nil
	}
	return nil, u.Errors.Pop()
}

func (u *UsersMock) DeleteProfilePicture(context.Context, *DeleteProfilePictureParams) (*ProfilePicture, error) {
	panic("implement me")
}

func NewUserServiceMock(errors []error, isAdmin bool) UserService {
	return &UsersMock{Errors: errorz.NewErrorStack(errors), IsAdmin:isAdmin}
}

func (u *UsersMock) user() *User {
	roles := []string{"User"}
	if u.IsAdmin {
		roles = append(roles, "Admin")
	}
	return &User{
		Id:                   1,
		Username:             "kakkaliisa",
		Email:                "kakkaliisa@gmail.com",
		Roles:                roles,
	}
}

type UsersMock struct {
	Errors 		errorz.Stack
	IsAdmin 	bool
	User        *User
}



