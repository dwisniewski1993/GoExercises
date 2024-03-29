package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/net/context/ctxhttp"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func parseUrl(urlStr string) *url.URL {
	if !strings.Contains(urlStr, "://") && !strings.HasPrefix(urlStr, "//") {
		urlStr = fmt.Sprintf("//%v", urlStr)
	}

	url, err := url.Parse(urlStr)
	if err != nil {
		log.Fatalf("parseUrl: could not parse url %q: %v", urlStr, err)
	}

	if url.Scheme == "" {
		url.Scheme = defaultUrlScheme
	}

	return url
}

func performRequest(ctx context.Context, url *url.URL, c *config) {
	req := newRequest(c.httpMethod, url, c.body, c.httpHeaders)

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := ctxhttp.Do(ctx, client, req)
	if err != nil {
		log.Fatalf("executeRequest: failed to read response: %v", err)
	}

	err = readResponseBody(req, resp, c)
	if err != nil {
		log.Fatalf("executeRequest: failed to read body: %v", err)
	}

	defer resp.Body.Close()

	if c.followRedirects && isHTTPRedirect(resp) {
		loc, err := resp.Location()
		if err != nil {
			if err == http.ErrNoLocation {
				log.Fatalf("redirect: unable to follow redirect")
			}
		}

		redirectsFollowedCount++
		if redirectsFollowedCount > maxRedirects {
			log.Fatalf("redirect: maximum number of redirects followed")
		}

		performRequest(ctx, loc, c)
	}
}

func newRequest(method string, url *url.URL, body string, headers headers) *http.Request {
	req, err := http.NewRequest(method, url.String(), strings.NewReader(body))
	if err != nil {
		log.Fatalf("newRequest: unable to create request: %v", err)
	}

	for _, header := range headers {
		k, v := parseHeader(header)
		req.Header.Add(k, v)
	}
	return req
}

func parseHeader(h string) (string, string) {
	i := strings.IndexRune(h, ':')
	if i == -1 {
		log.Fatalf("parseHeader: '%s' has invalid format", h)
	}
	return strings.TrimRight(h[:i], " "), strings.TrimLeft(h[i:], " :")
}

func readResponseBody(req *http.Request, resp *http.Response, c *config) error {
	if isHTTPRedirect(resp) || req.Method == http.MethodHead {
		return nil
	}

	var out io.Writer

	if c.saveOutput || c.outputFile != "" {
		filename := outputFile

		if filename == "" {
			tmpFile, err := ioutil.TempFile(".", req.URL.Path)
			if err != nil {
				return errors.New("readResponseBody: unable to create output file")
			}
			filename = tmpFile.Name()
		}

		f, err := os.Create(filename)
		if err != nil {
			return errors.New("readResponseBody: unable to create output file")
		}
		defer f.Close()
		out = f
	} else {
		out = os.Stdout
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("readResponseBody: failed to read response body")
	}

	out.Write(respBytes)
	return nil
}

func isHTTPRedirect(resp *http.Response) bool {
	return resp.StatusCode > 299 && resp.StatusCode < 400
}

type config struct {
	httpMethod      string
	body            string
	followRedirects bool
	httpHeaders     []string
	saveOutput      bool
	outputFile      string
}

type headers []string

func (h *headers) String() string {
	var o []string

	for _, v := range *h {
		o = append(o, fmt.Sprintf("-H %s", v))
	}
	return strings.Join(o, " ")
}

func (h *headers) Set(v string) error {
	*h = append(*h, v)
	return nil
}
