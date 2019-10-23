package main

import (
	"errors"
	"flag"
	"math"
	"regexp"
)

var (
	// Konfiguracja lini komend
	ignorePattern        string
	matchPattern         string
	searchLimit          int64
	matchDirectoriesOnly bool

	ignoreRegexp            *regexp.Regexp
	matchRegex              *regexp.Regexp
	searchLimitCount        int64 = math.MaxInt64
	searchLimitReachedError       = errors.New("maximum file limit reached")
)

const defaultBasePath string = "."
const defaultIgnorePattern = `^(\.git|\.hg|\.svn)$`

func init() {
	flag.StringVar(&ignorePattern, "i", "", "Ignore pattern")
	flag.StringVar(&matchPattern, "m", "", "Match file/folder pattern")
	flag.Int64Var(&searchLimit, "l", -1, "Max files limit")
	flag.BoolVar(&matchDirectoriesOnly, "d", false, "Search directories only")
}
