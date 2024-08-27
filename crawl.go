package main

import (
	"fmt"
	"net/url"
	"log"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawBaseURL, err)
		return 
	}
	
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return 
	}

	if baseURL.Host != currentURL.Host {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - normalizedURL: %v", err)
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
