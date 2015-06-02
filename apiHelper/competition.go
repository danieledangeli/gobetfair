package apiHelper

import (
	"github.com/danieledangeli/gobetfair/config"
	"github.com/danieledangeli/gobetfair/response"
	"net/http"
	"github.com/danieledangeli/gobetfair/session"
)

type Competition struct {
	Config config.Config
	Session session.Session
}

func (c Competition) ApiRequest(parameters []string) (response.Response, error) {
	var resp response.Response
	
	httpResponse, err := http.Get(c.Config.ApiEndpoint)

	if err != nil {
		return resp, err
	}
	
	return response.ParseResponse(httpResponse)
}
