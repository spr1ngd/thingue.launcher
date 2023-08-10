package util

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func HttpGet(url string) ([]byte, error) {
	var result []byte
	response, err := http.Get(url)
	defer response.Body.Close()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		result, err = io.ReadAll(response.Body)
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

func HttpUrlToAgentWsUrl(httpUrl string) string {
	wsUrl := strings.Replace(httpUrl, "http://", "ws://", 1)
	wsUrl = strings.Replace(wsUrl, "https://", "wss://", 1)
	if strings.HasSuffix(wsUrl, "/") {
		wsUrl = wsUrl + "ws/agent"
	} else {
		wsUrl = wsUrl + "/ws/agent"
	}
	return wsUrl
}

func HttpUrlToStreamerWsUrl(httpUrl string) string {
	wsUrl := strings.Replace(httpUrl, "http://", "ws://", 1)
	wsUrl = strings.Replace(wsUrl, "https://", "wss://", 1)
	if strings.HasSuffix(wsUrl, "/") {
		wsUrl = wsUrl + "ws/streamer"
	} else {
		wsUrl = wsUrl + "/ws/streamer"
	}
	return wsUrl
}
