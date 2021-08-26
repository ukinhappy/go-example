package main
//
//import (
//	"fmt"
//	"github.com/coreos/etcd/clientv3"
//	"github.com/ukinhappy/go-example/grpc/pkg/proto"
//	"time"
//
//	"github.com/coreos/etcd/clientv3/naming"
//	"golang.org/x/net/context"
//	"google.golang.org/grpc"
//	"strconv"
//)
//
//func main() {
//	cli, cerr := clientv3.NewFromURL("http://127.0.0.1:2379")
//	if cerr != nil {
//		panic(cerr)
//	}
//
//	b := grpc.RoundRobin(&naming.GRPCResolver{Client: cli})
//	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
//	conn, err := grpc.DialContext(ctx, "/hello_service", grpc.WithInsecure(), grpc.WithBalancer(b))
//	if err != nil {
//		panic(err)
//	}
//
//	ticker := time.NewTicker(1 * time.Second)
//	for t := range ticker.C {
//		fmt.Println("client.....")
//		client := proto.NewHelloWorldClient(conn)
//		resp, err := client.SayHelloWorld(context.Background(), &proto.HelloWorldRequest{Referer: "world " + strconv.Itoa(t.Second())})
//		if err == nil {
//			fmt.Printf("%v: Reply is %s\n", t, resp.Message)
//		}
//		fmt.Println(err)
//	}
//}
