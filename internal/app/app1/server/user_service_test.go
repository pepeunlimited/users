package server

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/wrappers"
	rpc2 "github.com/pepeunlimited/authorization-twirp/rpc"
	"github.com/pepeunlimited/files/rpcspaces"
	"github.com/pepeunlimited/microservice-kit/mail"
	"github.com/pepeunlimited/microservice-kit/rpcz"
	"github.com/pepeunlimited/users/internal/app/app1/mysql"
	"github.com/pepeunlimited/users/rpcusers"
	"github.com/twitchtv/twirp"
	"log"
	"testing"
	"time"
)

var username string 		= "username"
var password string 		= "password"
var provider mail.Provider 	= mail.Mock

func TestUserServer_CreateUser(t *testing.T) {
	ctx := context.TODO()
	server := NewUserServer(mysql.NewEntClient(), rpc2.NewAuthorizationMock(nil), username, password, provider, rpcspaces.NewSpacesMock(nil))
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
	server.authorization.(*rpc2.Mock).Username = "simo"
	server.authorization.(*rpc2.Mock).Email 	 = "simo@gmail.com"
	server.authorization.(*rpc2.Mock).UserId 	 = resp0.Id
	ctx = rpcz.AddAuthorization("1")
	user, err := server.GetUser(ctx, &rpcusers.GetUserParams{})
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
	server := NewUserServer(mysql.NewEntClient(), rpc2.NewAuthorizationMock(nil), username, password, provider, spacesMock)
	server.users.DeleteAll(ctx)
	resp0,_ := server.CreateUser(ctx, &rpcusers.CreateUserParams{
		Username: "simo",
		Password: "siimoo",
		Email:    "simo@gmail.com",
	})
	spacesMock.(*rpcspaces.SpacesMock).File = &rpcspaces.File{
		UserId: 1111111,
	}
	server.authorization.(*rpc2.Mock).Username = "simo"
	server.authorization.(*rpc2.Mock).Email 	 = "simo@gmail.com"
	server.authorization.(*rpc2.Mock).UserId 	 = resp0.Id
	ctx = rpcz.AddAuthorization("1")
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
	server := NewUserServer(mysql.NewEntClient(), rpc2.NewAuthorizationMock(nil), username, password, provider, spacesMock)
	server.users.DeleteAll(ctx)
	resp0,_ := server.CreateUser(ctx, &rpcusers.CreateUserParams{
		Username: "simo",
		Password: "siimoo",
		Email:    "simo@gmail.com",
	})
	spacesMock.(*rpcspaces.SpacesMock).File = &rpcspaces.File{
		UserId: resp0.Id,
	}
	server.authorization.(*rpc2.Mock).Username = "simo"
	server.authorization.(*rpc2.Mock).Email 	 = "simo@gmail.com"
	server.authorization.(*rpc2.Mock).UserId 	 = resp0.Id
	ctx = rpcz.AddAuthorization("1")

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
	_, err = server.DeleteProfilePicture(ctx, &rpcusers.DeleteProfilePictureParams{})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestUserServer_CreateUserFail(t *testing.T) {
	ctx := context.TODO()
	server := NewUserServer(mysql.NewEntClient(), rpc2.NewAuthorizationMock(nil), username, password, provider, rpcspaces.NewSpacesMock(nil))
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
	ctx := rpcz.AddUserId(3)
	server := NewUserServer(mysql.NewEntClient(), rpc2.NewAuthorizationMock([]error{fmt.Errorf("custom-error")}), username, password, provider, rpcspaces.NewSpacesMock(nil))
	server.users.DeleteAll(ctx)
	ctx = rpcz.AddAuthorization("1")
	_, err := server.GetUser(ctx, &rpcusers.GetUserParams{})
	if err == nil {
		t.FailNow()
	}
	if err.Error() != "custom-error" {
		t.FailNow()
	}
}

func TestUserServer_SignInOk(t *testing.T) {
	ctx := context.TODO()
	server := NewUserServer(mysql.NewEntClient(), rpc2.NewAuthorizationMock(nil), username, password, provider, rpcspaces.NewSpacesMock(nil))
	server.users.DeleteAll(ctx)

	email := "simo@gmail.com"
	username := email
	password := "p4sw0rd"

	user0, err := server.CreateUser(ctx, &rpcusers.CreateUserParams{
		Username: username,
		Password: password,
		Email:    email,
	})

	server.authorization.(*rpc2.Mock).Username 	= username
	server.authorization.(*rpc2.Mock).Email 		= email
	server.authorization.(*rpc2.Mock).Roles 		= []string{"User"}
	server.authorization.(*rpc2.Mock).UserId		= user0.Id

	user, err := server.VerifySignIn(ctx, &rpcusers.VerifySignInParams{
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
	server := NewUserServer(mysql.NewEntClient(), rpc2.NewAuthorizationMock(nil), username, password, provider, rpcspaces.NewSpacesMock(nil))
	server.users.DeleteAll(ctx)
	_, err := server.VerifySignIn(ctx, &rpcusers.VerifySignInParams{
		Username: "simo",
		Password: "p4sw0rd",
	})
	if err == nil {
		t.FailNow()
	}
	if !rpcusers.IsReason(err.(twirp.Error), rpcusers.UserNotFound) {
		t.Log(err.(twirp.Error).Error())
		t.FailNow()
	}

}

func TestUserServer_SignInFailCred(t *testing.T) {
	ctx := context.TODO()
	server := NewUserServer(mysql.NewEntClient(), rpc2.NewAuthorizationMock(nil), username, password, provider, rpcspaces.NewSpacesMock(nil))
	server.users.DeleteAll(ctx)
	_, err := server.VerifySignIn(ctx, &rpcusers.VerifySignInParams{
		Username: "simo",
		Password: "p4sw0rd",
	})
	if err == nil {
		t.FailNow()
	}
	if !rpcusers.IsReason(err.(twirp.Error), rpcusers.UserNotFound) {
		t.Log(err.(twirp.Error).Error())
		t.FailNow()
	}

}

func TestUserServer_ForgotPasswordSuccess(t *testing.T) {
	ctx := context.TODO()
	server := NewUserServer(mysql.NewEntClient(), rpc2.NewAuthorizationMock(nil), username, password, provider, rpcspaces.NewSpacesMock(nil))
	server.users.DeleteAll(ctx)
	username := "simo"
	user,_ := server.CreateUser(ctx, &rpcusers.CreateUserParams{
		Username: username,
		Password: "p4sw04d",
		Email:    "simo@gmail.com",
	})
	_, err := server.ForgotPassword(ctx, &rpcusers.ForgotPasswordParams{
		Email: &wrappers.StringValue{
			Value: user.Email,
		},
		Language: "fi",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_, tickets, err := server.users.GetUserTicketsByUserId(ctx, int(user.Id))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if len(tickets) != 1 {
		t.FailNow()
	}
}

func TestUserServer_ForgotPasswordFailure1(t *testing.T) {
	ctx := context.TODO()
	server := NewUserServer(mysql.NewEntClient(), rpc2.NewAuthorizationMock(nil), username, password, mail.MockFail, rpcspaces.NewSpacesMock(nil))
	server.users.DeleteAll(ctx)
	username := "simo"
	user,_ := server.CreateUser(ctx, &rpcusers.CreateUserParams{
		Username: username,
		Password: "p4sw04d",
		Email:    "simo@gmail.com",
	})
	_, err := server.ForgotPassword(ctx, &rpcusers.ForgotPasswordParams{
		Username: &wrappers.StringValue{
			Value: username,
		},
		Language: "fi",
	})
	if err == nil {
		t.FailNow()
	}
	if !rpcusers.IsReason(err.(twirp.Error), mail.MailSendFailed) {
		t.FailNow()
	}
	_, tickets, err := server.users.GetUserTicketsByUserId(ctx, int(user.Id))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if len(tickets) != 0 {
		t.FailNow()
	}
}

func TestUserServer_ForgotPasswordFailure2(t *testing.T) {
	ctx := context.TODO()
	server := NewUserServer(mysql.NewEntClient(), rpc2.NewAuthorizationMock(nil), username, password, mail.MockFail, rpcspaces.NewSpacesMock(nil))
	server.users.DeleteAll(ctx)
	username := "simo"
	_, err := server.ForgotPassword(ctx, &rpcusers.ForgotPasswordParams{
		Username: &wrappers.StringValue{
			Value: username,
		},
		Language: "fi",
	})
	if err == nil {
		t.FailNow()
	}
	if !rpcusers.IsReason(err.(twirp.Error), rpcusers.UserNotFound) {
		t.FailNow()
	}
}

func TestUserServer_VerifyResetPasswordExpired(t *testing.T) {
	ctx := context.TODO()
	server := NewUserServer(mysql.NewEntClient(), rpc2.NewAuthorizationMock(nil), username, password, mail.Mock, rpcspaces.NewSpacesMock(nil))
	server.users.DeleteAll(ctx)
	user,_ := server.CreateUser(ctx, &rpcusers.CreateUserParams{
		Username: "simo",
		Password: "simo",
		Email:    "simo@gmail.com",
	})
	ticket,_ := server.tickets.CreateTicket(ctx, time.Now().UTC().Add(1*time.Second), int(user.Id))
	time.Sleep(2 * time.Second)
	_, err := server.VerifyResetPassword(ctx, &rpcusers.VerifyPasswordParams{TicketToken: ticket.Token})
	if err == nil {
		t.FailNow()
	}
	if !rpcusers.IsReason(err.(twirp.Error), rpcusers.TicketExpired) {
		t.FailNow()
	}
}

func TestUserServer_VerifyResetPasswordNotFound(t *testing.T) {
	ctx := context.TODO()
	server := NewUserServer(mysql.NewEntClient(), rpc2.NewAuthorizationMock(nil), username, password, mail.Mock, rpcspaces.NewSpacesMock(nil))
	server.users.DeleteAll(ctx)
	server.CreateUser(ctx, &rpcusers.CreateUserParams{
		Username: "simo",
		Password: "simo",
		Email:    "simo@gmail.com",
	})
	_, err := server.VerifyResetPassword(ctx, &rpcusers.VerifyPasswordParams{TicketToken: "asd"})
	if err == nil {
		t.FailNow()
	}
	if !rpcusers.IsReason(err.(twirp.Error), rpcusers.TicketNotFound) {
		t.FailNow()
	}
}

func TestUserServer_VerifyResetPasswordAndResetPasswordSuccess(t *testing.T) {
	ctx := context.TODO()
	server := NewUserServer(mysql.NewEntClient(), rpc2.NewAuthorizationMock(nil), username, password, provider, rpcspaces.NewSpacesMock(nil))
	server.users.DeleteAll(ctx)
	username := "simo"
	user,_ := server.CreateUser(ctx, &rpcusers.CreateUserParams{
		Username: username,
		Password: "p4sw04d",
		Email:    "simo@gmail.com",
	})
	server.ForgotPassword(ctx, &rpcusers.ForgotPasswordParams{
		Username: &wrappers.StringValue{
			Value: username,
		},
		Language: "fi",
	})
	_, tickets,_ := server.users.GetUserTicketsByUserId(ctx, int(user.Id))
	if len(tickets) != 1 {
		t.FailNow()
	}
	token := tickets[0].Token
	_, err := server.VerifyResetPassword(ctx, &rpcusers.VerifyPasswordParams{
		TicketToken: token,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_, err = server.ResetPassword(ctx, &rpcusers.ResetPasswordParams{
		TicketToken:    token,
		Password: "newpw",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_, err = server.VerifySignIn(ctx, &rpcusers.VerifySignInParams{
		Username: username,
		Password: "newpw",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_, tickets2,_ := server.users.GetUserTicketsByUserId(ctx, int(user.Id))
	if len(tickets2) != 0 {
		t.FailNow()
	}
}

func TestUserServer_VerifyResetPasswordAndResetPasswordSuccess2(t *testing.T) {
	ctx := context.TODO()
	server := NewUserServer(mysql.NewEntClient(), rpc2.NewAuthorizationMock(nil), username, password, provider, rpcspaces.NewSpacesMock(nil))
	server.users.DeleteAll(ctx)
	username := "simo"
	user,_ := server.CreateUser(ctx, &rpcusers.CreateUserParams{
		Username: username,
		Password: "p4sw04d",
		Email:    "simo@gmail.com",
	})
	server.ForgotPassword(ctx, &rpcusers.ForgotPasswordParams{
		Email: &wrappers.StringValue{Value: user.Email},
		Language: "fi",
	})
	_, tickets,_ := server.users.GetUserTicketsByUserId(ctx, int(user.Id))
	if len(tickets) != 1 {
		t.FailNow()
	}
	token := tickets[0].Token
	_, err := server.VerifyResetPassword(ctx, &rpcusers.VerifyPasswordParams{
		TicketToken: token,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_, err = server.ResetPassword(ctx, &rpcusers.ResetPasswordParams{
		TicketToken:    token,
		Password: "newpw",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_, err = server.VerifySignIn(ctx, &rpcusers.VerifySignInParams{
		Username: username,
		Password: "newpw",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_, tickets2,_ := server.users.GetUserTicketsByUserId(ctx, int(user.Id))
	if len(tickets2) != 0 {
		t.FailNow()
	}
}

func TestUserServer_UpdatePassword(t *testing.T) {
	ctx := context.TODO()
	server := NewUserServer(mysql.NewEntClient(), rpc2.NewAuthorizationMock(nil), username, password, provider, rpcspaces.NewSpacesMock(nil))
	server.users.DeleteAll(ctx)

	email := "simo@gmail.com"
	username := email
	password := "p4sw0rd"

	user0,_ := server.CreateUser(ctx, &rpcusers.CreateUserParams{
		Username: username,
		Password: password,
		Email:    email,
	})
	server.authorization.(*rpc2.Mock).Username 	= username
	server.authorization.(*rpc2.Mock).Email 		= email
	server.authorization.(*rpc2.Mock).Roles 		= []string{"User"}
	server.authorization.(*rpc2.Mock).UserId		= user0.Id
	ctx = rpcz.AddAuthorization("token")
	_, err := server.UpdatePassword(ctx, &rpcusers.UpdatePasswordParams{
		CurrentPassword: password,
		NewPassword:     "newpw",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestUserServer_UpdatePasswordFail(t *testing.T) {
	ctx := context.TODO()
	server := NewUserServer(mysql.NewEntClient(), rpc2.NewAuthorizationMock(nil), username, password, provider, rpcspaces.NewSpacesMock(nil))
	server.users.DeleteAll(ctx)

	email := "simo@gmail.com"
	username := email
	password := "p4sw0rd"

	user0,_ := server.CreateUser(ctx, &rpcusers.CreateUserParams{
		Username: username,
		Password: password,
		Email:    email,
	})
	server.authorization.(*rpc2.Mock).Username 	= username
	server.authorization.(*rpc2.Mock).Email 		= email
	server.authorization.(*rpc2.Mock).Roles 		= []string{"User"}
	server.authorization.(*rpc2.Mock).UserId		= user0.Id
	ctx = rpcz.AddAuthorization("token")
	_, err := server.UpdatePassword(ctx, &rpcusers.UpdatePasswordParams{
		CurrentPassword: "wronpw",
		NewPassword:     "newpw",
	})
	if err == nil {
		t.FailNow()
	}
	if !rpcusers.IsReason(err.(twirp.Error), rpcusers.InvalidCredentials) {
		t.FailNow()
	}
}