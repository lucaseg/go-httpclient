package gohttp

import "time"

type builder struct {
	maxIdleConnections int
	connectionTimeout  time.Duration
	responseTimeout    time.Duration
	disableTimeout     bool
}

type ClientBuilder interface {
	SetMaxIdleConnections(connections int) ClientBuilder
	SetConnectionTimeout(timeout time.Duration) ClientBuilder
	SetResponseTimeout(timeout time.Duration) ClientBuilder
	DisableTimeout(b bool) ClientBuilder
}

func NewBuilder() ClientBuilder {
	return &builder{}
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
