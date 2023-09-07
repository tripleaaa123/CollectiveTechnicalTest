package main

import (
	"io"
	"net/http"
	"regexp"
)

func GetProjectURLs() ([]string, error) {
	resp, err := http.Get("https://raw.githubusercontent.com/avelino/awesome-go/master/README.md")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	bodyString := string(bodyBytes)
	urlRegex := regexp.MustCompile(`https://github.com/[\w-]+/[\w-]+`)
	projectURLs := urlRegex.FindAllString(bodyString, -1)

	return projectURLs, nil
}
