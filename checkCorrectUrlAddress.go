package main

import (
	"net/url"
	"strings"
)

func IsCorrectUrlAddress(urlStr string) bool {
	if !strings.HasPrefix(urlStr, "http://") && !strings.HasPrefix(urlStr, "https://") {
		return false
	}

	_, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return false
	}

	return true
}