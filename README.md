# README

## 设置goproxy加速国内依赖下载
```shell
go env -w GOPROXY=https://goproxy.cn,direct
```

## 1、GUI应用程序
### 安装
```bash
#安装wails cli
go install github.com/wailsapp/wails/v2/cmd/wails@latest
#检查wails开发环境
wails doctor
```
### 实时开发预览
```bash
wails dev
```
### 构建
Windows
```shell
$AppVersion = "0.0.10"
wails build -ldflags "-X main.GitCommit=$(git rev-parse HEAD) -X 'main.BuildDate=$(Get-Date -Format 'yyyy-MM-dd HH:mm:ss')' -X main.AppVersion=$AppVersion" -o thingue-launcher-v$AppVersion.exe
wails build -ldflags "-X main.GitCommit=$(git rev-parse HEAD) -X 'main.BuildDate=$(Get-Date -Format 'yyyy-MM-dd HH:mm:ss')' -X main.AppVersion=$AppVersion" -o thingue-launcher-v$AppVersion.exe -windowsconsole
```
Linux
```shell
export AppVersion="0.0.10"
wails build -ldflags "-X main.GitCommit=`git rev-parse HEAD` -X 'main.BuildDate=`date "+%Y-%m-%d %H:%M:%S"`' -X main.AppVersion=0.0.4 -X main.AppVersion=$AppVersion" -o thingue-launcher-v$AppVersion
```

## 2、命令行工具
### 构建
Windows
```shell
$AppVersion = "0.0.10"
go build -o build/bin/cli.exe -ldflags "-X main.GitCommit=$(git rev-parse HEAD) -X 'main.BuildDate=$(Get-Date -Format 'yyyy-MM-dd HH:mm:ss')' -X main.AppVersion=$AppVersion" ./cli
```
Linux
```shell
export AppVersion=0.0.10 GOOS=linux GOARCH=amd64
go build -o build/bin/cli -ldflags "-X main.GitCommit=`git rev-parse HEAD` -X 'main.BuildDate=`date "+%Y-%m-%d %H:%M:%S"`' -X main.AppVersion=$AppVersion" ./cli
```
### 通过设置GOOS、GOARCH跨平台构建
powershell
```shell
$GOOS = windows
$GOARCH = amd64
```
bash
```shell
export GOOS=linux GOARCH=amd64
```