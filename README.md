## Overview
A simple API wrapper for [Email Hunter](https://emailhunter.co), in Golang.

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
Each of these methods returns a `response` struct, which is just a representation of the JSON returned by Email Hunter (refer to their [API documentation](https://emailhunter.co/api/docs) for those details); an `exception` struct containing the fields `Status` and `Message`, both of which are populated if an API's response code is not 200; and a standard Go `error` you can check against to make sure your API call was executed unexceptionally.

- `client.Search("stripe.com")` returns all known stripe.com emails
- `client.Find("stripe.com", "Dustin", "Moskovitz")` generates an address for this domain and name based on available data or known pattern
- `client.Verify("steli@close.io")` verifies an email address
- `client.Count("stripe.com")` returns the number of known addresses for stripe.com
- `client.Account()` returns your plan information and usage
