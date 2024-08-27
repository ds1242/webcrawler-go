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

	body, err := getHTML(baseURL)
	if err != nil {
		fmt.Errorf("unable to get HTML body: %v", err)
		os.Exit(1)
	}

	var sliceOfURLs []string

	sliceOfURLs, err = getURLsFromHTML(body, baseURL)
	if err != nil {
		fmt.Errorf("an error occurred parsing the urls from the html: %v", err)
		os.Exit(1)
	}

	fmt.Println(sliceOfURLs)
}
