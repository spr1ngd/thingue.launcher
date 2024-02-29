$AppVersion = "1.0.0"
$GitCommit = git rev-parse HEAD
$BuildDate = Get-Date -Format 'yyyy-MM-dd HH:mm:ss'
$env:GOOS="linux"
$env:GOARCH="amd64"

wails build -ldflags "-X main.GitCommit=$GitCommit -X 'main.BuildDate=$BuildDate' -X main.AppVersion=$AppVersion" -o thingue-launcher-v$AppVersion
