package credentials

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pepeunlimited/microservice-kit/errorz"
	"github.com/pepeunlimited/users/pkg/rpc/user"
)

type CredentialsMock struct {
	Errors 		errorz.Stack
	IsAdmin 	bool
	User        *user.User
}

func (c *CredentialsMock) UpdatePassword(context.Context, *UpdatePasswordParams) (*empty.Empty, error) {
	panic("implement me")
}

func (c *CredentialsMock) ForgotPassword(context.Context, *ForgotPasswordParams) (*empty.Empty, error) {
	panic("implement me")
}

func (c *CredentialsMock) VerifyResetPassword(context.Context, *VerifyPasswordParams) (*empty.Empty, error) {
	panic("implement me")
}

func (c *CredentialsMock) ResetPassword(context.Context, *ResetPasswordParams) (*empty.Empty, error) {
	panic("implement me")
}

func (c *CredentialsMock) VerifySignIn(context.Context, *VerifySignInParams) (*VerifySignInResponse, error) {
	if c.Errors.IsEmpty() {
		if c.User == nil {
			return &VerifySignInResponse{
				Id:       c.user().Id,
				Username: c.user().Username,
				Email:    c.user().Email,
				Roles:    c.user().Roles,
			}, nil
		}
		return &VerifySignInResponse{
			Id:       c.User.Id,
			Username: c.User.Username,
			Email:    c.User.Email,
			Roles:    c.User.Roles,
		}, nil
	}
	return nil, c.Errors.Pop()
}

func NewCredentialsMock(errors []error, isAdmin bool) CredentialsService {
	return &CredentialsMock{Errors: errorz.NewErrorStack(errors), IsAdmin:isAdmin}
}

func (u *CredentialsMock) user() *user.User {
	roles := []string{"User"}
	if u.IsAdmin {
		roles = append(roles, "Admin")
	}
	return &user.User{
		Id:                   1,
		Username:             "kakkaliisa",
		Email:                "kakkaliisa@gmail.com",
		Roles:                roles,
	}
}


