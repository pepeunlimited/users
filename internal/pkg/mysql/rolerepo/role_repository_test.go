package rolerepo

import (
	"context"
	"github.com/pepeunlimited/users/internal/pkg/ent"
	"github.com/pepeunlimited/users/internal/pkg/mysql/userrepo"
	"testing"
)

func TestRolesMySQL_AddRole(t *testing.T) {
	ctx := context.TODO()
	client := ent.NewEntClient()
	users := userrepo.NewUserRepository(client)
	users.DeleteAll(ctx)
	roles := NewRolesRepository(client)
	user,_, err := users.CreateUser(ctx, "simo", "simo@gmail.com", "p4sw0rd")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	err = roles.AddRole(ctx, user.ID, Admin)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	err = roles.AddRole(ctx, user.ID, Admin)
	if err == nil {
		t.FailNow()
	}
	if err != ErrDuplicateRole {
		t.Error(err)
		t.FailNow()
	}

	err = roles.DeleteRole(ctx, user.ID, Admin)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	err = roles.AddRole(ctx, user.ID, Admin)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}