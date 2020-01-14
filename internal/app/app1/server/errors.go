package server

import (
	"github.com/pepeunlimited/microservice-kit/rpcz"
	"github.com/pepeunlimited/users/internal/app/app1/ticketrepo"
	"github.com/pepeunlimited/users/internal/app/app1/userrepo"
	"github.com/pepeunlimited/users/rpcusers"
	"github.com/twitchtv/twirp"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func isUserError(err error) error {
	switch err {
	case userrepo.ErrUserNotExist:
		return twirp.NotFoundError("user not exist").WithMeta(rpcz.Reason, rpcusers.UserNotFound)
	case userrepo.ErrUserLocked:
		return twirp.NewError(twirp.PermissionDenied ,"user is locked").WithMeta(rpcz.Reason, rpcusers.UserIsLocked)
	case userrepo.ErrUserBanned:
		return twirp.NewError(twirp.PermissionDenied ,"user is banned").WithMeta(rpcz.Reason, rpcusers.UserIsBanned)
	}
	log.Print("user-service: unknown isUserError: "+err.Error())
	//unknown
	return twirp.NewError(twirp.Internal ,"unknown error: "+err.Error())
}

func isCryptoError(err error) error {
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return twirp.NewError(twirp.InvalidArgument, err.Error()).WithMeta(rpcz.Reason, rpcusers.InvalidCredentials)
	}
	return twirp.InternalError("user-service: unknown isCryptoError: "+err.Error())
}

func isTicketError(err error) error {
	switch err {
	case ticketrepo.ErrTicketNotExist:
		return twirp.NewError(twirp.NotFound, "ticket not found").WithMeta(rpcz.Reason, rpcusers.TicketNotFound)
	case ticketrepo.ErrTicketExpired:
		return twirp.NewError(twirp.InvalidArgument, "token expired").WithMeta(rpcz.Reason, rpcusers.TicketExpired)
	}
	log.Print("user-service: unknown isTicketError: "+err.Error())
	// unknown
	return twirp.InternalErrorWith(err)
}