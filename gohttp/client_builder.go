package gohttp

import (
	"net/http"
	"time"
)

type builder struct {
	headers            http.Header
	maxIdleConnections int
	connectionTimeout  time.Duration
	responseTimeout    time.Duration
	disableTimeout     bool
	client             *http.Client
	userAgent          string
}

type ClientBuilder interface {
	SetHeaders(headers http.Header) ClientBuilder
	SetMaxIdleConnections(connections int) ClientBuilder
	SetConnectionTimeout(timeout time.Duration) ClientBuilder
	SetResponseTimeout(timeout time.Duration) ClientBuilder
	DisableTimeout(b bool) ClientBuilder
	SetHttpClient(c *http.Client) ClientBuilder
	SetUserAgent(user string) ClientBuilder
	Build() Client
}

func NewBuilder() ClientBuilder {
	return &builder{}
}

func (c *builder) Build() Client {
	client := &httpClient{
		builder: c,
	}
	return client
}

func (c *builder) SetHeaders(headers http.Header) ClientBuilder {
	c.headers = headers
	return c
}

func (c *builder) SetMaxIdleConnections(connections int) ClientBuilder {
	c.maxIdleConnections = connections
	return c
}
func (c *builder) SetConnectionTimeout(timeout time.Duration) ClientBuilder {
	c.connectionTimeout = timeout
	return c
}
func (c *builder) SetResponseTimeout(timeout time.Duration) ClientBuilder {
	c.responseTimeout = timeout
	return c
}
func (c *builder) DisableTimeout(b bool) ClientBuilder {
	c.disableTimeout = b
	return c
}

func (c *builder) SetHttpClient(client *http.Client) ClientBuilder {
	c.client = client
	return c
}

func (c *builder) SetUserAgent(user string) ClientBuilder {
	c.userAgent = user
	return c
}
