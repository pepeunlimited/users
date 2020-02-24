package errorz

import (
	"github.com/pepeunlimited/users/internal/pkg/mysql/ticket"
	"github.com/pepeunlimited/users/pkg/rpc/user"
	"github.com/twitchtv/twirp"
	"log"
)

func Ticket(err error) error {
	switch err {
	case ticket.ErrTicketNotExist:
		return twirp.NewError(twirp.NotFound, user.TicketNotFound)
	case ticket.ErrTicketExpired:
		return twirp.NewError(twirp.InvalidArgument, user.TicketExpired)
	case ticket.ErrUserNotExist:
		return twirp.NotFoundError(user.UserNotFound)
	case ticket.ErrErrTicketExist:
		return twirp.NewError(twirp.Aborted, user.TicketExist)
	}
	log.Print("user-service: unknown isTicketError: "+err.Error())
	return twirp.InternalErrorWith(err)
}