package httpio

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
)

func DoRequest(request *http.Request) BetfairResponse {
	hc := http.Client{}
	response := decodeResponse(hc.Do(request))
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

func debug(data []byte, err error) {
	if err == nil {
		fmt.Printf("%s\n\n", data)
	} else {
		log.Fatalf("%s\n\n", err)
	}
}
