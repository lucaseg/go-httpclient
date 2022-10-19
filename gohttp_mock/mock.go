package gohttp_mock

import (
	"fmt"

	"github.com/lucaseg/go-httpclient/core"
)

type Mock struct {
	Method      string
	Url         string
	RequestBody string

	Error          error
	ResponseBody   string
	ResponseStatus int
}

func (m *Mock) GetResponse() (*core.Response, error) {

	if m.Error != nil {
		return nil, m.Error
	}

	return &core.Response{
		StatusCode: m.ResponseStatus,
		Status:     fmt.Sprintf("%d %s", m.ResponseStatus, m.ResponseBody),
		Body:       []byte(m.ResponseBody),
	}, nil
}
