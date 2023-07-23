package main

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"thingue-launcher/agent"
	"thingue-launcher/server"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) SelectExePath(name string) string {
	selection, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select File",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "ThingUE (*.exe)",
				Pattern:     "*.exe",
			},
		},
	})
	if err != nil {
		return ""
	}
	return selection
}

func (a *App) ServerStart(addr string, basePath string) {
	server.Start(addr, basePath)
}

func (a *App) ServerShutdown() {
	server.Shutdown()
}

func (a *App) UnrealStart(exePath string, params []string) int {
	return agent.UnrealStart(exePath, params...)
}

func (a *App) UnrealStop(pid int) {
	agent.UnrealStop(pid)
}
