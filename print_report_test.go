package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)


func TestByPageCountAndTitle(t *testing.T) {
	pageSlice := []Page {
		{
			PageURL: "https://example.com/page2",
			LinkCount: 2,
		},
		{
			PageURL: "https://example.com/zsldjflaskjf",
			LinkCount: 1,
		},
		{
			PageURL: "https://example.com/page1",
			LinkCount: 3,
		},
		{
			PageURL: "https://example.com/page5",
			LinkCount: 1,
		},
		{
			PageURL: "https://example.com/another",
			LinkCount: 1,
		},
		
	}

	expected := []Page{
		{
			PageURL: "https://example.com/page1",
			LinkCount: 3,
		},
		{
			PageURL: "https://example.com/page2",
			LinkCount: 2,
		},
		{
			PageURL: "https://example.com/another",
			LinkCount: 1,
		},
		{
			PageURL: "https://example.com/page5",
			LinkCount: 1,
		},
		{
			PageURL: "https://example.com/zsldjflaskjf",
			LinkCount: 1,
		},
	}

	sort.Sort(ByPageCountAndTitle(pageSlice))


	if !reflect.DeepEqual(pageSlice, expected) {
		t.Errorf("Test FAIL: sorted pageSlice does not match expected\n")
		fmt.Printf("Page Slice: %v\n", pageSlice)
		fmt.Printf("Expected: %v\n", expected)
		return
	}
}