package main

import (
	"flag"
	"fmt"
	"io"
	"log"
)

var (
	outputFile = flag.String("o", "phish_list.txt", "Output file name")
	format     = flag.String("f", "txt", "Output format (txt/json)")
	timeRange  = flag.Int("t", 0, "Time range in hours (e.g., 24)")
	verbose    = flag.Bool("verbose", false, "Enable verbose logging")
)

func main() {
	flag.Parse()

	if *verbose {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	} else {
		log.SetOutput(io.Discard)
	}

	
	urls, err := fetchPhishingURLs()
	if err != nil {
		log.Fatalf("Error fetching URLs: %v", err)
	}

	filteredURLs := filterURLs(urls, *timeRange)
	if err := writeOutput(filteredURLs, *outputFile, *format); err != nil {
		log.Fatalf("Error writing output: %v", err)
	}

	fmt.Printf("Saved %d phishing URLs to %s\n", len(filteredURLs), *outputFile)

	
	if err := analyzeURLs(*outputFile); err != nil {
		log.Fatalf("Error analyzing URLs: %v", err)
	}
}
