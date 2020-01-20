package server

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pepeunlimited/files/rpcspaces"
	"github.com/pepeunlimited/microservice-kit/mail"
	"github.com/pepeunlimited/microservice-kit/rpcz"
	"github.com/pepeunlimited/users/internal/app/app1/mysql"
	"github.com/pepeunlimited/users/rpcusers"
	"github.com/twitchtv/twirp"
	"log"
	"testing"
)

var username string 		= "username"
var password string 		= "password"
var provider mail.Provider 	= mail.Mock

func TestUserServer_CreateUser(t *testing.T) {
	ctx := context.TODO()
	server := NewUserServer(mysql.NewEntClient(), username, password, provider, rpcspaces.NewSpacesMock(nil))
	server.users.DeleteAll(ctx)
	resp0, err := server.CreateUser(ctx, &rpcusers.CreateUserParams{
		Username: "simo",
		Password: "siimoo",
		Email:    "simo@gmail.com",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	user, err := server.GetUser(ctx, &rpcusers.GetUserParams{
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


func TestUserServer_SetProfilePictureFail(t *testing.T) {
	ctx := context.TODO()

	spacesMock := rpcspaces.NewSpacesMock(nil)
	server := NewUserServer(mysql.NewEntClient(), username, password, provider, spacesMock)
	server.users.DeleteAll(ctx)
	resp0,_ := server.CreateUser(ctx, &rpcusers.CreateUserParams{
		Username: "simo",
		Password: "siimoo",
		Email:    "simo@gmail.com",
	})
	spacesMock.(*rpcspaces.SpacesMock).File = &rpcspaces.File{
		UserId: 1111111,
	}

	ctx = rpcz.AddUserId(resp0.Id)
	_, err := server.SetProfilePicture(ctx, &rpcusers.SetProfilePictureParams{
		ProfilePictureId: 3,
	})
	if err == nil {
		t.FailNow()
	}
	if !rpcusers.IsReason(err.(twirp.Error), rpcusers.ProfilePictureAccessDenied) {
		t.FailNow()
	}
}

func TestUserServer_SetDeleteProfilePicture(t *testing.T) {
	ctx := context.TODO()

	spacesMock := rpcspaces.NewSpacesMock(nil)
	server := NewUserServer(mysql.NewEntClient(), username, password, provider, spacesMock)
	server.users.DeleteAll(ctx)
	resp0,_ := server.CreateUser(ctx, &rpcusers.CreateUserParams{
		Username: "simo",
		Password: "siimoo",
		Email:    "simo@gmail.com",
	})
	spacesMock.(*rpcspaces.SpacesMock).File = &rpcspaces.File{
		UserId: resp0.Id,
	}
	ctx = rpcz.AddUserId(resp0.Id)

	_, err := server.SetProfilePicture(ctx, &rpcusers.SetProfilePictureParams{
		ProfilePictureId: 3,
	})

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	user,_ := server.users.GetUserById(ctx, int(resp0.Id))
	if *user.ProfilePictureID != 3 {
		t.FailNow()
	}
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_, err = server.DeleteProfilePicture(ctx, &empty.Empty{})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestUserServer_CreateUserFail(t *testing.T) {
	ctx := context.TODO()
	server := NewUserServer(mysql.NewEntClient(), username, password, provider, rpcspaces.NewSpacesMock(nil))
	server.users.DeleteAll(ctx)
	_, err := server.CreateUser(ctx, &rpcusers.CreateUserParams{
		Username: "simo",
		Password: "siimoo",
		Email:    "simo@gmail.com",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_, err = server.CreateUser(ctx, &rpcusers.CreateUserParams{
		Username: "simo",
		Password: "siimoo",
		Email:    "simo2@gmail.com",
	})
	if err == nil {
		t.FailNow()
	}
	if !rpcusers.IsReason(err.(twirp.Error), rpcusers.UsernameExist) {
		t.Log(err.(twirp.Error).Error())
		t.FailNow()
	}
}

func TestUserServer_GetUserNotFound(t *testing.T) {
	ctx := context.TODO()
	server := NewUserServer(mysql.NewEntClient(), username, password, provider, rpcspaces.NewSpacesMock(nil))
	server.users.DeleteAll(ctx)
	ctx = rpcz.AddUserId(12312312)
	_, err := server.GetUser(ctx, &rpcusers.GetUserParams{
		UserId: 123123123,
	})
	if err == nil {
		t.FailNow()
	}

	if !rpcusers.IsReason(err.(twirp.Error), rpcusers.UserNotFound) {
		t.FailNow()
	}
}