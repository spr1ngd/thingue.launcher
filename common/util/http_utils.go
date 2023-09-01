package util

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func HttpGet(url string) ([]byte, error) {
	var result []byte
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		result, err = io.ReadAll(response.Body)
		defer response.Body.Close()
	}
	return result, err
}

func HttpPost(url string, data []byte) ([]byte, error) {
	var result []byte
	response, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		result, err = io.ReadAll(response.Body)
		defer response.Body.Close()
	}
	return result, err
}

func HttpUrlToWsUrl(httpBaseUrl string, paths ...string) string {
	wsBaseUrl := strings.Replace(httpBaseUrl, "http://", "ws://", 1)
	wsBaseUrl = strings.Replace(wsBaseUrl, "https://", "wss://", 1)
	parsedURL, _ := url.Parse(wsBaseUrl)
	path := parsedURL.JoinPath(paths...)
	//parsedURL = parsedURL.ResolveReference(&url.URL{Path: path.Join(paths...)})
	return path.String()
}
