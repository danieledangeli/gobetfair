package session

import (
	"testing"
	"net/http"
	"github.com/danieledangeli/gobetfair/httpio"
	"github.com/danieledangeli/gobetfair/credential"
	"strings"
)

func TestItLoginAndCreatesAToken(t *testing.T) {
	httpIO := &MockHttpIo{}

	httpIO.DoRequestFunc = (func(request *http.Request) httpio.BetfairResponse {
		betfairResponse := httpio.BetfairLoginResponse{"superSecret", "150", "200", ""}
		return betfairResponse;
	})

	brf := &MockBetfairRequestFactory{}
	brf.CreateLoginRequestFunc = (func(username string, password string, applicationKey string, urlEndpoint string) (*http.Request) {
		req, _ := http.NewRequest("POST", "www.test.it", strings.NewReader("body"))
		return req
	})

	sessionService := SessionService{brf: brf, httpio: httpIO}

	credential := credential.Credential{"username", "password", "key"}
	session, err := sessionService.Login(credential, "wwwwww.login.com")

	if err != nil {
		t.Error("not expected error", err)
		t.FailNow()
	}

	if session.Token != "superSecret" {
		t.Error("not expected session", session)
		t.FailNow()
	}
}

func TestItReturnErrorIdLoginResponseHasErrors(t *testing.T) {
	httpIO := &MockHttpIo{}

	httpIO.DoRequestFunc = (func(request *http.Request) httpio.BetfairResponse {
		betfairResponse := httpio.BetfairLoginResponse{"", "", "400", "an error"}
		return betfairResponse;
	})

	brf := &MockBetfairRequestFactory{}
	brf.CreateLoginRequestFunc = (func(username string, password string, applicationKey string, urlEndpoint string) (*http.Request) {
		req, _ := http.NewRequest("POST", "www.test.it", strings.NewReader("body"))
		return req
	})

	sessionService := SessionService{brf: brf, httpio: httpIO}

	credential := credential.Credential{"username", "password", "key"}
	session, err := sessionService.Login(credential, "wwww.login.com")

	if err == nil {
		t.Error("Expected return error, got session", session)
		t.FailNow()
	}

	if(err.Error() != "Betfair response error: an error") {
		t.Error("Assert response Error: \nExpected Error message\n", "Betfair response error: an error", "\nActual:\n",err.Error())
		t.FailNow()
	}
}

type MockHttpIo struct {
	DoRequestFunc func(request *http.Request) httpio.BetfairResponse
}

type MockBetfairRequestFactory struct {
	CreateLoginRequestFunc func(username string, password string, applicationKey string, urlEndpoint string) (*http.Request)
}

func (s *MockBetfairRequestFactory) CreateLoginRequest(username string, password string, applicationKey string, urlEndpoint string)(*http.Request) {
	return s.CreateLoginRequestFunc(username, password, applicationKey, urlEndpoint)
}

func (h *MockHttpIo) DoRequest(request *http.Request) httpio.BetfairResponse {
	return h.DoRequestFunc(request)
}
