package main

import (
	rpc2 "github.com/pepeunlimited/authorization-twirp/rpc"
	rpc3 "github.com/pepeunlimited/files/rpc"
	"github.com/pepeunlimited/microservice-kit/headers"
	"github.com/pepeunlimited/microservice-kit/mail"
	"github.com/pepeunlimited/microservice-kit/middleware"
	"github.com/pepeunlimited/microservice-kit/misc"
	"github.com/pepeunlimited/users/internal/app/app1/repository"
	"github.com/pepeunlimited/users/internal/app/app1/server"
	"github.com/pepeunlimited/users/rpc"
	"log"
	"net/http"
)

const (
	Version = "0.1.3.11"
)

func main() {
	log.Printf("Starting the UsersServer... version=[%v]", Version)

	client := repository.NewEntClient()
	authorizationAddress := misc.GetEnv(rpc2.RpcAuthorizationHost, "http://api.dev.pepeunlimited.com")
	spacesAddress 		 := misc.GetEnv(rpc3.RpcFilesHost, "http://api.dev.pepeunlimited.com")

	stmpUsername := misc.GetEnv(mail.SmtpPassword, "us3rn4m3")
	stmpPassword := misc.GetEnv(mail.SmtpPassword, "p4sw0rd")
	smtpProvider := mail.Provider(misc.GetEnv(mail.SmtpClient,   mail.Mock))

	us := rpc.NewUserServiceServer(server.NewUserServer(
		client,
		rpc2.NewAuthorizationServiceProtobufClient(authorizationAddress, http.DefaultClient),
		stmpUsername,
		stmpPassword,
		smtpProvider,
		rpc3.NewSpacesServiceProtobufClient(spacesAddress, http.DefaultClient)),
		nil)

	mux := http.NewServeMux()
	mux.Handle(us.PathPrefix(), middleware.Adapt(us, headers.Authorizationz()))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Panic(err)
	}
}