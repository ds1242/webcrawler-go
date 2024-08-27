package main

import (
	"fmt"
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

	baseURL := os.Args[1]

	fmt.Printf("starting crawl of: %s...\n", baseURL)

	pages := make(map[string]int)

	crawlPage(baseURL, baseURL, pages)

	for normalizedURL, count := range pages {
		fmt.Printf("%s: %d\n", normalizedURL, count)
	}
}
