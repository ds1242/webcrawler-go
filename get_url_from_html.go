package main

import (
	"fmt"
	"strings"
	"net/url"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return []string{}, err
	}

	var linkSlice []string
	var parsedURLSlice []string
	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return []string{}, err
	}

	linkSlice = parseNode(doc, linkSlice)
	for _, link := range linkSlice {
		u, err := url.Parse(link)
		if err != nil {
			return []string{}, err
		}
		
		resolvedURL := baseURL.ResolveReference(u)

		parsedURLSlice = append(parsedURLSlice, resolvedURL.String())

	}

	return parsedURLSlice, err
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
