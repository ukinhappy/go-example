package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/ukinhappy/go-example/grpc/pkg/proto"
)
// http 请求

// curl -X POST -k https://localhost:8000/hello_world -d '{"referer": "restful_api"}'
func main() {
	creds, err := credentials.NewClientTLSFromFile("server.pem", "grpc-gateway")
	if err != nil {
		log.Println("Failed to create TLS credentials %v", err)
	}
	conn, err := grpc.Dial(":8000", grpc.WithTransportCredentials(creds))
	defer conn.Close()

	if err != nil {
		log.Println(err)
	}

	c := pb.NewHelloWorldClient(conn)
	context := context.Background()
	body := &pb.HelloWorldRequest{
		Referer : "Grpc",
	}

	r, err := c.SayHelloWorld(context, body)
	if err != nil {
		log.Println(err)
	}

	log.Println(r.Message)
}