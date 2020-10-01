package main

import (
	"fmt"
	"log"
	"net"

	"github.com/rsachdeva/zerogrpcservice/third_party/zeropb"
	"google.golang.org/grpc"
)

type exampleZeroServiceServer struct{}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("could not listen %v", err)
	}
	s := grpc.NewServer()
	zeropb.RegisterZeroServiceServer(s, exampleZeroServiceServer{})

	fmt.Println("Ready to serve")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("error is %#v", err)
	}
}
