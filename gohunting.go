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

// Exception is a representation of a Email Hunter exception.
type Exception struct {
	Status   int    `json:"status"`    // HTTP specific error code
	Message  string `json:"message"`   // HTTP error message
	Code     int    `json:"code"`      // Twilio specific error code
	MoreInfo string `json:"more_info"` // Additional info from Twilio
}

// Create a new Email Hunter http client
func Client(ApiKey string) *EmailHunter {
	baseUrl := "https://api.emailhunter.co/v1"
	HTTPClient := http.DefaultClient

	return &EmailHunter{ApiKey, baseUrl, HTTPClient}
}

func (emailHunter *EmailHunter) sendRequest(formValues url.Values, emailHunterUrl string) (*http.Response, error) {
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
