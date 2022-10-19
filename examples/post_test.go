package examples

import (
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/lucaseg/go-httpclient/gohttp"
)

func TestPostGitHub(t *testing.T) {
	t.Run("Error creating repo", func(t *testing.T) {
		// Initialization
		gohttp.StartMockServer()
		gohttp.FlushMocks()
		gohttp.AddMock(gohttp.Mock{
			Method:      http.MethodPost,
			Url:         "https://api.github.com/user/repos",
			RequestBody: "{\"name\":\"test\",\"description\":\"\",\"private\":true}",

			Error: errors.New("request timeout"),
		})
		repository := Repository{
			Name:        "test",
			Description: "",
			Private:     true,
		}
		// Execution
		repo, err := CreateRepository(repository)

		// Validation
		if repo != nil {
			t.Error("no repository expected")
		}
		if err == nil {
			t.Error("error expected")
		}

	})
	t.Run("Creation success", func(t *testing.T) {
		// Initialization
		gohttp.StartMockServer()
		gohttp.FlushMocks()
		gohttp.AddMock(gohttp.Mock{
			Method:         http.MethodPost,
			Url:            "https://api.github.com/user/repos",
			RequestBody:    "{\"name\":\"test\",\"description\":\"\",\"private\":true}",
			ResponseStatus: 201,
			ResponseBody:   "{\"name\":\"test\",\"description\":\"\",\"private\":true}",
		})
		repository := Repository{
			Name:        "test",
			Description: "",
			Private:     true,
		}
		// Execution
		repo, err := CreateRepository(repository)

		// Validation
		if err != nil {
			t.Error("error unexpected")
		}
		if repo == nil {
			t.Error("response distinct nil expected")
		}
		if repo.Name != "test" {
			t.Error("response name unexpected")
		}

	})
	t.Run("Status Code != 201", func(t *testing.T) {
		gohttp.StartMockServer()
		gohttp.FlushMocks()
		gohttp.AddMock(gohttp.Mock{
			Method:         http.MethodPost,
			Url:            "https://api.github.com/user/repos",
			RequestBody:    "{\"name\":\"test\",\"description\":\"\",\"private\":true}",
			ResponseStatus: 400,
			ResponseBody:   "{\"message\":\"bad_request\",\"documentation_url\":\"documentation.com.ar\"}",
		})
		repository := Repository{
			Name:        "test",
			Description: "",
			Private:     true,
		}
		// Execution
		repo, err := CreateRepository(repository)

		// Validation
		if repo != nil {
			t.Error("unexpected repo value")
		}
		if err == nil {
			t.Error("unexpected error value")
		}
		if !strings.Contains(err.Error(), "bad_request") {
			t.Error("unexpected error message")
		}
	})
	t.Run("Status Code != 201 and unexpected response of error", func(t *testing.T) {
		gohttp.StartMockServer()
		gohttp.FlushMocks()
		gohttp.AddMock(gohttp.Mock{
			Method:         http.MethodPost,
			Url:            "https://api.github.com/user/repos",
			RequestBody:    "{\"name\":\"test\",\"description\":\"\",\"private\":true}",
			ResponseStatus: 400,
			ResponseBody:   "{\"message\":123,\"documentation_url\":\"documentation.com.ar\"}",
		})
		repository := Repository{
			Name:        "test",
			Description: "",
			Private:     true,
		}
		// Execution
		repo, err := CreateRepository(repository)

		// Validation
		if repo != nil {
			t.Error("unexpected repo value")
		}
		if err == nil {
			t.Error("unexpected error value")
		}
		if !strings.Contains(err.Error(), "unexpected response from github") {
			t.Error("unexpected error message")
		}
	})

	t.Run("Unmarshall response success fail", func(t *testing.T) {
		// Initialization
		gohttp.StartMockServer()
		gohttp.FlushMocks()
		gohttp.AddMock(gohttp.Mock{
			Method:         http.MethodPost,
			Url:            "https://api.github.com/user/repos",
			RequestBody:    "{\"name\":\"test\",\"description\":\"\",\"private\":true}",
			ResponseStatus: 201,
			ResponseBody:   "{\"name\":123,\"description\":\"\",\"private\":true}",
		})
		repository := Repository{
			Name:        "test",
			Description: "",
			Private:     true,
		}
		// Execution
		repo, err := CreateRepository(repository)

		// Validation
		// Validation
		if repo != nil {
			t.Error("unexpected repo value")
		}
		if err == nil {
			t.Error("unexpected error value")
		}
		if !strings.Contains(err.Error(), "error parsing response: ") {
			t.Error("unexpected error message")
		}

	})
}
