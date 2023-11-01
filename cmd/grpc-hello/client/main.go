// grpc hello
package main

import (
	"context"
	pb "github.com/Akai66/pandora/internal/pb/hello"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

const (
	_address     = "localhost:50051"
	_defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(_address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("client: failed to connect: %s\n", err)
	}
	log.Printf("client: start connect to: %s\n", _address)
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	name := _defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.HelloRequest{
		Name: name,
	}
	res, err := c.SayHello(ctx, req)
	if err != nil {
		log.Fatalf("client: failed to sayhello: %s\n", err)
	}
	log.Printf("client: Reply: %s\n", res.Reply)
}
