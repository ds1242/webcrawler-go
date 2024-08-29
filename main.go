package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var maxPagesArg string
	var maxConcurrencyArg string

	var maxPages int
	var maxConcurrency int

	if len(os.Args) < 4 {
		fmt.Println("not enough arguments provided")
		fmt.Println("usage: crawler <baseURL> <maxConcurrency> <maxPages>")
		return
	}
	if len(os.Args) > 4 {
		fmt.Println("too many arguments provided")
		return
	}

	rawBaseURL := os.Args[1]

	if len(os.Args) == 3 {
		maxConcurrencyArg = os.Args[2]
	}

	if len(os.Args) == 4 {
		maxPagesArg = os.Args[3]
	} 

	maxConcurrency, err := strconv.Atoi(maxConcurrencyArg)
	if err != nil {
		fmt.Printf("Error converting arg type: %v", err)
		return
	}

	maxPages, err = strconv.Atoi(maxPagesArg)
	if err != nil {
		fmt.Printf("Error converting arg type: %v", err)
		return
	}

	cfg, err := configure(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
		return
	}

	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	fmt.Println("All pages crawled")

	for normalizedURL, count := range cfg.pages {
		fmt.Printf("%s: %d\n", normalizedURL, count)
	}

	// fmt.Printf("number of pages crawled: %v\n", len(cfg.pages))
}
