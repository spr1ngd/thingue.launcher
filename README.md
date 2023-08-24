# README

## About

This is the official Wails Vue template.

You can configure the project by editing `wails.json`. More information about the project settings can be found
here: https://wails.io/docs/reference/project-config

## Install

## Live Development
`sed -i "s#https\?://mirror.msys2.org/#https://mirrors.tuna.tsinghua.edu.cn/msys2/#g" /etc/pacman.d/mirrorlist*`
`pacman -S mingw-w64-x86_64-toolchain`
`go env -w CGO_ENABLED=1`
`go env -w GOPROXY=https://goproxy.cn,direct`
To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use 
`wails build -ldflags "-X main.GitCommit=$(git rev-parse HEAD) -X 'main.BuildDate=$(Get-Date -Format 'yyyy-MM-dd HH:mm:ss')'"`.

