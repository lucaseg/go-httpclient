package examples

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Repository struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
}

type GitHubError struct {
	StatusCode    int    `json:"-"`
	Message       string `json:"message"`
	Documentation string `json:"documentation_url" `
}

func CreateRepository(request Repository) (*Repository, error) {
	bite, err := json.Marshal(request)
	fmt.Println(string(bite))
	response, err := httpClient.Post("https://api.github.com/user/repos", request)

	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusCreated {
		var gitHubError GitHubError
		if err := response.UnmarshalJson(&gitHubError); err != nil {
			return nil, errors.New("unexpected response from github")
		}
		return nil, errors.New(gitHubError.Message)
	}

	var repository Repository
	if err := response.UnmarshalJson(&repository); err != nil {
		return nil, errors.New("error parsing response: " + err.Error())
	}
	return &repository, nil
}
