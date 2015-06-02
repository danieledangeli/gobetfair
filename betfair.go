package main

import (
	"fmt"
	"github.com/danieledangeli/gobetfair/config"
	"log"
	"github.com/danieledangeli/gobetfair/session"
)

func main() {
	config, err := config.GetConfig("config/conf.yaml")
	session, err:= session.Login(config.Credential, config.LoginEndpoint)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Println(session.Token)
}

