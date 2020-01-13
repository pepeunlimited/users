package rpcusers

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pepeunlimited/microservice-kit/errorz"
)

type Mock struct {
	Errors 		errorz.Stack
	IsAdmin 	bool
	User 		*User
}

func (u *Mock) SetProfilePicture(ctx context.Context,  params *SetProfilePictureParams) (*ProfilePicture, error) {
	if u.Errors.IsEmpty() {
		return &ProfilePicture{
			ProfilePictureId: params.ProfilePictureId,
		}, nil
	}
	return nil, u.Errors.Pop()
}

func (u *Mock) DeleteProfilePicture(context.Context, *DeleteProfilePictureParams) (*ProfilePicture, error) {
	if u.Errors.IsEmpty() {
		return &ProfilePicture{ProfilePictureId: 3}, nil
	}
	return nil, u.Errors.Pop()
}

func (u *Mock) CreateUser(context.Context, *CreateUserParams) (*User, error) {
	if u.Errors.IsEmpty() {
		return u.user(), nil
	}
	return nil, u.Errors.Pop()
}

func (u *Mock) UpdatePassword(context.Context, *UpdatePasswordParams) (*UpdatePasswordResponse, error) {
	panic("implement me")
}

func (u *Mock) ForgotPassword(context.Context, *ForgotPasswordParams) (*empty.Empty, error) {
	panic("implement me")
}

func (u *Mock) VerifyResetPassword(context.Context, *VerifyPasswordParams) (*VerifyPasswordResponse, error) {
	panic("implement me")
}

func (u *Mock) ResetPassword(context.Context, *ResetPasswordParams) (*ResetPasswordResponse, error) {
	panic("implement me")
}

func (u *Mock) GetUser(context.Context, *GetUserParams) (*User, error) {
	if u.Errors.IsEmpty() {
		if u.User == nil {
			return u.user(), nil
		}
		return u.User, nil
	}
	return nil, u.Errors.Pop()
}

func (u *Mock) VerifySignIn(context.Context, *VerifySignInParams) (*User, error) {
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