package util

import (
	"encoding/json"
	"fmt"
)

func JsonStrToMapData(jsonStr string) map[string]interface{} {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		fmt.Println("JSON解码错误:", err)
		panic(err)
	}
	return data
}

func MapDataToJsonStr(data map[string]interface{}) string {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println("JSON编码错误:", err)
		panic(err)
	}
	return string(jsonBytes)
}
