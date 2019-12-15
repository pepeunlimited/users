package validator

import (
	"github.com/pepeunlimited/microservice-kit/validator"
	"github.com/pepeunlimited/users/rpc"
	"github.com/twitchtv/twirp"
)

type UserServerValidator struct {}


func NewUserServerValidator() UserServerValidator {
	return UserServerValidator{}
}

func (UserServerValidator) CreateUser(params *rpc.CreateUserParams) error {
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

func (UserServerValidator) GetUser(params *rpc.GetUserParams) error {
	return nil
}

func (UserServerValidator) SignIn(params *rpc.VerifySignInParams) error {
	if validator.IsEmpty(params.Username) {
		return twirp.RequiredArgumentError("username")
	}
 	if validator.IsEmpty(params.Password) {
 		return twirp.RequiredArgumentError("password")
	}
	return nil
}