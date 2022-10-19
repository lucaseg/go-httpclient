package examples

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/lucaseg/go-httpclient/gohttp_mock"
)

func TestMain(m *testing.M) {
	gohttp_mock.StartMockServer()
	os.Exit(m.Run())
}

func TestGet(t *testing.T) {

	t.Run("When the request fail", func(t *testing.T) {
		// Given
		gohttp_mock.FlushMocks()
		mock := gohttp_mock.Mock{
			Method:         http.MethodGet,
			Url:            "https://api.github.com",
			Error:          errors.New("request time out"),
			ResponseStatus: 500,
		}
		gohttp_mock.AddMock(mock)
		// When
		endpoints, err := GetEndpoints()

		// Then
		if endpoints != nil {
			t.Error("Expected endpoint equals to nil")
		}

		if err == nil {
			t.Error("Error expected")
		}

		if !strings.Contains(err.Error(), "request time out") {
			t.Error("The error message received is invalid")
		}
	})

	t.Run("When the unmarshal fail", func(t *testing.T) {
		// Given
		gohttp_mock.FlushMocks()
		mock := gohttp_mock.Mock{
			Method: http.MethodGet,
			Url:    "https://api.github.com",
			Error:  errors.New("request time out"),

			ResponseBody: "{" +
				"\"current_user_url\": 123" +
				"}",
			ResponseStatus: 500,
		}
		gohttp_mock.AddMock(mock)
		// When
		endpoints, err := GetEndpoints()

		// Then
		if endpoints != nil {
			t.Error("endpoints expected equals to nil")
		}

		if err == nil {
			t.Error("error expected")
		}

		if !strings.Contains(err.Error(), "request time out") {
			t.Error("error message received is invalid")
		}
	})

	t.Run("Success case", func(t *testing.T) {
		// Given
		gohttp_mock.FlushMocks()
		mock := gohttp_mock.Mock{
			Method: http.MethodGet,
			Url:    "https://api.github.com",
			ResponseBody: "{" +
				"\"current_user_url\" : \"test\"," +
				"\"authorizations_url\": \"test\"," +
				"\"repository_url\": \"test\"" +
				"}",
			ResponseStatus: 200,
		}
		gohttp_mock.AddMock(mock)
		// When
		endpoints, err := GetEndpoints()
		fmt.Println(endpoints)
		fmt.Println(err)
		// Then
		if endpoints == nil {
			t.Error("endpoint expected")
		}
		if err != nil {
			t.Error("no error expected")
		}
	})
}
