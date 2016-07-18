package gohunting

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (emailHunter *EmailHunter) Find(domain, firstName, lastName string) (findResponse *FindResponse, exception *Exception, err error) {
	searchUrl := emailHunter.BaseUrl + "/generate"

	formValues := url.Values{}
	formValues.Set("domain", domain)
	formValues.Set("first_name", firstName)
	formValues.Set("last_name", lastName)

	res, err := emailHunter.sendRequest(formValues, searchUrl)
	if err != nil {
		return findResponse, exception, err
	}
	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return findResponse, exception, err
	}

	if res.StatusCode != http.StatusOK {
		exception = new(Exception)
		err = json.Unmarshal(responseBody, exception)

		// We aren't checking the error because we don't actually care.
		// It's going to be passed to the client either way.
		return findResponse, exception, err
	}

	findResponse = new(FindResponse)
	err = json.Unmarshal(responseBody, findResponse)

	return
}
