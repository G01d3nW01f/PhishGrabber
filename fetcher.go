package main

import (
	"io"
	"net/http"
	"strings"
)

func fetchPhishingURLs() ([]string, error) {
	urls := []string{}
	sources := []string{
		"https://openphish.com/feed.txt",
		//someday add but need API	
	}

	for _, source := range sources {
		resp, err := http.Get(source)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		for _, line := range strings.Split(string(body), "\n") {
			if line != "" {
				urls = append(urls, line)
			}
		}
	}
	return urls, nil
}
