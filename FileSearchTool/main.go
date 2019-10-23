package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"regexp"
)

func main() {
	log.Printf("File Search Tool Application")

	flag.Parse()

	if matchPattern != "" {
		matchRegex = regexp.MustCompile(matchPattern)
	}

	if ignorePattern != "" {
		ignoreRegexp = regexp.MustCompile(ignorePattern)
	} else {
		ignoreRegexp = regexp.MustCompile(defaultIgnorePattern)
	}

	if searchLimit > 0 {
		searchLimitCount = searchLimit
	}

	basePath := defaultBasePath
	if flag.NArg() > 0 {
		basePath = filepath.FromSlash(flag.Arg(0))
	}

	results := performSearch(basePath)

	for result := range results {
		fmt.Println(result)
	}
}
