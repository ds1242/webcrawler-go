package main

import (
	"fmt"
	"sort"
)

type Page struct {
	pageURL 	string
	linkCount	int
}

func printReport(pages map[string]int, baseURL string) {
	
	var pageSlice []Page
	
	for normalizedURL, count := range pages {
		pageSlice = append(pageSlice, Page{
			pageURL: 	normalizedURL,
			linkCount: 	count,
		})
	}
	
	fmt.Printf(`=============================
REPORT for %s
=============================`, baseURL)
	fmt.Println()
	
	pageSlice = sortPageStruct(pageSlice)
	fmt.Println(pageSlice)
	// for _, page := range pageSlice{
	// 	fmt.Printf("Found %d internal links to %s\n", page.linkCount, page.pageURL)
	// }

}


func sortPageStruct(pageSlice []Page) []Page {
	sort.Slice(pageSlice, func(i, j int) bool {
		return pageSlice[i].linkCount > pageSlice[j].linkCount
	})
	return pageSlice
}
