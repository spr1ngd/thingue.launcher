package util

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func HttpGet(url string) (string, error) {
	var result []byte
	response, err := http.Get(url)
	defer response.Body.Close()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		result, err = io.ReadAll(response.Body)
	}
	return string(result), err
}

func HttpPost(url, data string) (string, error) {
	var result []byte
	response, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(data)))
	defer response.Body.Close()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		result, err = io.ReadAll(response.Body)
	}
	return string(result), err
}
