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

	pages := make(map[string]int)
	
	pages, err := crawlPage(baseURL, baseURL, pages)
	if err != nil {
		fmt.Printf("error crawling page: %v\n", err)
		os.Exit(1)
	}

	for key, val := range pages {
		fmt.Printf("%s: %d\n", key, val)
	}
}
