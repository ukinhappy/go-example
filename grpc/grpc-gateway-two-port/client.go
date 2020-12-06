package main

import (
	"log"

	pb "github.com/ukinhappy/go-example/grpc/pkg/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)
// http 请求

// -k参数忽略证书不受信问题
// curl -X POST -k https://localhost:8000/hello_world -d '{"referer": "restful_api"}'

func main() {
	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	defer conn.Close()

	if err != nil {
		log.Println(err)
	}

	c := pb.NewHelloWorldClient(conn)
	context := context.Background()
	body := &pb.HelloWorldRequest{
		Referer : "ukin",
	}

	r, err := c.SayHelloWorld(context, body)
	if err != nil {
		log.Println(err)
	}

	log.Println(r.Message)
}