@echo off

:: 设置应用版本
set AppVersion=0.0.12

:: 获取 Git 提交哈希
for /f "delims=" %%i in ('git rev-parse HEAD') do set GitCommit=%%i

:: 获取当前日期和时间
for /f "delims=" %%i in ('powershell -Command "Get-Date -Format 'yyyy-MM-dd HH:mm:ss'"') do set BuildDate=%%i

:: 设置目标平台环境变量
set GOOS=windows
set GOARCH=amd64

:: 执行构建命令
wails build -ldflags "-X main.GitCommit=%GitCommit% -X 'main.BuildDate=%BuildDate%' -X main.AppVersion=%AppVersion%" -o thingue-launcher-v%AppVersion%.exe
