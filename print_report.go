package main

import (
	"fmt"
)

type Page struct {
	pageURL 	string
	linkCount	int
}

func printReport(pages map[string]int, baseURL string) {
	fmt.Printf(`=============================
REPORT for %s
=============================`, baseURL)
	fmt.Println()

	var pageSlice []Page

	for normalizedURL, count := range pages {
		pageSlice = append(pageSlice, Page{
			pageURL: 	normalizedURL,
			linkCount: 	count,
		})
	}

	for _, page := range pageSlice{
		fmt.Printf("Found %d internal links to %s\n", page.linkCount, page.pageURL)
	}

}
