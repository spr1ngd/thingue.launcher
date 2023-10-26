package server

import (
	"sync"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	ConnManager.Init()
	go func() {
		time.Sleep(6 * time.Second)
		ConnManager.Disconnect()
	}()
	wg.Wait()
}
