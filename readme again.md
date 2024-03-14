
打包linux版本：

1. 注释

- runner.go : line 5
- runner.go : line 94

2. build client/frontend

- npm run build

3. 注意事项

1. 安装完go，要配置到etc/environment
2. 使用go install wails，安装完在当前用户目录/go/bin下，也要配置到etc/environment
3. 