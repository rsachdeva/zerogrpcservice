package main

import (
	"fmt"
	"log"

	"github.com/rsachdeva/zerogrpcservice/third_party/zeropb"
	"google.golang.org/grpc"
)

func main() {
	// grpc.WithBlock() used for testing
	// grpc.WithInsecure() for now
	cc, err := grpc.Dial("localhost:50051", grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("err is %v", err)
	}
	defer cc.Close()

	c := zeropb.NewZeroServiceClient(cc)
	fmt.Printf(" Created client c is %#v", c)
}
