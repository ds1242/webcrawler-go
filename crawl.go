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

	fmt.Println("normalize url and add to count")
	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		return pages, err
	}
	
	_, ok := pages[normalizedURL]
	if ok {
		pages[normalizedURL] += 1
	} else {
		pages[normalizedURL] = 1
	}

	fmt.Println("getting html")
	body, err := getHTML(rawBaseURL)
	if err != nil {
		return pages, fmt.Errorf("unable to get HTML body: %v", err)
	}

	// var sliceOfURLs []string
	// fmt.Println("getting urls from html body")
	sliceOfURLs, err := getURLsFromHTML(body, rawBaseURL)
	if err != nil {
		return pages, fmt.Errorf("error getting urls from body: %v", err)
	}
	
	for _, url := range sliceOfURLs {
		fmt.Printf("crawling slice currentURL: %v, new currentURL: %v", rawCurrentURL, url)
		pages, err = crawlPage(rawCurrentURL, url, pages)
		if err != nil {
			return pages, err
		}
	}
	
	return pages, nil

}
