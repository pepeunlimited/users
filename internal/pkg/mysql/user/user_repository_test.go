package user

import (
	"context"
	"github.com/pepeunlimited/microservice-kit/cryptoz"
	"github.com/pepeunlimited/users/internal/pkg/ent"
	"github.com/pepeunlimited/users/internal/pkg/mysql/role"
	"github.com/pepeunlimited/users/internal/pkg/mysql/ticket"
	"testing"
	"time"
)

func TestUserMySQL_ResetPassword(t *testing.T) {
	ctx := context.TODO()
	client := ent.NewEntClient()
	repo := New(client)
	repo.DeleteAll(ctx)
	username := "ssiimoo"
	email := "simo.alakotila@gmail.com"
	password := "p4sw0rd"
	user,_, err := repo.CreateUser(ctx, username, email, password)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	reseted, err := repo.ResetPassword(ctx, user.ID, "newpw")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if err := cryptoz.NewCrypto().Check(reseted.Password, "newpw"); err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestUserMySQL_CreateUser(t *testing.T) {
	ctx := context.TODO()
	client := ent.NewEntClient()
	repo := New(client)
	repo.DeleteAll(ctx)
	username := "ssiimoo"
	email := "simo.alakotila@gmail.com"
	password := "p4sw0rd"

	user, role, err := repo.CreateUser(ctx, username, email, password)
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
	if role.Role(role.Role) != role.User {
		t.FailNow()
	}
}

func TestUserMySQL_GetUserByIdNotFound(t *testing.T) {
	ctx := context.TODO()
	client := ent.NewEntClient()
	repo := New(client)
	repo.DeleteAll(ctx)
	user, err := repo.GetUserById(ctx, 100)
	if err == nil {
		t.Log(user)
		t.FailNow()
	}
	if err != ErrUserNotExist {
		t.FailNow()
	}
}

func TestUserMySQL_GetUsers(t *testing.T) {
	ctx := context.TODO()
	client := ent.NewEntClient()
	repo := New(client)

	repo.CreateUser(ctx, "ssiimoo", "simo.alakotila@gmail.com", "ssiimoo")
	repo.CreateUser(ctx, "piiia", "piiiaaa@gmail.com", "ssiimoo")

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
	client := ent.NewEntClient()
	users := New(client)
	ticketsrepo := ticket.New(client)
	users.DeleteAll(ctx)
	ssiimoo,_, _ := users.CreateUser(ctx, "ssiimoo", "simo.alakotila@gmail.com", "ssiimoo")
	piiia,_, _ := users.CreateUser(ctx, "piiia", "piiiaaa@gmail.com", "ssiimoo")
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
	client := ent.NewEntClient()
	users := New(client)
	users.DeleteAll(ctx)
	user,_,_ := users.CreateUser(ctx, "ssimoo", "simo.alakotila@gmail.com", "p4sw0rd")
	selected,_ := users.GetUserById(ctx, user.ID)
	updated, err := users.UpdateUser(ctx, selected.Update().SetUsername("ssimooo"))
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
	client := ent.NewEntClient()
	users := New(client)
	users.DeleteAll(ctx)
	password := "p4sw0rd"
	newpw := "new_"+password
	user,_, err := users.CreateUser(ctx, "ssimoo", "simo.alakotila@gmail.com", password)
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
	client := ent.NewEntClient()
	users := New(client)
	users.DeleteAll(ctx)
	password := "p4sw0rd"
	newpw := "new_"+password
	user,_, err := users.CreateUser(ctx, "ssimoo", "simo.alakotila@gmail.com", password)
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
	client := ent.NewEntClient()
	users := New(client)
	users.DeleteAll(ctx)
	user,_,_ := users.CreateUser(ctx, "simo", "simo.alakotila@gmail.com", "simo")
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
	client := ent.NewEntClient()
	users := New(client)
	users.DeleteAll(ctx)
	users.CreateUser(ctx, "simo1", "simo.alakotila@gmail.com", "simo")
	users.CreateUser(ctx, "simo", "simo1.alakotila@gmail.com", "simo")
	_,_, err := users.CreateUser(ctx, "simo1", "a@a.com", "asd")
	if err == nil {
		t.FailNow()
	}
	if err != ErrUsernameExist {
		t.FailNow()
	}
	_,_, err = users.CreateUser(ctx, "simo2", "simo1.alakotila@gmail.com", "asd")
	if err == nil {
		t.FailNow()
	}
	if err != ErrEmailExist {
		t.FailNow()
	}
}

func TestUserMySQL_GetUserRolesByUsername(t *testing.T) {
	ctx := context.TODO()
	client := ent.NewEntClient()
	users := New(client)
	users.DeleteAll(ctx)
	repo := role.New(client)
	user, _,_ := users.CreateUser(ctx, "simo", "simo.alakotila@gmail.com", "p4sw0rd")
	user, roles, err := users.GetUserRolesByUserId(ctx, user.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if len(roles) != 1 {
		t.FailNow()
	}
	err = repo.AddRole(ctx, user.ID, role.Admin)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	user, roles, err = users.GetUserRolesByUserId(ctx, user.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if len(roles) != 2 {
		t.FailNow()
	}
}

func TestUserMySQL_GetUserRolesByUserId(t *testing.T) {
	ctx := context.TODO()
	client := ent.NewEntClient()
	users := New(client)
	users.DeleteAll(ctx)
	repo := role.New(client)
	username := "simo"
	user, _,_ := users.CreateUser(ctx, username, "simo.alakotila@gmail.com", "p4sw0rd")
	user, roles, err := users.GetUserRolesByUsername(ctx, username)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if len(roles) != 1 {
		t.FailNow()
	}
	err = repo.AddRole(ctx, user.ID, role.Admin)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	user, roles, err = users.GetUserRolesByUserId(ctx, user.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if len(roles) != 2 {
		t.FailNow()
	}
}

func TestUserMySQL_SetProfilePictureIDAndDelete(t *testing.T) {
	ctx := context.TODO()
	client := ent.NewEntClient()
	users := New(client)
	users.DeleteAll(ctx)
	username := "simo0"
	users.DeleteAll(ctx)
	user, _,err := users.CreateUser(ctx, username, "sim0.alakotila@gmail.com", "p4sw0rd")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if err := users.SetProfilePictureID(ctx, user.ID, 3); err != nil {
		t.Error(err)
		t.FailNow()
	}
	afterUpdate, err := users.GetUserById(ctx, user.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if afterUpdate.ProfilePictureID == nil {
		t.FailNow()
	}
	if *afterUpdate.ProfilePictureID != 3 {
		t.FailNow()
	}
	if err := users.DeleteProfilePictureID(ctx, user.ID); err != nil {
		t.Error(err)
		t.FailNow()
	}
	afterUpdate2, err := users.GetUserById(ctx, user.ID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if afterUpdate2.ProfilePictureID != nil {
		t.FailNow()
	}
}