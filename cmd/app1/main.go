package main

import (
	"github.com/pepeunlimited/authorization-twirp/rpcauthorization"
	"github.com/pepeunlimited/files/rpcspaces"
	"github.com/pepeunlimited/microservice-kit/headers"
	"github.com/pepeunlimited/microservice-kit/mail"
	"github.com/pepeunlimited/microservice-kit/middleware"
	"github.com/pepeunlimited/microservice-kit/misc"
	"github.com/pepeunlimited/users/internal/app/app1/mysql"
	"github.com/pepeunlimited/users/internal/app/app1/server"
	"github.com/pepeunlimited/users/rpccredentials"
	"github.com/pepeunlimited/users/rpcusers"
	"log"
	"net/http"
)

const (
	Version = "0.1.3.14"
)

func main() {
	log.Printf("Starting the UsersServer... version=[%v]", Version)

	client := mysql.NewEntClient()
	authorizationAddress := misc.GetEnv(rpcauthorization.RpcAuthorizationHost, "http://api.dev.pepeunlimited.com")
	spacesAddress 		 := misc.GetEnv(rpcspaces.RpcSpacesHost, "http://api.dev.pepeunlimited.com")

	stmpUsername := misc.GetEnv(mail.SmtpPassword, "us3rn4m3")
	stmpPassword := misc.GetEnv(mail.SmtpPassword, "p4sw0rd")
	smtpProvider := mail.Provider(misc.GetEnv(mail.SmtpClient,   mail.Mock))


	css := rpccredentials.NewCredentialsServiceServer(server.NewCredentialsServer(
		client,
		rpcauthorization.NewAuthorizationServiceProtobufClient(authorizationAddress, http.DefaultClient),
		stmpUsername,
		stmpPassword,
		smtpProvider),nil)


	uss := rpcusers.NewUserServiceServer(server.NewUserServer(
		client,
		rpcauthorization.NewAuthorizationServiceProtobufClient(authorizationAddress, http.DefaultClient),
		stmpUsername,
		stmpPassword,
		smtpProvider,
		rpcspaces.NewSpacesServiceProtobufClient(spacesAddress, http.DefaultClient)),
		nil)

	mux := http.NewServeMux()
	mux.Handle(uss.PathPrefix(), middleware.Adapt(uss, headers.Authorizationz()))
	mux.Handle(css.PathPrefix(), middleware.Adapt(css, headers.Authorizationz()))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Panic(err)
	}
}