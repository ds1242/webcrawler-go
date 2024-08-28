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

	if cfg.baseURL.Host != currentURL.Host {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - normalizedURL: %v", err)
	}

	isFirst := cfg.addPageVisit(normalizedURL)
	if !isFirst {
		return
	}
	

	fmt.Printf("crawling %s\n", rawCurrentURL)
	body, err := getHTML(rawCurrentURL)
	if err != nil {
		log.Printf("unable to get HTML body: %v\n", err)
	}

	// fmt.Println("getting urls from html body")
	sliceOfURLs, err := getURLsFromHTML(body, rawCurrentURL)
	if err != nil {
		fmt.Printf("error getting urls from body: %v\n", err)
		return
	}
	for _, url := range sliceOfURLs {
		cfg.crawlPage(url)
	}
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	// lock the map to prevent it from being overwritten or edited by multiple goroutines
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if cfg.pages[normalizedURL] > 0 {
		cfg.pages[normalizedURL]++
		isFirst = false
	} else {
		cfg.pages[normalizedURL]++
		isFirst = true
	}
	
	return isFirst
}