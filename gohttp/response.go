package gohttp

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	status string
	statusCode int
	headers http.Header
	body []byte
}

func (r *Response) Bytes() []byte {
	return r.body
}

func (r *Response) String() string {
	return string(r.body)
}

func (r *Response) UnmarshalJson(target interface{}) error {
	return json.Unmarshal(r.Bytes(), target)
}