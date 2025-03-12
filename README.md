# Phishing URL Collector and Analyzer

This is a Go application that collects phishing URLs from various sources and analyzes them, providing output in either TXT or JSON format. The tool fetches known phishing URLs, filters them based on a specified time range, and performs basic analysis including HTTP status checks and IP resolution.

## Features

- Fetches phishing URLs from public feeds (currently OpenPhish)
- Supports TXT and JSON output formats
- Analyzes URLs with HTTP status and IP resolution
- Configurable time range filtering
- Verbose logging option
- Command-line interface with flags

## Prerequisites

- Go 1.16 or higher
- Internet connection for fetching URL feeds

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/phishing-url-analyzer.git
