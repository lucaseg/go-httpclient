package gohttp

import (
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {
	// initializacion
	client := httpClient{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("X-Auth-Token", "token")
	client.Headers = commonHeaders
	// execution
	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Header-Test", "test")
	fullHeaders := client.getHeaders(requestHeaders)

	// validation
	if fullHeaders.Get("X-Header-Test") != "test" {
		t.Error("The header expected is invalid")
	}

	if fullHeaders.Get("X-Auth-Token") != "token" {
		t.Error("The header expected is invalid")
	}
}

func TestGetRequestBody(t *testing.T) {
	client := httpClient{}

	t.Run("When the body is nil", func(t *testing.T) {
		client.getRequestBody("application/json", nil)
	})

	t.Run("When the content type is json", func(t *testing.T) {
		body := []string{"one", "two"}
		requestBody, err := client.getRequestBody("application/json", body)
		if err != nil {
			t.Error("Not error expected.")
		}
		if string(requestBody) != `["one","two"]` {
			t.Error("Error getting request body")
		}
	})

	t.Run("When the content type is xml", func(t *testing.T) {
		body := []string{"one", "two"}
		requestBody, err := client.getRequestBody("application/xml", body)
		if err != nil {
			t.Error("Not error expected.")
		}
		if string(requestBody) != `<string>one</string><string>two</string>` {
			t.Error("Error getting request body")
		}

	})
	t.Run("When the content type is json default", func(t *testing.T) {
		body := []string{"one", "two"}
		requestBody, err := client.getRequestBody("", body)
		if err != nil {
			t.Error("Not error expected.")
		}
		if string(requestBody) != `["one","two"]` {
			t.Error("Error getting request body")
		}
	})
}
