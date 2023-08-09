package unreal

import (
	"fmt"
	"testing"
	"thingue-launcher/agent/model"
	"time"
)

func TestStart(t *testing.T) {
	instance := model.Instance{
		Name:            "test1",
		ExecPath:        "E:\\UE\\ue4-game\\game\\Binaries\\Win64\\game.exe",
		LaunchArguments: []string{"-AudioMixer", "-RenderOffScreen", "-ForceRes", "-ResX=1920", "-ResX=1080"},
	}
	runner := NewRunner(&instance)

	i := 0
	for {
		i++
		fmt.Printf("main goroutine: i = %d,runner status = %t\n", i, runner.IsRunning)
		time.Sleep(time.Second)
		if i == 5 {
			runner.stop()
		}
	}
}
