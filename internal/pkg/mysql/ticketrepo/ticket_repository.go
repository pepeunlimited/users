package ticketrepo

import (
	"context"
	"errors"
	"github.com/pepeunlimited/microservice-kit/cryptoz"
	"github.com/pepeunlimited/users/internal/pkg/ent"
	"github.com/pepeunlimited/users/internal/pkg/ent/ticket"
	"time"
)


var (
	ErrTicketExpired 	= errors.New("tickets: ticket is expired")
	ErrTicketNotExist 	= errors.New("tickets: ticket not exist")
	ErrUserNotExist 	= errors.New("users: user not exist")
)

// Access the tickets table
// 	- many-to-one to users
type TicketRepository interface {
	CreateTicket(ctx context.Context, expiresAt time.Time, userId int) (*ent.Ticket, error)
	GetTicketUserByToken(ctx context.Context, token string) (*ent.Ticket, *ent.User, error)
	DeleteTickets(ctx context.Context)
	UseTicket(ctx context.Context, token string) error
}

type ticketMySQL struct {
	client *ent.Client
	crypto cryptoz.Crypto
}

func (repo ticketMySQL) isTicketValid(ticket *ent.Ticket) error {
	if ticket.ExpiresAt.Before(time.Now().UTC()) {
		return ErrTicketExpired
	}
	return nil
}

func (repo ticketMySQL) UseTicket(ctx context.Context, token string) error {
	_, err := repo.client.Ticket.Query().Where(ticket.Token(token)).Only(ctx)
	if err != nil {
		return err
	}

	_, err = repo.client.Ticket.Delete().Where(ticket.Token(token)).Exec(ctx)
	return err
}

func (repo ticketMySQL) GetTicketUserByToken(ctx context.Context, token string) (*ent.Ticket, *ent.User, error) {
	ticket, err := repo.client.Ticket.Query().Where(ticket.TokenEQ(token)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil, ErrTicketNotExist
		}
		return nil, nil, err
	}
	user, err := ticket.QueryUsers().Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil, ErrUserNotExist
		}
		return nil, nil, err
	}
	if err := repo.isTicketValid(ticket); err != nil {
		return nil, nil, err
	}
	return ticket, user, nil
}

func (repo ticketMySQL) CreateTicket(ctx context.Context, expiresAt time.Time, userId int) (*ent.Ticket, error) {
	token, err := repo.crypto.Random()
	if err != nil {
		return nil, err
	}
	ticket, err := repo.client.Ticket.Create().
		SetCreatedAt(time.Now().UTC()).
		SetExpiresAt(expiresAt.UTC()).
		SetToken(token).
		SetUsersID(userId).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (repo ticketMySQL) DeleteTickets(ctx context.Context) {
	repo.client.Ticket.Delete().ExecX(ctx)
}

func NewTicketRepository(client *ent.Client) TicketRepository {
	return ticketMySQL{client:client, crypto:cryptoz.NewCrypto()}
}

