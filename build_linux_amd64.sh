# $AppVersion = "0.0.12"
# $GitCommit = git rev-parse HEAD
# $BuildDate = Get-Date -Format 'yyyy-MM-dd HH:mm:ss'
# wails build -ldflags "-X main.GitCommit=$GitCommit -X 'main.BuildDate=$BuildDate' -X main.AppVersion=$AppVersion" -o thingue-launcher-v$AppVersion
$env:GOOS="linux"
$env:GOARCH="amd64"

export AppVersion="0.0.12"
wails build -ldflags "-X main.GitCommit=`git rev-parse HEAD` -X 'main.BuildDate=`date "+%Y-%m-%d %H:%M:%S"`' -X main.AppVersion=0.0.4 -X main.AppVersion=$AppVersion" -o thingue-launcher-v$AppVersion
