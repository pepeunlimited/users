package server

import (
	"context"
	"github.com/pepeunlimited/microservice-kit/rpcz"
	"github.com/pepeunlimited/users/internal/app/app1/repository"
	rpc2 "github.com/pepeunlimited/users/rpc"
	"github.com/twitchtv/twirp"
	"testing"
)

func TestUserServer_CreateUser(t *testing.T) {
	ctx := context.TODO()
	server := NewUserServer(repository.NewEntClient())
	server.users.DeleteAll(ctx)
	_, err := server.CreateUser(ctx, &rpc2.CreateUserParams{
		Username: "simo",
		Password: "siimoo",
		Email:    "simo@gmail.com",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestUserServer_CreateUserFail(t *testing.T) {
	ctx := context.TODO()
	server := NewUserServer(repository.NewEntClient())
	server.users.DeleteAll(ctx)
	_, err := server.CreateUser(ctx, &rpc2.CreateUserParams{
		Username: "simo",
		Password: "siimoo",
		Email:    "simo@gmail.com",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_, err = server.CreateUser(ctx, &rpc2.CreateUserParams{
		Username: "simo",
		Password: "siimoo",
		Email:    "simo2@gmail.com",
	})
	if err == nil {
		t.FailNow()
	}
	if err.(twirp.Error).Meta(rpcz.Unique) != "username" {
		t.FailNow()
	}
}

func TestUserServer_GetUserNotFound(t *testing.T) {
	ctx := rpcz.AddUserId(3)
	server := NewUserServer(repository.NewEntClient())
	server.users.DeleteAll(ctx)
	_, err := server.GetUser(ctx, &rpc2.GetUserParams{})
	if err == nil {
		t.FailNow()
	}
	if err.(twirp.Error).Meta(rpcz.NotFound) != "user" {
		t.FailNow()
	}
}

func TestUserServer_SignInOk(t *testing.T) {
	ctx := context.TODO()
	server := NewUserServer(repository.NewEntClient())
	server.users.DeleteAll(ctx)

	email := "simo@gmail.com"
	username := email
	password := "p4sw0rd"

	server.CreateUser(ctx, &rpc2.CreateUserParams{
		Username: username,
		Password: password,
		Email:    email,
	})
	user, err := server.SignIn(ctx, &rpc2.SignInParams{
		Username: username,
		Password: password,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if user == nil {
		t.FailNow()
	}
}

func TestUserServer_SignInFail(t *testing.T) {
	ctx := context.TODO()
	server := NewUserServer(repository.NewEntClient())
	server.users.DeleteAll(ctx)
	_, err := server.SignIn(ctx, &rpc2.SignInParams{
		Username: "simo",
		Password: "p4sw0rd",
	})
	if err == nil {
		t.FailNow()
	}
	if err.(twirp.Error).Meta(rpcz.NotFound) != "user" {
		t.FailNow()
	}
}