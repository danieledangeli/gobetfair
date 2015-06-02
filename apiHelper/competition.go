package apiHelper

import (
	"github.com/danieledangeli/gobetfair/config"
	"github.com/danieledangeli/gobetfair/response"
	"net/http"
)

type Competition struct {
	Config config.Config
}

func (c Competition) ApiRequest(parameters []string) (response.Response, error) {
	var resp response.Response
	
	resp, err := http.Get("http://example.com/")

	if err != nil {
		return resp, err
	}
	
	return response.ParseResponse(resp)
}
