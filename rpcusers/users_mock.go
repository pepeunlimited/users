package rpcusers

import (
	"context"
	"github.com/pepeunlimited/microservice-kit/errorz"
)

func (u *CredentialsMock) SetProfilePicture(ctx context.Context,  params *SetProfilePictureParams) (*ProfilePicture, error) {
	if u.Errors.IsEmpty() {
		return &ProfilePicture{
			ProfilePictureId: params.ProfilePictureId,
		}, nil
	}
	return nil, u.Errors.Pop()
}

func (u *CredentialsMock) DeleteProfilePicture(context.Context, *DeleteProfilePictureParams) (*ProfilePicture, error) {
	if u.Errors.IsEmpty() {
		return &ProfilePicture{ProfilePictureId: 3}, nil
	}
	return nil, u.Errors.Pop()
}

func (u *CredentialsMock) CreateUser(context.Context, *CreateUserParams) (*User, error) {
	if u.Errors.IsEmpty() {
		return u.user(), nil
	}
	return nil, u.Errors.Pop()
}

func (u *CredentialsMock) GetUser(context.Context, *GetUserParams) (*User, error) {
	if u.Errors.IsEmpty() {
		if u.User == nil {
			return u.user(), nil
		}
		return u.User, nil
	}
	return nil, u.Errors.Pop()
}

func NewUserServiceMock(errors []error, isAdmin bool) UserService {
	return &Mock{Errors: errorz.NewErrorStack(errors), IsAdmin:isAdmin}
}

func (u *Mock) user() *User {
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

}


