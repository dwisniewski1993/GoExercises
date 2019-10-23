package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func performSearch(basePath string) chan string {
	q := make(chan string, 32)

	go func() {
		n := int64(0)
		basePath = appendPathSep(basePath)

		performMatch := func(path string, info os.FileInfo) error {
			if matchRegex != nil && !matchRegex.MatchString(info.Name()) {
				return nil
			}

			n++
			if n > searchLimitCount {
				return searchLimitReachedError
			}

			q <- filepath.Clean(path)
			return nil
		}

		var err error
		err = filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
			if info == nil {
				return err
			}
			name := info.Name()
			if matchDirectoriesOnly {
				if info.IsDir() && name != defaultBasePath {
					if ignoreRegexp.MatchString(name) {
						return filepath.SkipDir
					}
					return performMatch(path, info)
				}
				return nil
			} else {
				if !info.IsDir() {
					if ignoreRegexp.MatchString(name) {
						return nil
					}
					return performMatch(path, info)
				} else {
					if ignoreRegexp.MatchString(name) {
						return filepath.SkipDir
					}
				}
			}

			return nil
		})

		if err != nil && err != searchLimitReachedError {
			log.Fatal(err)
		}

		close(q)
	}()

	return q
}

func appendPathSep(basePath string) string {
	sep := string(os.PathSeparator)

	if !strings.HasSuffix(basePath, sep) {
		return fmt.Sprintf("%s%s", basePath, sep)
	}

	return basePath
}
