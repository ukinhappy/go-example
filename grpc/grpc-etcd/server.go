package main

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/ukinhappy/go-example/grpc/pkg/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type HelloWorldService struct{}

func (s *HelloWorldService) SayHelloWorld(ctx context.Context, r *proto.HelloWorldRequest) (*proto.HelloWorldResponse, error) {
	return &proto.HelloWorldResponse{Message: r.GetReferer() + " hello"}, nil
}
func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:18888")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	proto.RegisterHelloWorldServer(s, &HelloWorldService{})
	go func() {
		cli, err := clientv3.New(clientv3.Config{Endpoints: []string{"127.0.0.1:2379"}, DialTimeout: time.Second * 5})
		if err != nil {
			panic(err)
		}
		cli.Put(context.TODO(), "/grpc-etcd/server/127.0.0.1:18888", "127.0.0.1:18888")
	}()
	s.Serve(listen)
}
