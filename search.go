package gohunting

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type SearchResponse struct {
	Status  string `json:"status"`
	Domain  string `json:"domain"`
	Results int    `json:"results"`
	Webmail bool   `json:"webmail"`
	Pattern string `json:"pattern"`
	Offset  int    `json:"offset"`
	Emails  []struct {
		Value      string `json:"value"`
		Type       string `json:"type"`
		Confidence int    `json:"confidence"`
		Sources    []struct {
			Domain       string `json:"domain"`
			Uri          string `json:"uri"`
			Extracted_on string `json:"extract_on"`
		} `json:"sources"`
	} `json:"emails"`
}

// Search uses Email Hunter to send a text message.
// See https://emailhunter.co/api/docs#domain-search for more information.
func (emailHunter *EmailHunter) Search(domain string) (searchResponse *SearchResponse, exception *Exception, err error) {
	searchUrl := emailHunter.BaseUrl + "/search"

	formValues := url.Values{}
	formValues.Set("domain", domain)
	formValues.Set("api_key", emailHunter.ApiKey)

	res, err := emailHunter.sendRequest(formValues, searchUrl)
	if err != nil {
		return searchResponse, exception, err
	}
	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return searchResponse, exception, err
	}

	if res.StatusCode != http.StatusOK {
		exception = new(Exception)
		err = json.Unmarshal(responseBody, exception)

		// We aren't checking the error because we don't actually care.
		// It's going to be passed to the client either way.
		return searchResponse, exception, err
	}

	searchResponse = new(SearchResponse)
	err = json.Unmarshal(responseBody, searchResponse)

	return
}
