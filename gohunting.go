// gohunting: but a humble API wrapper for Mail Hunter.
package gohunting

import (
	"net/http"
	"net/url"
	"strings"
)

// EmailHunter to store client info in
type EmailHunter struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
}

// Exception is a representation of a http exception.
type Exception struct {
	Message string `json:"message"`
}

// Client creates a http client and EmailHunter struct
func Client(APIKey string) *EmailHunter {
	baseURL := "https://api.hunter.io/v2"
	HTTPClient := http.DefaultClient

	return &EmailHunter{APIKey, baseURL, HTTPClient}
}

// internal function to actually communicate with Email Hunter
func (emailHunter *EmailHunter) sendRequest(formValues url.Values, emailHunterURL string) (*http.Response, error) {
	formValues.Set("api_key", emailHunter.APIKey)
	req, err := http.NewRequest("GET", emailHunterURL, strings.NewReader(formValues.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := emailHunter.HTTPClient.Do(req)

	return resp, err
}
