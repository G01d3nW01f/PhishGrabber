package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func writeOutput(urls []string, filename, format string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	if format == "txt" {
		_, err = f.WriteString(strings.Join(urls, "\n"))
	} else if format == "json" {
	
		type URLData struct {
			URLs []string `json:"urls"`
		}
		
		// Create JSON data
		data := URLData{URLs: urls}
		
		// Create JSON encoder
		encoder := json.NewEncoder(f)
		// Set indentation for readable output
		encoder.SetIndent("", "  ")
		
		// Encode and write JSON
		err = encoder.Encode(data)
	}
	return err
}

// analyzeURLs File read
func analyzeURLs(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	client := &http.Client{
		Timeout: 5 * time.Second, // Timeout 5sec
	}

	for scanner.Scan() {
		url := scanner.Text()
		if url == "" {
			continue
		}

		// send HTTP Request
		status := "N/A"
		resp, err := client.Get(url)
		if err == nil {
			status = resp.Status
			resp.Body.Close()
		} else {
			log.Printf("Request failed for %s: %v", url, err)
		}

		// resolution host name
		ip := "N/A"
		host := strings.TrimPrefix(strings.TrimPrefix(url, "http://"), "https://")
		if idx := strings.Index(host, "/"); idx != -1 {
			host = host[:idx] 
		}
		ips, err := net.LookupIP(host)
		if err == nil && len(ips) > 0 {
			ip = ips[0].String() 
		}

		// result 
		fmt.Printf("%s [%s] %s\n", url, status, ip)
	}

	return scanner.Err()
}
