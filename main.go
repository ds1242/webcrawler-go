package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURL := os.Args[1]
	fmt.Printf("starting crawl of: %s\n", baseURL)

	pages := make(map[string]int)
	
	pages, err := crawlPage(baseURL, baseURL, pages)
	if err != nil {
		fmt.Printf("error crawling page: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(pages)
	//if err != nil {
	//	fmt.Errorf("an error occurred parsing the urls from the html: %v", err)
	//	os.Exit(1)
	//}

	//fmt.Println(sliceOfURLs)
}
