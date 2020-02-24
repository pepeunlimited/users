package errorz

import (
	"github.com/pepeunlimited/users/pkg/rpc/user"
	"github.com/twitchtv/twirp"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func Crypto(err error) error {
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return twirp.NewError(twirp.InvalidArgument, user.InvalidCredentials)
	}
	log.Print("user-service: unknown isCryptoError: "+err.Error())
	return twirp.InternalErrorWith(err)
}