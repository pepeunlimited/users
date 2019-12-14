package repository

import (
	"context"
	"github.com/pepeunlimited/microservice-kit/cryptoz"
	"github.com/pepeunlimited/users/internal/app/app1/ent"
	"testing"
	"time"
)

func TestUserMySQL_CreateUser(t *testing.T) {
	ctx := context.TODO()
	client := NewEntClient()
	repo := NewUserRepository(client)
	repo.DeleteAll(ctx)
	username := "ssiimoo"
	email := "simo.alakotila@gmail.com"
	password := "p4sw0rd"

	user, err := repo.CreateUser(ctx, username, email, password, Admin)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	userById, err := repo.GetUserById(ctx, user.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if userById.ID != user.ID {
		t.FailNow()
	}
}

func TestUserMySQL_GetUserByIdNotFound(t *testing.T) {
	ctx := context.TODO()
	client := NewEntClient()
	repo := NewUserRepository(client)
	repo.DeleteAll(ctx)
	user, err := repo.GetUserById(ctx, 100)
	if err == nil {
		t.Log(user)
		t.FailNow()
	}
	if !ent.IsNotFound(err) {
		t.FailNow()
	}
}

func TestUserMySQL_GetUsers(t *testing.T) {
	ctx := context.TODO()
	client := NewEntClient()
	repo := NewUserRepository(client)

	repo.CreateUser(ctx, "ssiimoo", "simo.alakotila@gmail.com", "ssiimoo", User)
	repo.CreateUser(ctx, "piiia", "piiiaaa@gmail.com", "ssiimoo", User)

	users, err := repo.GetUsers(ctx, 0, 20)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if len(users) != 2 {
		t.Log(len(users))
		t.FailNow()
	}
}

func TestUserMySQL_GetUserTicketsByUserId(t *testing.T) {
	ctx := context.TODO()
	client := NewEntClient()
	users := NewUserRepository(client)
	ticketsrepo := NewTicketRepository(client)
	users.DeleteAll(ctx)
	ssiimoo,_ := users.CreateUser(ctx, "ssiimoo", "simo.alakotila@gmail.com", "ssiimoo", User)
	piiia,_ := users.CreateUser(ctx, "piiia", "piiiaaa@gmail.com", "ssiimoo", User)
	_, tickets, err := users.GetUserTicketsByUserId(ctx, ssiimoo.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if len(tickets) != 0 {
		t.FailNow()
	}
	ticketsrepo.CreateTicket(ctx, time.Now().Add(1*time.Minute), ssiimoo.ID)
	ticketsrepo.CreateTicket(ctx, time.Now().Add(2*time.Minute), ssiimoo.ID)
	ticketsrepo.CreateTicket(ctx, time.Now().Add(3*time.Minute), ssiimoo.ID)
	_, tickets, err = users.GetUserTicketsByUserId(ctx, ssiimoo.ID)
	if len(tickets) != 3 {
		t.FailNow()
	}
	_, tickets, err = users.GetUserTicketsByUserId(ctx, piiia.ID)
	if len(tickets) != 0 {
		t.FailNow()
	}
}

func TestUserMySQL_UpdateUser(t *testing.T) {
	ctx := context.TODO()
	client := NewEntClient()
	users := NewUserRepository(client)
	users.DeleteAll(ctx)
	user,_ := users.CreateUser(ctx, "ssimoo", "simo.alakotila@gmail.com", "p4sw0rd", Admin)
	updated, err := users.UpdateUser(ctx, user.Update().SetUsername("ssimooo"))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if updated.Username != "ssimooo" {
		t.Log(updated.Username)
		t.FailNow()
	}
}

func TestUserMySQL_UpdatePasswordOk(t *testing.T) {
	ctx := context.TODO()
	client := NewEntClient()
	users := NewUserRepository(client)
	users.DeleteAll(ctx)
	password := "p4sw0rd"
	newpw := "new_"+password
	user, err := users.CreateUser(ctx, "ssimoo", "simo.alakotila@gmail.com", password, Admin)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	updated, err := users.UpdatePassword(ctx, user.ID, password, newpw)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	crypto := cryptoz.NewCrypto()
	if err := crypto.Check(updated.Password, newpw); err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestUserMySQL_UpdatePasswordFail(t *testing.T) {
	ctx := context.TODO()
	client := NewEntClient()
	users := NewUserRepository(client)
	users.DeleteAll(ctx)
	password := "p4sw0rd"
	newpw := "new_"+password
	user, err := users.CreateUser(ctx, "ssimoo", "simo.alakotila@gmail.com", password, Admin)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	updated, err := users.UpdatePassword(ctx, user.ID, password, newpw)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	crypto := cryptoz.NewCrypto()
	wrong := "wrong"
	err = crypto.Check(updated.Password, wrong)
	if err == nil {
		t.FailNow()
	}
	if err.Error() != "crypto/bcrypt: hashedPassword is not the hash of the given password" {
		t.FailNow()
	}

}

func TestUserMySQL_BanUserUnLock(t *testing.T) {
	ctx := context.TODO()
	client := NewEntClient()
	users := NewUserRepository(client)
	users.DeleteAll(ctx)
	user,_ := users.CreateUser(ctx, "simo", "simo.alakotila@gmail.com", "simo", Admin)
	err := users.BanUser(ctx, user.ID)
	if err != nil {
		t.FailNow()
	}
	_, err = users.UnbanUser(ctx, user.ID)
	if err != nil {
		t.FailNow()
	}
	users.LockUser(ctx, user.ID)
	if _, err := users.GetUserById(ctx, user.ID); err != nil {
		if err != ErrUserLocked {
			t.Log(err)
			t.FailNow()
		}
	}
	arrays, err := users.GetUsers(ctx, 0, 20)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if len(arrays) != 0 {
		t.FailNow()
	}
}

func TestUserMySQL_CreateUserEmailAndUsernameExist(t *testing.T) {
	ctx := context.TODO()
	client := NewEntClient()
	users := NewUserRepository(client)
	users.DeleteAll(ctx)
	users.CreateUser(ctx, "simo1", "simo.alakotila@gmail.com", "simo", Admin)
	users.CreateUser(ctx, "simo", "simo1.alakotila@gmail.com", "simo", Admin)
	_, err := users.CreateUser(ctx, "simo1", "a@a.com", "asd", Admin)
	if err == nil {
		t.FailNow()
	}
	if err != ErrUsernameExist {
		t.FailNow()
	}
	_, err = users.CreateUser(ctx, "simo2", "simo1.alakotila@gmail.com", "asd", Admin)
	if err == nil {
		t.FailNow()
	}
	if err != ErrEmailExist {
		t.FailNow()
	}
}
