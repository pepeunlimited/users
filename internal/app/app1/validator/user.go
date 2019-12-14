package validator

import (
	"github.com/pepeunlimited/microservice-kit/validator"
	rpc2 "github.com/pepeunlimited/users/rpc"
	"github.com/twitchtv/twirp"
)

type UserServerValidator struct {}


func NewUserServerValidator() UserServerValidator {
	return UserServerValidator{}
}

func (UserServerValidator) CreateUser(params *rpc2.CreateUserParams) error {
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

func (UserServerValidator) GetUser(params *rpc2.GetUserParams) error {
	return nil
}