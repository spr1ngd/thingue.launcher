package util

import (
	"fmt"
	"testing"
)

func TestHttpUrlToWsUrl(t *testing.T) {
	fmt.Println(HttpUrlToWsUrl("http://127.0.0.1:8090/", "/ws/streamer"))
	fmt.Println(HttpUrlToWsUrl("http://127.0.0.1:8090", "ws/node"))
	fmt.Println(HttpUrlToWsUrl("http://127.0.0.1:8090/uino/", "/ws", "/streamer"))
}

func TestHash(t *testing.T) {
	hash := CalculateFileHash("E:\\UE\\ue5-fps\\fps.exe")
	fmt.Println(hash)
}
