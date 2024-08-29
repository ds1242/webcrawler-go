package main

import (
	"fmt"
	"log"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {

	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	maxPagesReached := cfg.checkPagesLength()
	if maxPagesReached {
		return
	}
	fmt.Printf("crawling page: %s\n", rawCurrentURL)
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}

	if cfg.baseURL.Hostname() != currentURL.Hostname() {
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

	// fmt.Printf("crawling %s\n", rawCurrentURL)
	body, err := getHTML(rawCurrentURL)
	if err != nil {
		log.Printf("unable to get HTML body: %v\n", err)
	}

	sliceOfNextURLs, err := getURLsFromHTML(body, cfg.baseURL)
	if err != nil {
		fmt.Printf("error getting urls from body: %v\n", err)
		return
	}
	for _, url := range sliceOfNextURLs {
		cfg.wg.Add(1)
		go cfg.crawlPage(url)
	}
}

func (cfg *config) checkPagesLength() (maxPagesReached bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	return len(cfg.pages) > cfg.maxPages
}
