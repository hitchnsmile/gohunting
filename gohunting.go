package gohunting

import (
	"net/http"
	"net/url"
	"strings"
)

type EmailHunter struct {
	ApiKey     string
	BaseUrl    string
	HTTPClient *http.Client
}

// Exception is a representation of a http exception.
type Exception struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Create a new Email Hunter http client
func Client(ApiKey string) *EmailHunter {
	baseUrl := "https://api.emailhunter.co/v1"
	HTTPClient := http.DefaultClient

	return &EmailHunter{ApiKey, baseUrl, HTTPClient}
}

// internal function to actually communicate with Email Hunter
func (emailHunter *EmailHunter) sendRequest(formValues url.Values, emailHunterUrl string) (*http.Response, error) {
	formValues.Set("api_key", emailHunter.ApiKey)
	req, err := http.NewRequest("GET", emailHunterUrl, strings.NewReader(formValues.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := emailHunter.HTTPClient
	if client == nil {
		client = http.DefaultClient
	}

	return client.Do(req)
}
