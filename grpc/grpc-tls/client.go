package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "grpc-hello-world/proto"
	"google.golang.org/grpc/credentials"
)

func main() {

	// 客户端认证TLS 需要服务端的证书server.pem
	t,_:=credentials.NewClientTLSFromFile("server.pem","go-grpc-tls")
	conn, err := grpc.Dial(":9001", grpc.WithTransportCredentials(t))
	defer conn.Close()
	if err != nil {
		log.Println(err)
	}

	c := pb.NewHelloWorldClient(conn)
	context := context.Background()
	body := &pb.HelloWorldRequest{
		Referer : "happy",
	}

	r, err := c.SayHelloWorld(context, body)
	if err != nil {
		log.Println(err)
	}

	log.Println(r.Message)
}
