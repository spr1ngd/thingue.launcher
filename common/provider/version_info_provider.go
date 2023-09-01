package provider

import (
	"thingue-launcher/common/constants"
	"thingue-launcher/common/domain"
)

var VersionInfo = new(domain.VersionInfo)

func SetVersionBuildInfo(gitCommit, buildDate string) {
	VersionInfo.Version = constants.APP_VERSION
	VersionInfo.GitCommit = gitCommit
	VersionInfo.BuildDate = buildDate
}
