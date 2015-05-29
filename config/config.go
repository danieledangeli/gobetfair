package config

import (
	"gopkg.in/yaml.v2"
	"github.com/danieledangeli/gobetfair/credential"
	"io/ioutil"
)

type Config struct {
	ApiEndpoint string `yaml:"apiEndpoint"`
	LoginEndpoint string `yaml:"loginEndpoint"`
	Credential credential.Credential `yaml:"credential"`
}

func GetConfig(filepath string) (Config, error) {

	yamlFileByteArray, err := getYamlFile(filepath)

	if err != nil {
		return Config{},err
	}

	return unmarshalFileByteArrayIntoConfig(yamlFileByteArray)
}

func getYamlFile(filepath string) ([]byte, error) {
	return ioutil.ReadFile(filepath)
}

func unmarshalFileByteArrayIntoConfig(yamlFileByteArray []byte) (Config, error) {
	config := Config{}
	err := yaml.Unmarshal(yamlFileByteArray, &config)
	return config, err
}