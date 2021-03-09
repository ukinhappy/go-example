package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/naming"
	"github.com/ukinhappy/go-example/grpc/pkg/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	cli, cerr := clientv3.NewFromURL("127.0.0.1:2379")
	if cerr != nil {
		log.Fatal(cerr)
	}
	r := &naming.GRPCResolver{Client: cli}
	b := grpc.RoundRobin(r)
	conn, gerr := grpc.Dial("/grpc-etcd/server/127.0.0.1:18888", grpc.WithBalancer(b))

	if gerr != nil {
		log.Fatal(gerr)
	}

	client := proto.NewHelloWorldClient(conn)
	resp, err := client.SayHelloWorld(context.TODO(), &proto.HelloWorldRequest{Referer: "happy"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Message)
}
