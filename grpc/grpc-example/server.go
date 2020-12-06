package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/ukinhappy/go-example/grpc/pkg/proto"
)


type HelloWorldService struct{}

func (s *HelloWorldService) SayHelloWorld(ctx context.Context, r *proto.HelloWorldRequest) (*proto.HelloWorldResponse, error) {
	return &proto.HelloWorldResponse{Message: r.GetReferer() + " hello"}, nil
}

const PORT = "9000"

func main() {

	server := grpc.NewServer()
	proto.RegisterHelloWorldServer(server, &HelloWorldService{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	server.Serve(lis)
}
