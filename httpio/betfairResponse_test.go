package httpio

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestItDecodeResponseByteDataIntoBetfairResponseObject(t *testing.T) {
	expectedDecodedResponse := BetfairLoginResponse{"123", "product", "200", "error"}

	var response BetfairLoginResponse
	var body = []byte("{\"token\": \"123\", \"Product\": \"product\", \"status\": \"200\", \"error\": \"error\"}")
	var decodedResponse = response.Decode(body)

	assert.Equal(t, expectedDecodedResponse, decodedResponse)
}

func TestItHasErrorsIfPresent(t *testing.T) {
	loginResponse := BetfairLoginResponse{"123", "product", "404", "error"}
	assert.True(t, loginResponse.HasError())
}

func TestItDontHaveErrorsIfNotPresent(t *testing.T) {
	loginResponse := BetfairLoginResponse{"123", "product", "404", ""}
	assert.False(t, loginResponse.HasError())
}

func TestGetErrorItReturnError(t *testing.T) {
	loginResponse := BetfairLoginResponse{"123", "product", "", "horror error"}
	assert.Equal(t, "horror error", loginResponse.GetError())
}
