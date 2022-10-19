package gohttp

import (
	"fmt"
)

type Mock struct {
	Method      string
	Url         string
	RequestBody string

	Error          error
	ResponseBody   string
	ResponseStatus int
}

func (m *Mock) getResponse() (*Response, error) {

	if m.Error != nil {
		return nil, m.Error
	}

	return &Response{
		statusCode: m.ResponseStatus,
		status:     fmt.Sprintf("%d %s", m.ResponseStatus, m.ResponseBody),
		body:       []byte(m.ResponseBody),
	}, nil
}
