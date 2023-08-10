package server

import (
	"fmt"
	"testing"
)

func TestDeviceInfo(t *testing.T) {
	info := GetDeviceInfo()
	fmt.Println(info)
}
