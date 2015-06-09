package httpio

import (
	"net/http"
	"bytes"
	"strings"
	"log"
)

type BetfairRequestFactoryInterface interface {
	CreateLoginRequest(username string, password string, applicationKey string, urlEndpoint string) (*http.Request)
}

type BetfairRequestFactory struct {
}

func (brf *BetfairRequestFactory) CreateLoginRequest(username string, password string, applicationKey string, urlEndpoint string) (*http.Request) {
	var buffer bytes.Buffer

	buffer.WriteString("username=")
	buffer.WriteString(username)
	buffer.WriteString("&password=")
	buffer.WriteString(password)

	req, err := http.NewRequest("POST", urlEndpoint, strings.NewReader(buffer.String()))

	if err != nil {
		log.Panic(err)
	}

	req.Header.Add("X-Application", applicationKey)
	req.Header.Add("Accept","application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return req
}
