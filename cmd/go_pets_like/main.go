package main

import (
	"github.com/max-rodziyevsky/go-pets-like/api"
	"github.com/max-rodziyevsky/go-pets-like/configs"
	"log"
)

func main() {
	cfg, err := configs.NewConfig()
	if err != nil {
		panic(err)
	}

	serverAddress, err := cfg.GetServerAddress()
	if err != nil {
		log.Fatal(err)
	}

	if err := api.NewHttpServer(serverAddress).Start(); err != nil {
		log.Fatal(err)
	}
}
