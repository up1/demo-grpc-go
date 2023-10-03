# Workshop with gRPC

## Step 1 :: Design proto file
* /proto/hello.proto

## Step 2 :: Generate code with [Protocol Buffer Compile](https://grpc.io/docs/protoc-installation/)
* Server-side code
* Client-side code

Working with [go](https://grpc.io/docs/languages/go/quickstart/)
```
$go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

$protoc --proto_path=. --go_out=./hello --go_opt=paths=source_relative --go-grpc_out=./hello --go-grpc_opt=paths=source_relative hello.proto
```

## Step 3 :: Start server and client

Start server
```
$go run cmd/server.go
```

Call from client
```
$go run cmd/client.go
```

## Step 4 :: Performance testing with [ghz](https://ghz.sh/)
```
ghz --insecure \
  --proto ./hello.proto \
  --call hello.HelloWorld.SayHi \
  -c 100 -n 100000  \
  -d '{"id": 1, "message":"Hello"}' \
  0.0.0.0:9000
```


