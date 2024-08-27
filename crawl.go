package main

import (
	"fmt"
	"net/url"
	"log"
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
	if err != nil {
		return pages, err
	}

	if pages[normalizedURL] > 0 {
		pages[normalizedURL]++
		return pages, nil
	} else {
		pages[normalizedURL]++
	}

	fmt.Printf("crawling %s\n", rawCurrentURL)
	body, err := getHTML(rawCurrentURL)
	if err != nil {
		log.Printf("unable to get HTML body: %v\n", err)
	}

	// fmt.Println("getting urls from html body")
	sliceOfURLs, err := getURLsFromHTML(body, rawCurrentURL)
	if err != nil {
		return pages, fmt.Errorf("error getting urls from body: %v", err)
	}
	for _, url := range sliceOfURLs {
		pages, err = crawlPage(rawBaseURL, url, pages)
		if err != nil {
			return pages, err
		}
	}
	
	return pages, nil

}
