package gohttp

import (
	"net/http"
	"sync"
)

type httpClient struct {
	builder   *builder
	client    *http.Client
	clientOne sync.Once
}

type Client interface {
	Get(url string, headers ...http.Header) (*Response, error)
	Post(url string, body interface{}, headers ...http.Header) (*Response, error)
	Put(url string, body interface{}, headers ...http.Header) (*Response, error)
	Delete(url string, headers ...http.Header) (*Response, error)
	Patch(url string, headers ...http.Header) (*Response, error)
}

func getHeaders(headers ...http.Header) http.Header {
	if len(headers) > 0 {
		return headers[0]
	}
	return http.Header{}
}

func (c *httpClient) Get(url string, headers ...http.Header) (*Response, error) {

	return c.do(http.MethodGet, url, getHeaders(headers...))
}
func (c *httpClient) Post(url string, body interface{}, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodPost, url, getHeaders(headers...), body)
}
func (c *httpClient) Put(url string, body interface{}, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodPut, url, getHeaders(headers...), body)
}
func (c *httpClient) Delete(url string, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodDelete, url, getHeaders(headers...), nil)
}
func (c *httpClient) Patch(url string, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodPatch, url, getHeaders(headers...), nil)
}
