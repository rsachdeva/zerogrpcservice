# zerogrpcservice
Zero in Go style here refers to Empty gRPC Service.
This project is to help developers starting with protobuf and gRPC using Go.

Shows common code from scratch for protobuf/gRPC -- proto with ZeroService; 
grpc Server start and Client connection with ZeroService.
The project already has go.mod for protobuf and gRPC libraries.

**Steps to get started After cloning/forking project:(assuming root of the project)**

brew install protobuf (for MacOS and similarly for other OS; for protoc)

export GO111MODULE=on go get -u github.com/golang/protobuf/protoc-gen-go (builds and installs protoc-gen-go plugin)

That's it....You can start with protobuf/gRPC with Golang.

Execute generatezeroservice.sh ( overwrites existing internal/zeropb/zeroservice.pb.go generated from proto/zeroservice.proto)  

Execute runserver.sh  

Execute runclient.sh

Go through the 'Zero' service server and client code.

Tweak Module Name.

Add your Service to zeroservice.proto and rename proto file as appropriate.

Update server.go and client.go accordingly.

Profit!
