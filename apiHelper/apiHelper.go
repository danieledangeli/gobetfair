package apiHelper
import "github.com/danieledangeli/gobetfair/response"

type ApiHelper interface {
	ApiRequest(parameters []string)  (response.Response, error)
}