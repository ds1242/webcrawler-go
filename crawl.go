package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) (map[string]int, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return pages, err
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		return pages, err
	}

	if baseURL.Host != currentURL.Host {
		return pages, nil
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)

	return pages, nil

}
