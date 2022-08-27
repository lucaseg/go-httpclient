package gohttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	defaultMaxIdleConnections = 5
	defaultResponseTimeout    = 5
	defaultConnectionTimeout  = 5
)

func (c *httpClient) getRequestBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}
	switch strings.ToLower(contentType) {
	case "application/json":
		return json.Marshal(body)
	case "application/xml":
		return xml.Marshal(body)
	default:
		return json.Marshal(body)
	}
}

func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*http.Response, error) {
	fullHeaders := c.getHeaders(headers)

	requestBody, err := c.getRequestBody(fullHeaders.Get("Content-type"), body)

	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	client := c.getHttpClient()
	response, err := client.Do(request)

	if err != nil {
		return nil, errors.New("unable to create a new request")
	}

	request.Header = fullHeaders

	return response, nil
}

func (c *httpClient) getHttpClient() *http.Client {
	if c.client != nil {
		return c.client
	}
	client := http.Client{
		Timeout: c.getConnectionTimeout() + c.getResponseTimeout(),
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   c.maxIdleConnections,
			ResponseHeaderTimeout: c.responseTimeout * time.Second,
			DialContext: net.Dialer{
				Timeout: c.connectionTimeout * time.Second,
			}.DialContext,
		},
	}
	return &client
}

func (c *httpClient) getMaxIdleConnections() int {
	if c.maxIdleConnections > 0 {
		return c.maxIdleConnections
	}
	return defaultMaxIdleConnections
}

func (c *httpClient) getResponseTimeout() time.Duration {
	if c.responseTimeout > 0 {
		return c.responseTimeout
	}
	if c.disableTimeout {
		return 0
	}
	return defaultResponseTimeout
}

func (c *httpClient) getConnectionTimeout() time.Duration {
	if c.responseTimeout > 0 {
		return c.responseTimeout
	}
	if c.disableTimeout {
		return 0
	}
	return defaultConnectionTimeout
}

func (c *httpClient) getHeaders(headers http.Header) http.Header {
	result := make(http.Header)

	for header, value := range c.Headers {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	for header, value := range headers {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	return result
}
