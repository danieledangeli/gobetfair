package main

import (
	"fmt"
	"github.com/danieledangeli/gobetfair/config"
	"log"
)

func main() {
	config, err := config.GetConfig("config/conf_correct.yml")

	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println(config)
}

