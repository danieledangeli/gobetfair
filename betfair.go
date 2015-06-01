package main

import (
	"fmt"
	"github.com/danieledangeli/gobetfair/config"
	"log"
	"github.com/danieledangeli/gobetfair/apiHelper"
)

func main() {
	config, err := config.GetConfig("config/conf.yaml.dist")

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	competition := apiHelper.Competition{config}

	params := []string{"leto", "paul", "teg"}
	response, err := competition.ApiRequest(params)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Println(response)
}

