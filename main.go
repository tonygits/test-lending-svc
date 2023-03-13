package main

import (
	"github.com/tonygits/test-lending-svc/services"
	"log"
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
		log.Fatal("cannot connect server", err)
	}
}
