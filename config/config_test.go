package config

import (
	"testing"
	"github.com/danieledangeli/gobetfair/credential"
)

func TestItParseYamlFile(t *testing.T) {
	expectedConfig  := Config{"www.google.com", "www.login.com", credential.Credential{"erlangb88", "annarita05011988", "appKey"}}
	config, err := GetConfig("testFixtures/conf_correct.yml")

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if config != expectedConfig {
		t.Error("The config are different: \nExpected Config\n",expectedConfig, "\nActual:\n",config)
		t.FailNow()
	}
}

func TestItRaiseAnErrorIfTheFileIsNotFound(t *testing.T) {
	config, err := GetConfig("testFixtures/conf_correct_not_found.yml")

	if err == nil {
		t.Error("Expected File Not Fond got:", config)
		t.FailNow()
	}
}