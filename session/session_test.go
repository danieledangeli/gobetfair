package session

import (
	"testing"
	"net/http"
	"github.com/danieledangeli/gobetfair/httpio"
	"github.com/danieledangeli/gobetfair/credential"
	"strings"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/assert"
)

func TestItLoginAndCreatesAToken(t *testing.T) {
	req, _ := http.NewRequest("POST", "www.test.it", strings.NewReader("body"))
	resp := httpio.BetfairLoginResponse{"superSecret", "", "200", ""}

	brf := new(MockBetfairRequestFactory)
	brf.On("CreateLoginRequest", "username", "password", "key", "www.login.com").Return(req)

	httpIO := new(MockHttpIo)
	httpIO.On("DoRequest", req).Return(resp)

	sessionService := SessionService{brf: brf, httpio: httpIO}

	credential := credential.Credential{"username", "password", "key"}
	session, err := sessionService.Login(credential, "www.login.com")

	assert.Nil(t, err)
	assert.Equal(t, "superSecret", session.Token)
}

func TestItReturnErrorIdLoginResponseHasErrors(t *testing.T) {
	req, _ := http.NewRequest("POST", "www.test.it", strings.NewReader("body"))
	resp := httpio.BetfairLoginResponse{"", "", "400", "an error"}

	brf := new(MockBetfairRequestFactory)
	brf.On("CreateLoginRequest", "username", "password", "key", "www.login.com").Return(req)

	httpIO := new(MockHttpIo)
	httpIO.On("DoRequest", req).Return(resp)

	sessionService := SessionService{brf: brf, httpio: httpIO}

	credential := credential.Credential{"username", "password", "key"}
	_, err := sessionService.Login(credential, "www.login.com")

	assert.NotNil(t, err)
	assert.Equal(t, "Betfair response error: an error", err.Error())
}

type MockHttpIo struct{
	mock.Mock
}

func (m *MockHttpIo) DoRequest(request *http.Request) httpio.BetfairResponse {
	args := m.Called(request)
	return args.Get(0).(httpio.BetfairResponse)

}

type MockBetfairRequestFactory struct {
	mock.Mock
}

func (s *MockBetfairRequestFactory) CreateLoginRequest(username string, password string, applicationKey string, urlEndpoint string)(*http.Request) {
	args := s.Called(username, password, applicationKey, urlEndpoint)
	return args.Get(0).(*http.Request)
}

