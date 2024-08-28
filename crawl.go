package main

import (
	"fmt"
	"log"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {

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
		return
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
		fmt.Errorf("error getting urls from body: %v", err)
		return
	}
	for _, url := range sliceOfURLs {
		crawlPage(rawBaseURL, url, pages)
	}
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {

}