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

	if pages[normalizedURL] > 0 {
		pages[normalizedURL]++
		return pages, nil
	}
	
	pages[normalizedURL]++

	fmt.Println("getting html")
	body, err := getHTML(rawCurrentURL)
	if err != nil {
		return pages, fmt.Errorf("unable to get HTML body: %v", err)
	}

	// fmt.Println("getting urls from html body")
	sliceOfURLs, err := getURLsFromHTML(body, rawCurrentURL)
	if err != nil {
		return pages, fmt.Errorf("error getting urls from body: %v", err)
	}
	for _, url := range sliceOfURLs {
		// fmt.Printf("crawling slice baseURL: %v, new currentURL: %v\n", rawBaseURL, url)
		pages, err = crawlPage(rawBaseURL, url, pages)
		if err != nil {
			return pages, err
		}
		// fmt.Println(url)
	}
	
	return pages, nil

}
