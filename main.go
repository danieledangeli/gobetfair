package main

import (
	"fmt"
	"github.com/danieledangeli/gobetfair/config"
)

func main() {
	s := config.GetConfig()
	fmt.Println(s.Credential.BetfairPassword)
}

