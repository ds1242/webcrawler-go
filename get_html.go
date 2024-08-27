package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(res.Body)

	defer res.Body.Close()
	if res.StatusCode > 299 {
		return "", errors.New("Response failed with bad status code")
	}
	if err != nil {
		return "", err
	}

	fmt.Println(body)
	return "", nil
}
