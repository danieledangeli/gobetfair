package httpio

import (
	"encoding/json"
)

type BetfairResponse interface {
	Decode(data []byte) BetfairResponse
	HasError() bool
	GetError() string
	GetData() interface{}
}

type BetfairLoginResponse struct {
	Token string
	product string
	status string
	error string
}
func (r BetfairLoginResponse) Decode(data []byte) (betfairResponse BetfairResponse) {
	json.Unmarshal(data, &r)
	return r
}

func (r BetfairLoginResponse) GetError() (error string) {
	return r.error
}

func (r BetfairLoginResponse) HasError() (hasError bool) {
	return len(r.error) > 0
}

func (r BetfairLoginResponse) GetData() (data interface{}) {
	return r.Token
}



