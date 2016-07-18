## Overview
A simple API wrapper for [Email Hunter](https://emailhunter.co), in Golang.

## License
gohunting is licensed under an MIT License.

## Installation
`go get github.com/dylanjt/gohunting`

## Usage

	package main

	import (
		"github.com/dylanjt/gohunting"
	)

	func main() {
		// prepare you a client
		client := gohunting.Client("your Email Hunter API key")

		// search a domain for emails
		client.Search("asana.com")
	}