package session
import (
	"github.com/danieledangeli/gobetfair/credential"
	"net/http"
	"errors"
	"github.com/danieledangeli/gobetfair/httpio"
)

type SessionServiceInterface interface {
	Login(credential credential.Credential, loginEndpoint string) (Session, error)
}

type SessionService struct {
	brf httpio.BetfairRequestFactoryInterface
	httpio httpio.HttpIOInterface
}

type Session struct {
	Token string
}

func (sS *SessionService) Login(credential credential.Credential, loginEndpoint string) (Session, error) {
	var session Session

	req := sS.brf.CreateLoginRequest(credential.BetfairUsername, credential.BetfairPassword, credential.ApplicationKey, loginEndpoint)
	betfairResponse := doLoginRequest(req,sS.httpio)

	if betfairResponse.HasError() {
		return session, errors.New("Betfair response error: " + betfairResponse.GetError())
	}

	session.Token = string(betfairResponse.Token)
	return session, nil
}

func doLoginRequest(req *http.Request, http httpio.HttpIOInterface) httpio.BetfairLoginResponse {
	betfairResponse := http.DoRequest(req)

	if betfairResponse, ok := betfairResponse.(httpio.BetfairLoginResponse); ok {
		return betfairResponse
	} else {
		panic("error on converting BetfairResponse into BetfairLoginResponseInterface")
	}
}
