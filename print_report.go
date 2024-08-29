package main

import (
	"fmt"
	"sort"
)

type Page struct {
	PageURL 	string
	LinkCount	int
}

func printReport(pages map[string]int, baseURL string) {
	
	var pageSlice []Page
	
	for normalizedURL, count := range pages {
		pageSlice = append(pageSlice, Page{
			PageURL: 	normalizedURL,
			LinkCount: 	count,
		})
	}
	sort.Sort(ByPageCountAndTitle(pageSlice))
	
	fmt.Printf(`=============================
REPORT for %s
=============================`, baseURL)
	fmt.Println()

	for _, page := range pageSlice{
		fmt.Printf("Found %d internal links to %s\n", page.LinkCount, page.PageURL)
	}
}


type ByPageCountAndTitle []Page

func (p ByPageCountAndTitle) Len() int           { return len(p) }
func (p ByPageCountAndTitle) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p ByPageCountAndTitle) Less(i, j int) bool { 
	if p[i].LinkCount == p[j].LinkCount {
		return p[i].PageURL < p[j].PageURL
	}
	return p[i].LinkCount > p[j].LinkCount 
}
