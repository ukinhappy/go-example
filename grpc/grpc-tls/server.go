package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"go-example/grpc/pkg/proto"
	"google.golang.org/grpc/credentials"
)

type HelloWorldService struct{}

func (s *HelloWorldService) SayHelloWorld(ctx context.Context, r *proto.HelloWorldRequest) (*proto.HelloWorldResponse, error) {
	return &proto.HelloWorldResponse{Message: r.GetReferer() + " hello"}, nil
}

const PORT = "9001"

func main() {

	//  开启tls
	tls,_:=credentials.NewServerTLSFromFile("server.pem","server.key")


	server := grpc.NewServer(grpc.Creds(tls))
	proto.RegisterHelloWorldServer(server, &HelloWorldService{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	server.Serve(lis)
}