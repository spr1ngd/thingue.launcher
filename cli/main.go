package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"thingue-launcher/common/provider"
)

var (
	GitCommit  string
	BuildDate  string
	AppVersion string
)

var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "ThingUE command line",
}

func main() {
	provider.SetBuildInfo(AppVersion, GitCommit, BuildDate)

	rootCmd.SetArgs(os.Args[1:])
	if err := rootCmd.Execute(); err != nil {
		e := err.Error()
		fmt.Println(strings.ToUpper(e[:1]) + e[1:])
		os.Exit(1)
	}
}
