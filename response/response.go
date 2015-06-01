package response
import (
	"net/http"
	"io/ioutil"
)

type Response struct {
	Body string
	Status int
}

func ParseResponse(response *http.Response) (Response, error) {
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Response{}, err
	}
	return Response{string(contents), 200}, nil
}