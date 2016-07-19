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

In the above examples, the `response` variables are struct representations of the JSON returned by Email Hunter (refer to their [API documentation](https://emailhunter.co/api/docs) for those details). `error` is the body message from Email Hunter when an error occurs (i.e. "invalid API key").
