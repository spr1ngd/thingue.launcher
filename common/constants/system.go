package constants

import "thingue-launcher/common/domain"

const (
	SaveDir    = "./thingue-launcher"
	ConfigName = "config.yaml"
)

var VersionInfo = new(domain.VersionInfo)

func SetVersionInfo(appVersion, gitCommit, buildDate string) {
	VersionInfo.Version = appVersion
	VersionInfo.GitCommit = gitCommit
	VersionInfo.BuildDate = buildDate
}
