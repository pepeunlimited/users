package main

import (
	"github.com/pepeunlimited/microservice-kit/headers"
	"github.com/pepeunlimited/microservice-kit/middleware"
	"github.com/pepeunlimited/users/internal/app/app1/repository"
	"github.com/pepeunlimited/users/internal/app/app1/server"
	"github.com/pepeunlimited/users/rpc"
	"log"
	"net/http"
)

const (
	Version = "0.1"
)

func main() {
	log.Printf("Starting the UsersServer... version=[%v]", Version)

	client := repository.NewEntClient()
	us := rpc.NewUserServiceServer(server.NewUserServer(client), nil)

	mux := http.NewServeMux()
	mux.Handle(us.PathPrefix(), middleware.Adapt(us, headers.Username(), headers.UserId(), headers.Email()))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Panic(err)
	}
}