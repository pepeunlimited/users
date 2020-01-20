package validator

import (
	"github.com/pepeunlimited/microservice-kit/validator"
	"github.com/pepeunlimited/users/rpccredentials"
	"github.com/pepeunlimited/users/rpcusers"
	"github.com/twitchtv/twirp"
)

type UserServerValidator struct {}


func NewUserServerValidator() UserServerValidator {
	return UserServerValidator{}
}

func (UserServerValidator) CreateUser(params *rpcusers.CreateUserParams) error {
	if validator.IsEmpty(params.Username) {
		return twirp.RequiredArgumentError("username")
	}
	if validator.IsEmpty(params.Email) {
		return twirp.RequiredArgumentError("email")
	}
	if validator.IsEmpty(params.Password) {
		return twirp.RequiredArgumentError("password")
	}
	return nil
}

func (UserServerValidator) SetProfilePicture(params *rpcusers.SetProfilePictureParams) error {
	if params.ProfilePictureId == 0 {
		return twirp.RequiredArgumentError("profile_picture_id")
	}
	return nil
}

func (UserServerValidator) GetUser(params *rpcusers.GetUserParams) error {
	if params.UserId == 0 {
		return twirp.RequiredArgumentError("user_id")
	}
	return nil
}

func (UserServerValidator) VerifySignIn(params *rpccredentials.VerifySignInParams) error {
	if validator.IsEmpty(params.Username) {
		return twirp.RequiredArgumentError("username")
	}
 	if validator.IsEmpty(params.Password) {
 		return twirp.RequiredArgumentError("password")
	}
	return nil
}

func (UserServerValidator) ValidForgotPassword(params *rpccredentials.ForgotPasswordParams) error {
	if params.Username == nil && params.Email == nil {
		return twirp.RequiredArgumentError("username_or_email")
	}

	if params.Email != nil && validator.IsEmpty(params.Email.Value) || params.Username != nil && validator.IsEmpty(params.Username.Value)  {
		return twirp.RequiredArgumentError("username_or_email")
	}

	if validator.IsEmpty(params.Language) {
		return twirp.RequiredArgumentError("language")
	}
	return nil
}

func (UserServerValidator) VerifyResetPassword(params *rpccredentials.VerifyPasswordParams) error {
	if validator.IsEmpty(params.TicketToken) {
		return twirp.RequiredArgumentError("ticket_token")
	}
	return nil
}

func (UserServerValidator) ResetPassword(params *rpccredentials.ResetPasswordParams) error {
	if validator.IsEmpty(params.TicketToken) {
		return twirp.RequiredArgumentError("ticket_token")
	}
	if validator.IsEmpty(params.Password) {
		return twirp.RequiredArgumentError("password")
	}
	return nil
}

func (v UserServerValidator) UpdatePassword(params *rpccredentials.UpdatePasswordParams) error {
	if params.UserId == 0 {
		return twirp.RequiredArgumentError("user_id")
	}
	if validator.IsEmpty(params.NewPassword) {
		return twirp.RequiredArgumentError("new_password")
	}
	if validator.IsEmpty(params.CurrentPassword) {
		return twirp.RequiredArgumentError("current_password")
	}
	return nil
}