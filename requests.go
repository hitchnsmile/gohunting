package gohunting

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Search uses Email Hunter find all the email addresses corresponding to a domain.
// See https://emailhunter.co/api/docs#domain-search for more information.
func (emailHunter *EmailHunter) Search(domain string) (searchResponse *SearchResponse, exception *Exception, err error) {
	searchUrl := emailHunter.BaseUrl + "/search"

	formValues := url.Values{}
	formValues.Set("domain", domain)

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

// Find uses Email Hunter to generate the most likely email address.
// See https://emailhunter.co/api/docs#email-finder for more information.
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

// Verify uses Email Hunter to verify the deliverability of an email address.
// See https://emailhunter.co/api/docs#email-verification for more information.
func (emailHunter *EmailHunter) Verify(email string) (verifyResponse *VerifyResponse, exception *Exception, err error) {
	searchUrl := emailHunter.BaseUrl + "/verify"

	formValues := url.Values{}
	formValues.Set("email", email)

	res, err := emailHunter.sendRequest(formValues, searchUrl)
	if err != nil {
		return verifyResponse, exception, err
	}
	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return verifyResponse, exception, err
	}

	if res.StatusCode != http.StatusOK {
		exception = new(Exception)
		err = json.Unmarshal(responseBody, exception)

		// We aren't checking the error because we don't actually care.
		// It's going to be passed to the client either way.
		return verifyResponse, exception, err
	}

	verifyResponse = new(VerifyResponse)
	err = json.Unmarshal(responseBody, verifyResponse)

	return
}

// Count tells you how many email addresses Email Hunter has for a domain.
// See https://emailhunter.co/api/docs#email-count for more information.
func (emailHunter *EmailHunter) Count(domain string) (countResponse *CountResponse, exception *Exception, err error) {
	searchUrl := emailHunter.BaseUrl + "/email-count"

	formValues := url.Values{}
	formValues.Set("domain", domain)

	res, err := emailHunter.sendRequest(formValues, searchUrl)
	if err != nil {
		return countResponse, exception, err
	}
	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return countResponse, exception, err
	}

	if res.StatusCode != http.StatusOK {
		exception = new(Exception)
		err = json.Unmarshal(responseBody, exception)
		return countResponse, exception, err
	}

	countResponse = new(CountResponse)
	err = json.Unmarshal(responseBody, countResponse)

	return
}

// Account gives you information, like usage, about your account.
// See https://emailhunter.co/api/docs#account for more information.
func (emailHunter *EmailHunter) Account() (accountResponse *AccountResponse, exception *Exception, err error) {
	searchUrl := emailHunter.BaseUrl + "/account"

	formValues := url.Values{}

	res, err := emailHunter.sendRequest(formValues, searchUrl)
	if err != nil {
		return accountResponse, exception, err
	}
	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return accountResponse, exception, err
	}

	if res.StatusCode != http.StatusOK {
		exception = new(Exception)
		err = json.Unmarshal(responseBody, exception)
		return accountResponse, exception, err
	}

	accountResponse = new(AccountResponse)
	err = json.Unmarshal(responseBody, accountResponse)

	return
}
