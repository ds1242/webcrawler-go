package main

import (
	"fmt"
	"log"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	defer cfg.wg.Done()
	defer func() { <-cfg.concurrencyControl }()


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

	fmt.Println("getting urls from html body")
	sliceOfURLs, err := getURLsFromHTML(body, rawCurrentURL)
	if err != nil {
		fmt.Printf("error getting urls from body: %v\n", err)
		return
	}
	for _, url := range sliceOfURLs {
		cfg.wg.Add(1)
		cfg.concurrencyControl <- struct{}{}
		go cfg.crawlPage(url)
	}
}
