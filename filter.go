package main

func filterURLs(urls []string, hours int) []string {
	uniqueURLs := make(map[string]bool)
	for _, url := range urls {
		uniqueURLs[url] = true
	}

	filtered := []string{}
	for url := range uniqueURLs {
		filtered = append(filtered, url)
	}
	return filtered
}
