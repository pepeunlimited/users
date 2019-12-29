package main

import (
	rpc2 "github.com/pepeunlimited/authorization-twirp/rpc"
	"github.com/pepeunlimited/microservice-kit/headers"
	"github.com/pepeunlimited/microservice-kit/middleware"
	"github.com/pepeunlimited/microservice-kit/misc"
	"github.com/pepeunlimited/users/internal/app/app1/repository"
	"github.com/pepeunlimited/users/internal/app/app1/server"
	"github.com/pepeunlimited/users/rpc"
	"log"
	"net/http"
)

const (
	Version = "0.1.2"
)

func main() {
	log.Printf("Starting the UsersServer... version=[%v]", Version)

	client := repository.NewEntClient()
	authorizationAddress := misc.GetEnv(rpc2.RpcAuthorizationHost, "http://localhost:8080")
	us := rpc.NewUserServiceServer(server.NewUserServer(client, rpc2.NewAuthorizationServiceProtobufClient(authorizationAddress, http.DefaultClient)), nil)
	mux := http.NewServeMux()
	mux.Handle(us.PathPrefix(), middleware.Adapt(us, headers.Authorizationz()))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Panic(err)
	}
}