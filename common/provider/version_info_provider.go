package provider

import "thingue-launcher/common/domain"

var VersionInfo = new(domain.VersionInfo)

func SetVersionBuildInfo(gitCommit, buildDate string) {
	VersionInfo.Version = "0.0.1"
	VersionInfo.GitCommit = gitCommit
	VersionInfo.BuildDate = buildDate
}
