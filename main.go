package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sync"
)

type config struct {
	pages 				map[string]int
	baseURL				*url.URL
	mu 					*sync.Mutex
	concurrencyControl 	chan struct{}
	wg 					*sync.WaitGroup
}


func main() {
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		return
	}
	if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		return
	}

	rawBaseURL := os.Args[1]
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawBaseURL, err)
		return
	}
	
	var waitGroup sync.WaitGroup

	var mux sync.Mutex

	channel := make(chan struct{}, 1)

	cfg := config {
		pages:	 	make(map[string]int),
		baseURL: 	baseURL,
		mu: 		&mux,
		concurrencyControl: channel,
		wg: 		&waitGroup,
	}
	
	fmt.Printf("starting crawl of: %s...\n", baseURL)

	pages := make(map[string]int)

	cfg.crawlPage()

	for normalizedURL, count := range pages {
		fmt.Printf("%s: %d\n", normalizedURL, count)
	}
}
