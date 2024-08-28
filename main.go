package main

import (
	"fmt"
	"os"
	"strconv"
)


func main() {
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		return
	}
	if len(os.Args) > 4 {
		fmt.Println("too many arguments provided")
		return
	}
	var maxPagesArg string
	var maxConcurrencyArg string

	var maxPages int
	var maxConcurrency int

	rawBaseURL := os.Args[1]
	if len(os.Args[2]) > 0 {
		maxConcurrencyArg = os.Args[2]
	} else {
		maxConcurrencyArg = "3"
	}
	if len(os.Args[3]) > 0 {
		maxPagesArg = os.Args[3]
	} else {
		maxPagesArg = "25"
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
}
