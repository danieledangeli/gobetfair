package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"github.com/danieledangeli/gobetfair/credential"
)

var data = `
apiEndpoint: www.google.com
loginEndpoint: www.login.com
credential:
  betfairUsername: username
  betfairPassword: password
`
type Config struct {
	ApiEndpoint string `yaml:"apiEndpoint"`
	LoginEndpoint string `yaml:"loginEndpoint"`
	Credential credential.Credential `yaml:"credential"`
}

func GetConfig() (Config) {
	config := Config{}
	err := yaml.Unmarshal([]byte(data), &config)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return config
}