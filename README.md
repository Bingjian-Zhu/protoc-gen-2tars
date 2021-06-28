# protoc-gen-2tars

## 安装

```
go get github.com/bingjian-zhu/protoc-gen-2tars@master
```

依赖：

- [protoc](https://github.com/google/protobuf)
- [protoc-gen-go](https://github.com/golang/protobuf)

go get -u github.com/golang/protobuf/proto@v1.3.5
go get -u github.com/golang/protobuf/protoc-gen-go@v1.3.5
注意：要选v1.3.5

```
go build .
```
把生成的protoc-gen-2tars复制到环境变量的路径下

## 使用

定义`helloworld.proto`

```
syntax = "proto3";
package pb;

// The greeting service definition.
service Greeter {
  // Sends a greeting
  rpc SayHello (HelloRequest) returns (HelloReply) {}
}

// The request message containing the user's name.
message HelloRequest {
  string name = 1;
}

// The response message containing the greetings
message HelloReply {
  string message = 1;
}
```

进入`helloworld.proto`所在目录proto转成tars

```
protoc --2tars_out=plugins=tarsrpc:. helloworld.proto
```

Your output result should be:

```
./
├── helloworld.pb.go
└── helloworld.proto
```

## 测试

server:
```
go run example/server/main.go -config example/config.conf
```

client:
```
go run example/client/main.go -config example/config.conf
```

output:
```
result is: hellosandyskies
```