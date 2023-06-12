package main

import (
	"log"

	"github.com/tonygits/test-lending-svc/services"
)

func main() {
	runGinServer()
}

func runGinServer() {
	address := ":8099"
	server, err := services.NewRouter()
	if err != nil {
		log.Fatal("cannot create server", err)
	}

	err = server.Start(address)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
