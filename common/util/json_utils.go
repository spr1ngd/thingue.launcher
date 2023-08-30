package util

import (
	"encoding/json"
	"fmt"
)

func JsonStrToMapData(jsonStr []byte) map[string]interface{} {
	var data map[string]interface{}
	err := json.Unmarshal(jsonStr, &data)
	if err != nil {
		fmt.Println("JSON解码错误:", err)
		panic(err)
	}
	return data
}

func MapToJson(data map[string]interface{}) []byte {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println("JSON编码错误:", err)
		panic(err)
	}
	return jsonBytes
}
