## Overview
A simple API wrapper for [Email Hunter](https://emailhunter.co), in Golang.

## License
gohunting is licensed under an MIT License. Due diligence: this is my first public Golang package, and because I wanted it to be up to community standards I modelled it after Sam Freiberg's [gotwillio](https://github.com/sfreiberg/gotwilio), which at the time of writing had a [rating of A by Go Report Card](https://goreportcard.com/report/github.com/sfreiberg/gotwilio).

## Installation
`go get github.com/dylanjt/gohunting`

## Usage

	package main

	import "github.com/dylanjt/gohunting"

	func main() {
		// prepare you a client
		client := gohunting.Client("API_KEY")

		var response, exception, error

		// get all known addresses for a domain
		response, exception, error = client.Search("stripe.com")

		// generate an address for a domain and first & last name
		response, exception, error = client.Find("stripe.com", "Dustin", "Moskovitz")

		// verify an address
		response, exception, error = client.Verify("steli@close.io")

		// count all addresses for a domain
		response, exception, error = client.Count("stripe.com")

		// information about your account
		response, exception, error = client.Account()
	}

In the above examples, `response` is a struct representation of the JSON returned by Email Hunter (refer to their [API documentation](https://emailhunter.co/api/docs) for those details). `exception` is a struct containing the fields `Status` and `Message`, both of which are populated if an API's response code is not 200. `error` is a standard Go error object you can check against to make sure your API call was executed unexceptionally.