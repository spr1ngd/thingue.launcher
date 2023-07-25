package util

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	servers := make(map[string]interface{})
	servers["server"] = []string{"1", "2"}
	data := make(map[string]interface{})
	data["type"] = "config"
	data["peerConnectionOptions"] = []map[string]interface{}{servers}
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println("JSON编码错误:", err)
		return
	}
	for k, v := range data {
		fmt.Println(k, v)
	}

	fmt.Println(string(jsonBytes))
}
