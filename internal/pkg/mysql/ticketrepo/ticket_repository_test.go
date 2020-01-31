package ticketrepo

import (
	"context"
	"github.com/pepeunlimited/microservice-kit/validator"
	"github.com/pepeunlimited/users/internal/pkg/ent"
	"github.com/pepeunlimited/users/internal/pkg/mysql/userrepo"
	"testing"
	"time"
)

func TestTicketMySQL_CreateTicket(t *testing.T) {

	ctx := context.TODO()
	client := ent.NewEntClient()

	users := userrepo.NewUserRepository(client)
	tickets := NewTicketRepository(client)

	users.DeleteAll(ctx)

	user,_, err := users.CreateUser(ctx, "ssiimoo", "simo.alakotila@gmail.com", "p4sw0rd")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	ticket, err := tickets.CreateTicket(ctx, time.Now().Add(1 * time.Minute), user.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if validator.IsEmpty(ticket.Token) {
		t.FailNow()
	}
	ticket0, user0, err := tickets.GetTicketUserByToken(ctx, ticket.Token)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if ticket0 == nil {
		t.FailNow()
	}
	if user0 == nil {
		t.FailNow()
	}
}

func TestTicketMySQL_GetTicketUserByTokenExpired(t *testing.T) {
	ctx := context.TODO()
	client := ent.NewEntClient()

	users := userrepo.NewUserRepository(client)
	tickets := NewTicketRepository(client)

	users.DeleteAll(ctx)
	user,_, err := users.CreateUser(ctx, "ssiimoo", "simo.alakotila@gmail.com", "p4sw0rd")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	ticket, err := tickets.CreateTicket(ctx, time.Now().Add(1 * time.Second), user.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if validator.IsEmpty(ticket.Token) {
		t.FailNow()
	}
	time.Sleep(2 * time.Second)
	ticket0, user0, err := tickets.GetTicketUserByToken(ctx, ticket.Token)
	if err == nil {
		t.FailNow()
	}
	if err != ErrTicketExpired {
		t.FailNow()
	}
	if ticket0 != nil {
		t.FailNow()
	}
	if user0 != nil {
		t.FailNow()
	}
}

func TestTicketMySQL_UseTicket(t *testing.T) {
	ctx := context.TODO()
	client := ent.NewEntClient()
	users := userrepo.NewUserRepository(client)
	tickets := NewTicketRepository(client)
	users.DeleteAll(ctx)
	user,_,err := users.CreateUser(ctx, "ssiimoo", "simo.alakotila@gmail.com", "p4sw0rd")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	ticket, err := tickets.CreateTicket(ctx, time.Now().Add(1 * time.Minute), user.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if err := tickets.UseTicket(ctx, ticket.Token); err != nil {
		t.Error(err)
		t.FailNow()
	}
	if err := tickets.UseTicket(ctx, "a"); err == nil {
		t.FailNow()
	} else {
		if err.Error() != "ent: ticket not found" {
			t.FailNow()
		}
	}
}

func TestTicketMySQL_GetTicketUserByToken(t *testing.T) {
	ctx := context.TODO()
	client := ent.NewEntClient()
	users := userrepo.NewUserRepository(client)
	tickets := NewTicketRepository(client)
	users.DeleteAll(ctx)
	_, _, err := tickets.GetTicketUserByToken(ctx, "a")
	if err == nil {
		t.FailNow()
	}
	if err != ErrTicketNotExist {
		t.FailNow()
	}
}