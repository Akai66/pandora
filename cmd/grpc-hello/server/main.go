// grpc hello server
package main

import (
	"context"
	"fmt"
	pb "github.com/Akai66/pandora/internal/pb/hello"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	_port = ":50051"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (*server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("server: Received: %s\n", req.Name)
	return &pb.HelloReply{
		Reply: fmt.Sprintf("hello %s", req.Name),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", _port)
	if err != nil {
		log.Fatalf("server: failed to listen: %s\n", err)
	}
	log.Printf("server: start listen to: %s\n", _port)
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("server: failed to serve: %s\n", err)
	}
}
