# README

## Install
### windows
#### gcc安装
由于依赖的sqlite数据库驱动需要gcc，所以在windows下需要通过msys2安装gcc并将bin目录配置到Path

1、下载安装[msys2](https://mirrors.tuna.tsinghua.edu.cn/msys2/distrib/msys2-x86_64-latest.exe)

2、打开`MSYS2 MINGW64`终端使用以下命令安装gcc
```bash
# 设置镜像源加速安装
sed -i "s#https\?://mirror.msys2.org/#https://mirrors.tuna.tsinghua.edu.cn/msys2/#g" /etc/pacman.d/mirrorlist*
# 安装gcc
pacman -S mingw-w64-x86_64-gcc
```
3、将`C:\msys64\mingw64\bin`配置到Path下
#### go安装配置
[下载地址](https://go.dev/dl/go1.21.6.windows-amd64.msi)，安装完成后执行
```bash
# 开启cgo
go env -w CGO_ENABLED=1
# 设置goproxy加速依赖下载
go env -w GOPROXY=https://goproxy.cn,direct
# 安装wails cli
go install github.com/wailsapp/wails/v2/cmd/wails@latest
# 检查开发环境，Linux下可能需要根据提示额外安装一些包
wails doctor
```

## Live Development
在项目目录下运行`wails dev`进入实时开发模式，代码变更会实时生效，如果提示找不到命令需要把`用户目录\go\bin`配置到Path里

## Building
### Windows客户端编译
```shell
# 设置版本号
$AppVersion = "0.0.11"
# 构建不带控制台的程序
wails build -ldflags "-X main.GitCommit=$(git rev-parse HEAD) -X 'main.BuildDate=$(Get-Date -Format 'yyyy-MM-dd HH:mm:ss')' -X main.AppVersion=$AppVersion" -o thingue-launcher-v$AppVersion.exe
# 构建带控制台的程序，闪退情况下比较方便排查问题
wails build -ldflags "-X main.GitCommit=$(git rev-parse HEAD) -X 'main.BuildDate=$(Get-Date -Format 'yyyy-MM-dd HH:mm:ss')' -X main.AppVersion=$AppVersion" -o thingue-launcher-v$AppVersion.exe -windowsconsole
```
### Linux客户端编译
```shell
export AppVersion="0.0.10"
wails build -ldflags "-X main.GitCommit=`git rev-parse HEAD` -X 'main.BuildDate=`date "+%Y-%m-%d %H:%M:%S"`' -X main.AppVersion=0.0.4 -X main.AppVersion=$AppVersion" -o thingue-launcher-v$AppVersion
```
### Windows独立服务程序编译
```shell
$AppVersion = "0.0.10"
go build -o build/bin/thingue-server-v$AppVersion.exe -ldflags "-X main.GitCommit=$(git rev-parse HEAD) -X 'main.BuildDate=$(Get-Date -Format 'yyyy-MM-dd HH:mm:ss')' -X main.AppVersion=$AppVersion" server/main.go
```
### Linux独立服务程序编译
```shell
export AppVersion="0.0.10"
go build -o build/bin/thingue-server_$AppVersion -ldflags "-X main.GitCommit=`git rev-parse HEAD` -X 'main.BuildDate=`date "+%Y-%m-%d %H:%M:%S"`' -X main.AppVersion=$AppVersion" server/main.go
```
