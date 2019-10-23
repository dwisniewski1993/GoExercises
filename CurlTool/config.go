package main

import "flag"

var (
	// Command line flags config.
	httpMethod      string
	body            string
	followRedirects bool
	httpHeaders     headers
	saveOutput      bool
	outputFile      string

	// number of redirects followed
	redirectsFollowedCount int
)

const (
	defaultUrlScheme = "http"
	maxRedirects     = 10
)

func init() {
	flag.Var(&httpHeaders, "H", "set HTTP headers")
	flag.StringVar(&httpMethod, "X", "GET", "HTTP method to use")
	flag.StringVar(&body, "d", "", "the body of a POST or PUT request")
	flag.BoolVar(&followRedirects, "L", false, "follow redirects")
	flag.BoolVar(&saveOutput, "O", false, "save body as remote filename")
	flag.StringVar(&outputFile, "o", "", "output file for body")
}
