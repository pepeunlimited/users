package user

import (
	"context"
	"github.com/pepeunlimited/users/internal/pkg/mysql/role"
	"github.com/pepeunlimited/users/internal/pkg/mysql/user"
	"log"
	"testing"
)

func TestUserServiceMock_GetUser(t *testing.T) {
	mock := NewUserServiceMock(nil, false)
	user, err := mock.GetUser(context.TODO(), &GetUserParams{})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	log.Print(user)
}

func TestUserServiceMock_GetUserError(t *testing.T) {
	mock := NewUserServiceMock([]error{role.ErrDuplicateRole, user.ErrUserNotExist}, true)
	_, err := mock.GetUser(context.TODO(), &GetUserParams{})
	if err == nil {
		t.FailNow()
	}
	_, err = mock.GetUser(context.TODO(), &GetUserParams{})
	if err == nil {
		t.FailNow()
	}
	user, err := mock.GetUser(context.TODO(), &GetUserParams{})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if user == nil {
		t.FailNow()
	}
}