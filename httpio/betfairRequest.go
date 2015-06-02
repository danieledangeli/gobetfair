package httpio

import (
	"net/http"
	"bytes"
	"strings"
	"log"
)

func CreateLoginRequest(username string, password string, applicationKey string, urlEndpoint string) (*http.Request) {
	var buffer bytes.Buffer
	buffer.WriteString("username=")
	buffer.WriteString(username)
	buffer.WriteString("&password=")
	buffer.WriteString(password)

	req, err := http.NewRequest("POST", urlEndpoint, strings.NewReader(buffer.String()))

	//debug(httputil.DumpRequestOut(req, true))

	if err != nil {
		log.Panic(err)
	}

	req.Header.Add("X-Application:", applicationKey)
	req.Header.Add("Accept","application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return req
}
