package server

import (
	"context"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/pepeunlimited/authorization-twirp/rpcauthorization"
	"github.com/pepeunlimited/files/rpcspaces"
	"github.com/pepeunlimited/microservice-kit/mail"
	"github.com/pepeunlimited/microservice-kit/rpcz"
	"github.com/pepeunlimited/users/internal/app/app1/mysql"
	"github.com/pepeunlimited/users/rpccredentials"
	"github.com/pepeunlimited/users/rpcusers"
	"github.com/twitchtv/twirp"
	"testing"
	"time"
)

func TestUserServer_SignInOk(t *testing.T) {
	ctx := context.TODO()
	server := NewCredentialsServer(mysql.NewEntClient(), rpcauthorization.NewAuthorizationMock(nil), username, password, provider)
	server.users.DeleteAll(ctx)

	email := "simo@gmail.com"
	username := email
	password := "p4sw0rd"

	userServer := NewUserServer(mysql.NewEntClient(), rpcauthorization.NewAuthorizationMock(nil), username, password, provider, rpcspaces.NewSpacesMock(nil))
	user0, err := userServer.CreateUser(ctx, &rpcusers.CreateUserParams{
		Username: username,
		Password: password,
		Email:    email,
	})

	server.authorization.(*rpcauthorization.Mock).Username 	= username
	server.authorization.(*rpcauthorization.Mock).Email 	= email
	server.authorization.(*rpcauthorization.Mock).Roles 	= []string{"User"}
	server.authorization.(*rpcauthorization.Mock).UserId	= user0.Id

	user, err := server.VerifySignIn(ctx, &rpccredentials.VerifySignInParams{
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
	server := NewCredentialsServer(mysql.NewEntClient(), rpcauthorization.NewAuthorizationMock(nil), username, password, provider)
	server.users.DeleteAll(ctx)
	_, err := server.VerifySignIn(ctx, &rpccredentials.VerifySignInParams{
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
	server := NewCredentialsServer(mysql.NewEntClient(), rpcauthorization.NewAuthorizationMock(nil), username, password, provider)
	server.users.DeleteAll(ctx)
	_, err := server.VerifySignIn(ctx, &rpccredentials.VerifySignInParams{
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
	server := NewCredentialsServer(mysql.NewEntClient(), rpcauthorization.NewAuthorizationMock(nil), username, password, provider)
	server.users.DeleteAll(ctx)
	username := "simo"

	userServer := NewUserServer(mysql.NewEntClient(), rpcauthorization.NewAuthorizationMock(nil), username, password, provider, rpcspaces.NewSpacesMock(nil))
	user,_ := userServer.CreateUser(ctx, &rpcusers.CreateUserParams{
		Username: username,
		Password: "p4sw04d",
		Email:    "simo@gmail.com",
	})
	_, err := server.ForgotPassword(ctx, &rpccredentials.ForgotPasswordParams{
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
	server := NewCredentialsServer(mysql.NewEntClient(), rpcauthorization.NewAuthorizationMock(nil), username, password, mail.MockFail)
	server.users.DeleteAll(ctx)
	username := "simo"
	userServer := NewUserServer(mysql.NewEntClient(), rpcauthorization.NewAuthorizationMock(nil), username, password, provider, rpcspaces.NewSpacesMock(nil))
	user,_ := userServer.CreateUser(ctx, &rpcusers.CreateUserParams{
		Username: username,
		Password: "p4sw04d",
		Email:    "simo@gmail.com",
	})
	_, err := server.ForgotPassword(ctx, &rpccredentials.ForgotPasswordParams{
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
	server := NewCredentialsServer(mysql.NewEntClient(), rpcauthorization.NewAuthorizationMock(nil), username, password, provider)

	server.users.DeleteAll(ctx)
	username := "simo"

	_, err := server.ForgotPassword(ctx, &rpccredentials.ForgotPasswordParams{
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

	server := NewCredentialsServer(mysql.NewEntClient(), rpcauthorization.NewAuthorizationMock(nil), username, password, provider)

	userServer := NewUserServer(mysql.NewEntClient(), rpcauthorization.NewAuthorizationMock(nil), username, password, provider, rpcspaces.NewSpacesMock(nil))

	server.users.DeleteAll(ctx)
	user,_ := userServer.CreateUser(ctx, &rpcusers.CreateUserParams{
		Username: "simo",
		Password: "simo",
		Email:    "simo@gmail.com",
	})
	ticket,_ := server.tickets.CreateTicket(ctx, time.Now().UTC().Add(1*time.Second), int(user.Id))
	time.Sleep(2 * time.Second)
	_, err := server.VerifyResetPassword(ctx, &rpccredentials.VerifyPasswordParams{TicketToken: ticket.Token})
	if err == nil {
		t.FailNow()
	}
	if !rpcusers.IsReason(err.(twirp.Error), rpcusers.TicketExpired) {
		t.FailNow()
	}
}

func TestUserServer_VerifyResetPasswordNotFound(t *testing.T) {
	ctx := context.TODO()
	server := NewCredentialsServer(mysql.NewEntClient(), rpcauthorization.NewAuthorizationMock(nil), username, password, provider)
	server.users.DeleteAll(ctx)
	userServer := NewUserServer(mysql.NewEntClient(), rpcauthorization.NewAuthorizationMock(nil), username, password, provider, rpcspaces.NewSpacesMock(nil))
	userServer.CreateUser(ctx, &rpcusers.CreateUserParams{
		Username: "simo",
		Password: "simo",
		Email:    "simo@gmail.com",
	})
	_, err := server.VerifyResetPassword(ctx, &rpccredentials.VerifyPasswordParams{TicketToken: "asd"})
	if err == nil {
		t.FailNow()
	}
	if !rpcusers.IsReason(err.(twirp.Error), rpcusers.TicketNotFound) {
		t.FailNow()
	}
}

func TestUserServer_VerifyResetPasswordAndResetPasswordSuccess(t *testing.T) {
	ctx := context.TODO()
	server := NewCredentialsServer(mysql.NewEntClient(), rpcauthorization.NewAuthorizationMock(nil), username, password, provider)
	userServer := NewUserServer(mysql.NewEntClient(), rpcauthorization.NewAuthorizationMock(nil), username, password, provider, rpcspaces.NewSpacesMock(nil))


	server.users.DeleteAll(ctx)
	username := "simo"
	user,_ := userServer.CreateUser(ctx, &rpcusers.CreateUserParams{
		Username: username,
		Password: "p4sw04d",
		Email:    "simo@gmail.com",
	})
	server.ForgotPassword(ctx, &rpccredentials.ForgotPasswordParams{
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
	_, err := server.VerifyResetPassword(ctx, &rpccredentials.VerifyPasswordParams{
		TicketToken: token,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_, err = server.ResetPassword(ctx, &rpccredentials.ResetPasswordParams{
		TicketToken:    token,
		Password: "newpw",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_, err = server.VerifySignIn(ctx, &rpccredentials.VerifySignInParams{
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
	server := NewCredentialsServer(mysql.NewEntClient(), rpcauthorization.NewAuthorizationMock(nil), username, password, provider)
	userServer := NewUserServer(mysql.NewEntClient(), rpcauthorization.NewAuthorizationMock(nil), username, password, provider, rpcspaces.NewSpacesMock(nil))
	server.users.DeleteAll(ctx)
	username := "simo"
	user,_ := userServer.CreateUser(ctx, &rpcusers.CreateUserParams{
		Username: username,
		Password: "p4sw04d",
		Email:    "simo@gmail.com",
	})
	server.ForgotPassword(ctx, &rpccredentials.ForgotPasswordParams{
		Email: &wrappers.StringValue{Value: user.Email},
		Language: "fi",
	})
	_, tickets,_ := server.users.GetUserTicketsByUserId(ctx, int(user.Id))
	if len(tickets) != 1 {
		t.FailNow()
	}
	token := tickets[0].Token
	_, err := server.VerifyResetPassword(ctx, &rpccredentials.VerifyPasswordParams{
		TicketToken: token,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_, err = server.ResetPassword(ctx, &rpccredentials.ResetPasswordParams{
		TicketToken:    token,
		Password: "newpw",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_, err = server.VerifySignIn(ctx, &rpccredentials.VerifySignInParams{
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

	server := NewCredentialsServer(mysql.NewEntClient(), rpcauthorization.NewAuthorizationMock(nil), username, password, provider)
	userServer := NewUserServer(mysql.NewEntClient(), rpcauthorization.NewAuthorizationMock(nil), username, password, provider, rpcspaces.NewSpacesMock(nil))
	server.users.DeleteAll(ctx)

	email := "simo@gmail.com"
	username := email
	password := "p4sw0rd"

	user0,_ := userServer.CreateUser(ctx, &rpcusers.CreateUserParams{
		Username: username,
		Password: password,
		Email:    email,
	})
	server.authorization.(*rpcauthorization.Mock).Username 	= username
	server.authorization.(*rpcauthorization.Mock).Email 	= email
	server.authorization.(*rpcauthorization.Mock).Roles 	= []string{"User"}
	server.authorization.(*rpcauthorization.Mock).UserId	= user0.Id
	ctx = rpcz.AddAuthorization("token")
	_, err := server.UpdatePassword(ctx, &rpccredentials.UpdatePasswordParams{
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
	server := NewCredentialsServer(mysql.NewEntClient(), rpcauthorization.NewAuthorizationMock(nil), username, password, provider)
	userServer := NewUserServer(mysql.NewEntClient(), rpcauthorization.NewAuthorizationMock(nil), username, password, provider, rpcspaces.NewSpacesMock(nil))
	server.users.DeleteAll(ctx)

	email := "simo@gmail.com"
	username := email
	password := "p4sw0rd"

	user0,_ := userServer.CreateUser(ctx, &rpcusers.CreateUserParams{
		Username: username,
		Password: password,
		Email:    email,
	})
	server.authorization.(*rpcauthorization.Mock).Username 	= username
	server.authorization.(*rpcauthorization.Mock).Email 	= email
	server.authorization.(*rpcauthorization.Mock).Roles 	= []string{"User"}
	server.authorization.(*rpcauthorization.Mock).UserId	= user0.Id
	ctx = rpcz.AddAuthorization("token")
	_, err := server.UpdatePassword(ctx, &rpccredentials.UpdatePasswordParams{
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