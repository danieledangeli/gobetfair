package httpio

import (
	"testing"
	"net/http"
	"io"
	"bytes"
	"net/url"
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
	if val, ok := header["Accept"]; ok == false || len(val) != 1 || val[0] != "application/json" {
		t.Error("Expected header: Accept: \"application/json\" got", val)
		t.FailNow()
	}

	if val, ok := header["Content-Type"]; ok == false || len(val) != 1 || val[0] != "application/x-www-form-urlencoded" {
		t.Error("Expected header: Accept: \"application/x-www-form-urlencoded\" got", val)
		t.FailNow()
	}

	if val, ok := header["X-Application"]; ok == false || len(val) != 1 || val[0] != "12345AppKey" {
		t.Error("Expected header: Accept: \"application/x-www-form-urlencoded\" \ngot", val)
		t.FailNow()
	}
}

func checkBodyOrFail(body io.ReadCloser, t *testing.T) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	bodyRequest := buf.String()

	if(bodyRequest!= "username=daniele&password=dangeli") {
		t.Error("Expected body: username=daniele&password=dangeli \n Got: ", bodyRequest)
		t.FailNow()
	}
}

func checkMethodOrFail(method string, t *testing.T) {
	if(method != "POST") {
		t.Error("expected HTTP method: POST \n got:", method)
		t.FailNow();
	}
}

func checkUrlOrFail(url *url.URL, t *testing.T) {
	if(url.String() != "www.login.com") {
		t.Error("expected HTTP url: www.login.com/auth \n got:", url.String())
		t.FailNow();
	}
}