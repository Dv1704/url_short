package util

import (
	"os"
	"strings"
)

func EnforceHTTP(url string) string {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		return "http://" + url
	}
	return url
}

func RemoveDomainError(url string) bool {
	// Removes prefixes like http://, https://, www.
	// Then checks if the cleaned URL equals the expected DOMAIN

	domain := os.Getenv("DOMAIN")

	
	if url == domain {
		return false
	}

	// Clean the URL
	newURL := strings.TrimPrefix(url, "http://")
	newURL = strings.TrimPrefix(newURL, "https://")
	newURL = strings.TrimPrefix(newURL, "www.")

	// Extract domain part before any path
	newURL = strings.SplitN(newURL, "/", 2)[0]

	// More concise return
	return newURL != domain
}
