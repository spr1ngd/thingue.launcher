package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
	"thingue-launcher/common/provider"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   `version`,
	Short: "打印版本信息",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("ThingUE Cli %s v%s\n", runtime.GOOS, provider.VersionInfo.Version)
		fmt.Printf("BuildDate %s\n", provider.VersionInfo.BuildDate)
		fmt.Printf("GitCommit %s\n", provider.VersionInfo.GitCommit)
		return nil
	},
}
