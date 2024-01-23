# usage
## 安装工具连
```shell
go install github.com/bufbuild/buf/cmd/buf@v1.28.1
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

## 检查protobuf语法
```shell
buf lint
```

## 生成代码
```shell
buf generate
```