package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(res.Body)

	defer res.Body.Close()
	if res.StatusCode > 399 {
		return "", errors.New("response failed with bad status code")
	}
	if err != nil {
		return "", err
	}
	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		err := fmt.Errorf("wrong content type: %v", contentType)
		return "", err
	}

	bodyHTML := string(body)

	return bodyHTML, nil
}
