package util

import (
	"encoding/json"
	"fmt"
)

func JsonStrToMapData(jsonStr []byte) map[string]interface{} {
	var data map[string]any
	err := json.Unmarshal(jsonStr, &data)
	if err != nil {
		fmt.Println("JSON解码错误:", err)
	}
	return data
}

func MapToJson(data map[string]any) []byte {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println("JSON编码错误:", err)
	}
	return jsonBytes
}
