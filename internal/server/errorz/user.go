package errorz

import (
	userrepo "github.com/pepeunlimited/users/internal/pkg/mysql/user"
	"github.com/pepeunlimited/users/pkg/rpc/user"
	"github.com/twitchtv/twirp"
	"log"
)

func User(err error) error {
	switch err {
	case userrepo.ErrUserNotExist:
		return twirp.NotFoundError(user.UserNotFound)
	case userrepo.ErrUserLocked:
		return twirp.NewError(twirp.PermissionDenied , user.UserIsLocked)
	case userrepo.ErrUserBanned:
		return twirp.NewError(twirp.PermissionDenied , user.UserIsBanned)
	case userrepo.ErrUsernameExist:
		return twirp.NewError(twirp.AlreadyExists, user.UsernameExist)
	case userrepo.ErrEmailExist:
		return twirp.NewError(twirp.AlreadyExists, user.EmailExist)
	}
	log.Print("user-service: unknown isUserError: "+err.Error())
	return twirp.InternalErrorWith(err)
}