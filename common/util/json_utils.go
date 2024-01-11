package util

import (
	"encoding/json"
	"thingue-launcher/common/logger"
)

func JsonStrToMapData(jsonStr []byte) map[string]any {
	var data map[string]any
	err := json.Unmarshal(jsonStr, &data)
	if err != nil {
		logger.Zap.Error("JSON解码错误:", err, string(jsonStr))
	}
	return data
}

func MapToJson(data map[string]any) []byte {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		logger.Zap.Error("JSON编码错误:", err)
	}
	return jsonBytes
}
