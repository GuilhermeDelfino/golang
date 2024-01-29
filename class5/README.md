# Class 5 - Starting with gRPC

#### Configuration

Go gRPC Lib
`$ go get -u google.golang.org/grpc`

You need to have **protoc** bin in your environment
https://grpc.io/docs/protoc-installation/

You need to install two Go plugins
`$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28`
`$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2`

It will create some necessary code for us

#### Proto File

You need a .proto file, in my case I created **hello.proto**

```
syntax = "proto3";

package helloworld;

option go_package = "./pb";
// Remote Procedure Call
service Greeter {
    rpc SayHello(HelloRequest) returns (HelloReply){}
}

message HelloRequest {
    string name = 1;
}
message HelloReply {
    string message = 1;
}
```

Then I created a folder named as **pb** in root dir for run a command later and the code generated goes to this folder

Later you created your .proto file, run this command

`$ go --go_out=. --go-grpc_out=. hello.proto`

It will read your .proto file and generate the code
In my case it created **hello_grpc.pb.go** and **hello.pb.go**
Now, with it you can create your **server** and **client** gRPC

---

Source:

- https://grpc.io/docs/languages/go/quickstart/
- https://github.com/grpc/grpc-go
- https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_client/main.go
