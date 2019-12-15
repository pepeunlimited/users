package rpc

import (
	"github.com/pepeunlimited/microservice-kit/rpcz"
	"github.com/pepeunlimited/microservice-kit/validator"
	"github.com/twitchtv/twirp"
)

const (
	UserNotFound 	= "user_not_found"
	UserIsBanned 	= "user_is_banned"
	UserIsLocked	= "user_is_locked"
	Credentials     = "credentials"
	UsernameExist   = "username_exist"
	EmailExist      = "email_exist"
)

func IsReason(error twirp.Error, key string) bool {
	reason := error.Meta(rpcz.Reason)
	if validator.IsEmpty(reason) {
		return false
	}
	return reason == key
}