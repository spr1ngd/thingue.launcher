package util

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
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
	defer response.Body.Close()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		result, err = io.ReadAll(response.Body)
	}
	return result, err
}
