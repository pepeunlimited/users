package server

import (
	"context"
	rpc2 "github.com/pepeunlimited/authorization-twirp/rpc"
	"github.com/pepeunlimited/microservice-kit/rpcz"
	"github.com/pepeunlimited/users/internal/app/app1/repository"
	"github.com/pepeunlimited/users/rpc"
	"github.com/twitchtv/twirp"
	"log"
	"testing"
)

func TestUserServer_CreateUser(t *testing.T) {
	ctx := context.TODO()
	server := NewUserServer(repository.NewEntClient(), rpc2.NewAuthorizationMock(nil))
	server.users.DeleteAll(ctx)
	resp0, err := server.CreateUser(ctx, &rpc.CreateUserParams{
		Username: "simo",
		Password: "siimoo",
		Email:    "simo@gmail.com",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	server.authService.(*rpc2.Mock).Username = "simo"
	server.authService.(*rpc2.Mock).Email 	 = "simo@gmail.com"
	server.authService.(*rpc2.Mock).UserId 	 = resp0.Id
	ctx = rpcz.AddAuthorization("1")
	user, err := server.GetUser(ctx, &rpc.GetUserParams{})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	log.Print(user)
	if user.Email != "simo@gmail.com" {
		t.FailNow()
	}
	if user.Username != "simo" {
		t.FailNow()
	}
	if user.Id != resp0.Id {
		t.FailNow()
	}
}

func TestUserServer_CreateUserFail(t *testing.T) {
	ctx := context.TODO()
	server := NewUserServer(repository.NewEntClient(), rpc2.NewAuthorizationMock(nil))
	server.users.DeleteAll(ctx)
	_, err := server.CreateUser(ctx, &rpc.CreateUserParams{
		Username: "simo",
		Password: "siimoo",
		Email:    "simo@gmail.com",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_, err = server.CreateUser(ctx, &rpc.CreateUserParams{
		Username: "simo",
		Password: "siimoo",
		Email:    "simo2@gmail.com",
	})
	if err == nil {
		t.FailNow()
	}
	if !rpc.IsReason(err.(twirp.Error), rpc.UsernameExist) {
		t.Log(err.(twirp.Error).Error())
		t.FailNow()
	}
}

func TestUserServer_GetUserNotFound(t *testing.T) {
	ctx := rpcz.AddUserId(3)
	server := NewUserServer(repository.NewEntClient(), rpc2.NewAuthorizationMock(nil))
	server.users.DeleteAll(ctx)
	_, err := server.GetUser(ctx, &rpc.GetUserParams{})
	if err == nil {
		t.FailNow()
	}
	if !rpc.IsReason(err.(twirp.Error), rpc.UserNotFound) {
		t.Log(err.(twirp.Error).Error())
		t.FailNow()
	}
}

func TestUserServer_SignInOk(t *testing.T) {
	ctx := context.TODO()
	server := NewUserServer(repository.NewEntClient(), rpc2.NewAuthorizationMock(nil))
	server.users.DeleteAll(ctx)

	email := "simo@gmail.com"
	username := email
	password := "p4sw0rd"

	server.CreateUser(ctx, &rpc.CreateUserParams{
		Username: username,
		Password: password,
		Email:    email,
	})
	user, err := server.VerifySignIn(ctx, &rpc.VerifySignInParams{
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
	server := NewUserServer(repository.NewEntClient(), rpc2.NewAuthorizationMock(nil))
	server.users.DeleteAll(ctx)
	_, err := server.VerifySignIn(ctx, &rpc.VerifySignInParams{
		Username: "simo",
		Password: "p4sw0rd",
	})
	if err == nil {
		t.FailNow()
	}
	if !rpc.IsReason(err.(twirp.Error), rpc.UserNotFound) {
		t.Log(err.(twirp.Error).Error())
		t.FailNow()
	}

}

func TestUserServer_SignInFailCred(t *testing.T) {
	ctx := context.TODO()
	server := NewUserServer(repository.NewEntClient(), rpc2.NewAuthorizationMock(nil))
	server.users.DeleteAll(ctx)
	_, err := server.VerifySignIn(ctx, &rpc.VerifySignInParams{
		Username: "simo",
		Password: "p4sw0rd",
	})
	if err == nil {
		t.FailNow()
	}
	if !rpc.IsReason(err.(twirp.Error), rpc.UserNotFound) {
		t.Log(err.(twirp.Error).Error())
		t.FailNow()
	}

}