package main

import (
	"fmt"
	"os"
)


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
	const maxConcurrency = 5
	
	cfg, err := configure(rawBaseURL, maxConcurrency)
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
