package twirp

import (
	"context"
	"github.com/pepeunlimited/microservice-kit/mail"
	"github.com/pepeunlimited/microservice-kit/rpcz"
	"github.com/pepeunlimited/users/internal/pkg/ent"
	"github.com/pepeunlimited/users/pkg/rpc/user"
	"github.com/twitchtv/twirp"
	"log"
	"testing"
)

var username string 		= "username"
var password string 		= "password"
var provider mail.Provider 	= mail.Mock

func TestUserServer_CreateUser(t *testing.T) {
	ctx := context.TODO()
	server := NewUserServer(ent.NewEntClient(), username, password, provider)
	server.users.DeleteAll(ctx)
	resp0, err := server.CreateUser(ctx, &user.CreateUserParams{
		Username: "simo",
		Password: "siimoo",
		Email:    "simo@gmail.com",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	user, err := server.GetUser(ctx, &user.GetUserParams{
		UserId: resp0.Id,
	})
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

func TestUserServer_SetDeleteProfilePicture(t *testing.T) {
	ctx := context.TODO()
	server := NewUserServer(ent.NewEntClient(), username, password, provider)
	server.users.DeleteAll(ctx)
	resp0,_ := server.CreateUser(ctx, &user.CreateUserParams{
		Username: "simo",
		Password: "siimoo",
		Email:    "simo@gmail.com",
	})

	_, err := server.SetProfilePicture(ctx, &user.SetProfilePictureParams{
		ProfilePictureId: 3,
		UserId: resp0.Id,
	})

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	fromDB,_ := server.users.GetUserById(ctx, int(resp0.Id))
	if *fromDB.ProfilePictureID != 3 {
		t.FailNow()
	}
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_, err = server.DeleteProfilePicture(ctx, &user.DeleteProfilePictureParams{
		UserId: resp0.Id,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestUserServer_CreateUserFail(t *testing.T) {
	ctx := context.TODO()
	server := NewUserServer(ent.NewEntClient(), username, password, provider)
	server.users.DeleteAll(ctx)
	_, err := server.CreateUser(ctx, &user.CreateUserParams{
		Username: "simo",
		Password: "siimoo",
		Email:    "simo@gmail.com",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_, err = server.CreateUser(ctx, &user.CreateUserParams{
		Username: "simo",
		Password: "siimoo",
		Email:    "simo2@gmail.com",
	})
	if err == nil {
		t.FailNow()
	}
	if err.(twirp.Error).Msg() != user.UsernameExist {
		t.Log(err.(twirp.Error).Error())
		t.FailNow()
	}
}

func TestUserServer_GetUserNotFound(t *testing.T) {
	ctx := context.TODO()
	server := NewUserServer(ent.NewEntClient(), username, password, provider)
	server.users.DeleteAll(ctx)
	ctx = rpcz.AddUserId(12312312)
	_, err := server.GetUser(ctx, &user.GetUserParams{
		UserId: 123123123,
	})
	if err == nil {
		t.FailNow()
	}
	if err.(twirp.Error).Msg() != user.UserNotFound {
		t.FailNow()
	}
}