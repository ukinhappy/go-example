package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "go-example/grpc/pkg/proto"
)

func main() {
	conn, err := grpc.Dial(":9001", grpc.WithInsecure())
	defer conn.Close()

	if err != nil {
		log.Println(err)
	}

	c := pb.NewHelloWorldClient(conn)
	context := context.Background()
	body := &pb.HelloWorldRequest{
		Referer : "uink",
	}

	r, err := c.SayHelloWorld(context, body)
	if err != nil {
		log.Println(err)
	}

	log.Println(r.Message)
}
