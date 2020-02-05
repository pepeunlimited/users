package main

import (
	"github.com/pepeunlimited/microservice-kit/mail"
	"github.com/pepeunlimited/microservice-kit/middleware"
	"github.com/pepeunlimited/microservice-kit/misc"
	"github.com/pepeunlimited/users/internal/pkg/ent"
	"github.com/pepeunlimited/users/internal/server/twirp"
	"github.com/pepeunlimited/users/pkg/credentialsrpc"
	"github.com/pepeunlimited/users/pkg/usersrpc"
	"log"
	"net/http"
)

const (
	Version = "0.1.3.21"
)

func main() {
	log.Printf("Starting the UsersServer... version=[%v]", Version)

	client 		  := ent.NewEntClient()

	stmpUsername  := misc.GetEnv(mail.SmtpPassword, "us3rn4m3")
	stmpPassword  := misc.GetEnv(mail.SmtpPassword, "p4sw0rd")
	smtpProvider  := mail.Provider(misc.GetEnv(mail.SmtpClient, mail.Mock))

	css := credentialsrpc.NewCredentialsServiceServer(twirp.NewCredentialsServer(
		client,
		stmpUsername,
		stmpPassword,
		smtpProvider),nil)

	uss := usersrpc.NewUserServiceServer(twirp.NewUserServer(
		client,
		stmpUsername,
		stmpPassword,
		smtpProvider),
		nil)

	mux := http.NewServeMux()
	mux.Handle(uss.PathPrefix(), middleware.Adapt(uss))
	mux.Handle(css.PathPrefix(), middleware.Adapt(css))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Panic(err)
	}

}