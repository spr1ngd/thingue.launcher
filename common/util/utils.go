package util

import (
	"net/url"
	"strings"
)

func HttpUrlToWsUrl(httpBaseUrl string, paths ...string) string {
	wsBaseUrl := strings.Replace(httpBaseUrl, "http://", "ws://", 1)
	wsBaseUrl = strings.Replace(wsBaseUrl, "https://", "wss://", 1)
	parsedURL, _ := url.Parse(wsBaseUrl)
	path := parsedURL.JoinPath(paths...)
	//parsedURL = parsedURL.ResolveReference(&url.URL{Path: path.Join(paths...)})
	return path.String()
}
