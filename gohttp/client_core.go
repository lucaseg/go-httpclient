package gohttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/lucaseg/go-httpclient/core"
	"github.com/lucaseg/go-httpclient/gohttp_mock"
	"github.com/lucaseg/go-httpclient/gomime"
)

const (
	defaultMaxIdleConnections = 5
	defaultResponseTimeout    = 5 * time.Second
	defaultConnectionTimeout  = 1 * time.Second
)

func (c *httpClient) getRequestBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}
	switch strings.ToLower(contentType) {
	case gomime.ContentTypeJson:
		return json.Marshal(body)
	case gomime.ContentTypeXml:
		return xml.Marshal(body)
	default:
		return json.Marshal(body)
	}
}

func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*core.Response, error) {
	fullHeaders := c.getHeaders(headers)

	requestBody, err := c.getRequestBody(fullHeaders.Get(gomime.HeaderContentType), body)

	if err != nil {
		return nil, err
	}

	if mock := gohttp_mock.GetMock(method, url, string(requestBody)); mock != nil {
		return mock.GetResponse()
	}

	// TODO: revisar el control de errores de este metodo
	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.New("unable to create a new request")
	}
	request.Header = fullHeaders

	client := c.getHttpClient()
	response, err := client.Do(request)
	if err != nil {
		return nil, errors.New("error trying to do request")
	}

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("error trying to read response body")
	}

	finalResponse := core.Response{
		StatusCode: response.StatusCode,
		Headers:    response.Header,
		Body:       responseBody,
		Status:     response.Status,
	}
	return &finalResponse, nil
}

func (c *httpClient) getHttpClient() *http.Client {
	c.clientOne.Do(func() {
		if c.builder.client != nil {
			c.client = c.builder.client
			return
		}
		c.client = &http.Client{
			Timeout: c.getConnectionTimeout() + c.getResponseTimeout(),
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   c.getMaxIdleConnections(),
				ResponseHeaderTimeout: c.getResponseTimeout(),
				DialContext: (&net.Dialer{
					Timeout: c.getConnectionTimeout(),
				}).DialContext,
			},
		}
	})
	return c.client
}

func (c *httpClient) getMaxIdleConnections() int {
	if c.builder.maxIdleConnections > 0 {
		return c.builder.maxIdleConnections
	}
	return defaultMaxIdleConnections
}

func (c *httpClient) getResponseTimeout() time.Duration {
	if c.builder.responseTimeout > 0 {
		return c.builder.responseTimeout
	}
	if c.builder.disableTimeout {
		return 0
	}
	return defaultResponseTimeout
}

func (c *httpClient) getConnectionTimeout() time.Duration {
	if c.builder.responseTimeout > 0 {
		return c.builder.responseTimeout
	}
	if c.builder.disableTimeout {
		return 0
	}
	return defaultConnectionTimeout
}

func (c *httpClient) getHeaders(headers http.Header) http.Header {
	result := make(http.Header)

	// Headers from httpclient
	for header, value := range c.builder.headers {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	// Custom headers
	for header, value := range headers {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	if c.builder.userAgent != "" {
		if result.Get(gomime.HeaderUserAgent) != "" {
			return result
		}
		result.Set(gomime.HeaderUserAgent, c.builder.userAgent)
	}
	return result
}
