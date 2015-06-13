package httpio

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
)

type HttpIOInterface interface {
	DoRequest(request *http.Request) BetfairResponse
}

type HttpIO struct {
	hc *http.Client
}
func (httpio *HttpIO) DoRequest(request *http.Request) BetfairResponse {
	response := decodeResponse(httpio.hc.Do(request))
	return response
}

func decodeResponse(response *http.Response, err error) BetfairResponse {
	var betfairResponse BetfairLoginResponse

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	return betfairResponse.Decode(contents)
}