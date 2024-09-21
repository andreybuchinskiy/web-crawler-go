package main

import (
	"fmt"
	"net/url"
)

func normalizeURL(s string) (string, error) {
	u, err := url.Parse(s)
	if err != nil {
		return "", err
	}
	normalizedUrl := fmt.Sprintf("%s%s", u.Host, u.Path)
	return normalizedUrl, nil
}
