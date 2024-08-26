package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	var linkSlice []string
	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return []string{"", ""}, err
	}

	linkSlice = parseNode(doc, linkSlice)

	fmt.Println(linkSlice)
	return linkSlice, err
}

func parseNode(n *html.Node, stringSlice []string) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				stringSlice = append(stringSlice, a.Val)
				break
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		stringSlice = parseNode(c, stringSlice)
	}
	return stringSlice
}
