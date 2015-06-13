package httpio

import (
	"testing"
	"net/http"
	"io"
	"bytes"
	"net/url"

	"github.com/stretchr/testify/assert"
)

func TestItCreateLoginRequest(t *testing.T) {

	brf := BetfairRequestFactory{};
	http := brf.CreateLoginRequest("daniele", "dangeli", "12345AppKey", "www.login.com")
	checkHeadersOrFail(http.Header, t);
	checkBodyOrFail(http.Body, t);
	checkMethodOrFail(http.Method, t)
	checkUrlOrFail(http.URL, t)

}

func checkHeadersOrFail(header http.Header, t *testing.T) {

	headerExpected := http.Header{}
	headerExpected.Add("Accept", "application/json")
	headerExpected.Add("Content-Type", "application/x-www-form-urlencoded")
	headerExpected.Add("X-Application", "12345AppKey")

	assert.ObjectsAreEqual(headerExpected, header)
}

func checkBodyOrFail(body io.ReadCloser, t *testing.T) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	bodyRequest := buf.String()

	assert.Equal(t, "username=daniele&password=dangeli", bodyRequest)

}

func checkMethodOrFail(method string, t *testing.T) {
	assert.Equal(t, "POST", method)
}

func checkUrlOrFail(url *url.URL, t *testing.T) {
	assert.Equal(t, "www.login.com", url.String())
}