## Overview
A simple API wrapper for [Email Hunter](https://emailhunter.co), in Golang.

## License
gohunting is licensed under an MIT License. Due diligence: this is my first public Golang package, and because I wanted it to be up to community standards I modelled it after Sam Freiberg's [gotwillio](https://github.com/sfreiberg/gotwilio), which at the time of writing had a [rating of A by Go Report Card](https://goreportcard.com/report/github.com/sfreiberg/gotwilio).

## Installation
`go get github.com/dylanjt/gohunting`

## Usage

	package main

	import (
		"fmt"
		"github.com/dylanjt/gohunting"
	)

	func main() {
		// prepare you a client
		client := gohunting.Client("API_KEY")

		// get all known addresses for a domain
		searchResponse, exception, err := client.Search("stripe.com")
		if err != nil {
			fmt.Println(err)
			fmt.Println(exception.Message)
		}
		fmt.Println(searchResponse)

		// generate an address for a domain and first & last name
		findResponse, _, _ := client.Find("stripe.com", "Dustin", "Moskovitz")
		fmt.Println(findResponse)

		// verify an address
		verifyResponse, _, _ := client.Verify("steli@close.io")
		fmt.Println(verifyResponse)

		// count all addresses for a domain
		countResponse, _, _ := client.Count("stripe.com")
		fmt.Println(countResponse)

		// information about your account
		accountResponse, _, _ := client.Account()
		fmt.Println(accountResponse)
	}

In the above examples, `response` is a struct representation of the JSON returned by Email Hunter (refer to their [API documentation](https://emailhunter.co/api/docs) for those details). `exception` is a struct containing the fields `Status` and `Message`, both of which are populated if an API's response code is not 200. `error` is a standard Go error object you can check against to make sure your API call was executed unexceptionally.
