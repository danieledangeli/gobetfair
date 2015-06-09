package httpio

import (
	"testing"
)

func TestItDecodeResponseByteDataIntoBetfairResponseObject(t *testing.T) {
	expectedDecodedResponse := BetfairLoginResponse{"123", "product", "200", "error"}

	var response BetfairLoginResponse
	var body = []byte("{\"token\": \"123\", \"Product\": \"product\", \"status\": \"200\", \"error\": \"error\"}")
	var decodedResponse = response.Decode(body)

	if decodedResponse != expectedDecodedResponse {
		t.Error("Expected: \n", expectedDecodedResponse, "\n Got: \n",  decodedResponse)
		t.FailNow()
	}
}

func TestItHasErrorsIfPresent(t *testing.T) {
	loginResponse := BetfairLoginResponse{"123", "product", "404", "error"}
	if loginResponse.HasError() == false {
		t.Error("Expected: login Response with errors \n got: \n without")
		t.FailNow()
	}
}

func TestItDontHaveErrorsIfNotPresent(t *testing.T) {
	loginResponse := BetfairLoginResponse{"123", "product", "404", ""}
	if loginResponse.HasError() {
		t.Error("Expected: login Response without errors \n got: \n with errors")
		t.FailNow()
	}
}

func TestGetErrorItReturnError(t *testing.T) {
	loginResponse := BetfairLoginResponse{"123", "product", "", "horror error"}
	if loginResponse.GetError() != "horror error" {
		t.Error("Expected: horror error \n Got: \n", loginResponse.GetError())
		t.FailNow()
	}
}
