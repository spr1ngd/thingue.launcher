package provider

import (
	"thingue-launcher/common/domain"
)

var VersionInfo = new(domain.VersionInfo)

func SetBuildInfo(appVersion, gitCommit, buildDate string) {
	VersionInfo.Version = appVersion
	VersionInfo.GitCommit = gitCommit
	VersionInfo.BuildDate = buildDate
}
