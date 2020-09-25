package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"grpc-hello-world/grpc-gateway-two-port/proto"
	"net/http"
	"net"
	"time"
)

var (
	EndPoint string = ":9000"
)

type HelloWorldService struct{}

func (s *HelloWorldService) SayHelloWorld(ctx context.Context, r *proto.HelloWorldRequest) (*proto.HelloWorldResponse, error) {
	return &proto.HelloWorldResponse{Message: r.GetReferer() + " hello"}, nil
}


// grpc 和 http 使用两个端口启动。
func main() {

	// 启动grpc服务
	go func() {
		server := grpc.NewServer()
		proto.RegisterHelloWorldServer(server, &HelloWorldService{})

		lis, err := net.Listen("tcp", EndPoint)
		if err != nil {
			log.Fatalf("net.Listen err: %v", err)
		}
		server.Serve(lis)
	}()

	time.Sleep(time.Second)

	// gw server
	ctx := context.Background()
	gwmux := runtime.NewServeMux()

	// register grpc-gateway pb 作为代理去连接grpc server
	if err := proto.RegisterHelloWorldHandlerFromEndpoint(ctx, gwmux, EndPoint, []grpc.DialOption{grpc.WithInsecure()}); err != nil {
		log.Printf("Failed to register gw server: %v\n", err)
	}

	// 启动http 服务
	http.ListenAndServe(":9001", gwmux)
}
