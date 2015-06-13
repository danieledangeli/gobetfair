package config

import (
	"testing"
	"github.com/danieledangeli/gobetfair/credential"
	"github.com/stretchr/testify/assert"
)

func TestItParseYamlFile(t *testing.T) {
	expectedConfig  := Config{"www.google.com", "www.login.com", credential.Credential{"username", "password", "appKey"}}
	config, err := GetConfig("testFixtures/conf_correct.yml")

	assert.Nil(t, err)
	assert.Equal(t, expectedConfig, config)
}

func TestItRaiseAnErrorIfTheFileIsNotFound(t *testing.T) {
	_, err := GetConfig("testFixtures/conf_correct_not_found.yml")
	assert.NotNil(t, err)
}