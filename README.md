## Overview
A simple API wrapper for [Email Hunter](https://emailhunter.co), in Go.

## License
gohunting is available for use under The MIT License.

## Install
`go get github.com/dylanjt/gohunting`

## Usage

### Import
`import github.com/dylanjt/gohunting`

### Prepare a client
A client is needed for all requests to Email Hunter:

	client := gohunting.Client("API_KEY")
	
### Methods
Each of these methods returns a `response` struct, which is just a representation of the JSON returned by Email Hunter (refer to their [API documentation](https://emailhunter.co/api/docs) for those details), and an `error` that is the body message from Email Hunter if an error occurred (i.e. "invalid API key").

- `client.Search("stripe.com")` returns all known stripe.com emails
- `client.Find("stripe.com", "Dustin", "Moskovitz")` generates an address for this domain and name based on available data or known pattern
- `client.Verify("steli@close.io")` verifies an email address
- `client.Count("stripe.com")` returns the number of known addresses for stripe.com
- `client.Account()` returns your plan information and usage

### Docs
[godoc.org/github.com/dylanjt/gohunting](https://godoc.org/github.com/dylanjt/gohunting)

### Full example

	package main
	
	import (
		"fmt"
		"github.com/dylanjt/gohunting"
	)
	
	func main() {
		client := gohunting.Client("API_KEY")

		// get all known addresses for a domain
		searchResponse, err := client.Search("stripe.com")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(searchResponse)

		// generate an address for a domain and first & last name
		findResponse, _ := client.Find("stripe.com", "Dustin", "Moskovitz")
		fmt.Println(findResponse)

		// verify an address
		verifyResponse, _ := client.Verify("steli@close.io")
		fmt.Println(verifyResponse)

		// count all addresses for a domain
		countResponse, _ := client.Count("stripe.com")
		fmt.Println(countResponse)

		// information about your account
		accountResponse, _ := client.Account()
		fmt.Println(accountResponse)
	}
