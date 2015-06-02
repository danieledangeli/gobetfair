package session
import (
	"github.com/danieledangeli/gobetfair/credential"
	"net/http"
	"errors"
	"github.com/danieledangeli/gobetfair/httpio"
)

type Session struct {
	Token string
}

func Login(credential credential.Credential, loginEndpoint string) (Session, error) {
	var session Session

	req := httpio.CreateLoginRequest(credential.BetfairUsername, credential.BetfairPassword, credential.ApplicationKey, loginEndpoint)
	betfairResponse := doLoginRequest(req)

	if betfairResponse.HasError() {
		return session, errors.New("Betfair response error" + betfairResponse.GetError())
	}

	session.Token = string(betfairResponse.Token)
	return session, nil
}

func doLoginRequest(req *http.Request) httpio.BetfairLoginResponse {
	betfairResponse := httpio.DoRequest(req)

	if betfairLoginResponse, ok := betfairResponse.(httpio.BetfairLoginResponse); ok {
		return betfairLoginResponse
	} else {
		panic("error on converting betfairResponse into BetfairLoginResponse")
	}
}
