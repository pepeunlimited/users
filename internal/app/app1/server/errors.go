package server

import (
	"github.com/pepeunlimited/users/internal/app/app1/ticketrepo"
	"github.com/pepeunlimited/users/internal/app/app1/userrepo"
	"github.com/pepeunlimited/users/usersrpc"
	"github.com/twitchtv/twirp"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func isUserError(err error) error {
	switch err {
	case userrepo.ErrUserNotExist:
		return twirp.NotFoundError(usersrpc.UserNotFound)
	case userrepo.ErrUserLocked:
		return twirp.NewError(twirp.PermissionDenied ,usersrpc.UserIsLocked)
	case userrepo.ErrUserBanned:
		return twirp.NewError(twirp.PermissionDenied ,usersrpc.UserIsBanned)
	}
	log.Print("user-service: unknown isUserError: "+err.Error())
	return twirp.InternalErrorWith(err)
}

func isCryptoError(err error) error {
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return twirp.NewError(twirp.InvalidArgument, usersrpc.InvalidCredentials)
	}
	log.Print("user-service: unknown isCryptoError: "+err.Error())
	return twirp.InternalErrorWith(err)
}

func isTicketError(err error) error {
	switch err {
	case ticketrepo.ErrTicketNotExist:
		return twirp.NewError(twirp.NotFound, usersrpc.TicketNotFound)
	case ticketrepo.ErrTicketExpired:
		return twirp.NewError(twirp.InvalidArgument, usersrpc.TicketExpired)
	}
	log.Print("user-service: unknown isTicketError: "+err.Error())
	return twirp.InternalErrorWith(err)
}