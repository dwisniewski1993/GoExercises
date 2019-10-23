package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Printf("Curl Tool Application")
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		flag.Usage()
		os.Exit(1)
	}

	if (httpMethod == http.MethodPost || httpMethod == http.MethodPut) && body == "" {
		log.Fatal("httpMethod: must supply body using -d flag")
	}

	url := parseUrl(args[0])

	c := config{
		httpHeaders:     httpHeaders,
		body:            body,
		followRedirects: followRedirects,
		saveOutput:      saveOutput,
		outputFile:      outputFile,
	}
	ctx := context.Background()

	performRequest(ctx, url, &c)
}
