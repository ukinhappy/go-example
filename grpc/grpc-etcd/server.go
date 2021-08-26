package main
//
//import (
//	"encoding/json"
//	"flag"
//	"fmt"
//	"github.com/coreos/etcd/clientv3"
//	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
//	"github.com/ukinhappy/go-example/grpc/pkg/proto"
//	"google.golang.org/grpc/naming"
//	"log"
//	"net"
//	"time"
//
//	"golang.org/x/net/context"
//	"google.golang.org/grpc"
//)
//
//var (
//	serv = flag.String("service", "hello_service", "service name")
//	port = flag.Int("port", 50001, "listening port")
//)
//
//// server is used to implement helloworld.GreeterServer.
//type server struct{}
//
//// SayHello implements helloworld.GreeterServer
//func (s *server) SayHelloWorld(ctx context.Context, in *proto.HelloWorldRequest) (*proto.HelloWorldResponse, error) {
//	fmt.Printf("%v: Receive is %s\n", time.Now(), in.Referer)
//	return &proto.HelloWorldResponse{Message: "Hello " + in.Referer}, nil
//}
//
//func main() {
//	flag.Parse()
//
//	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
//	if err != nil {
//		panic(err)
//	}
//
//	err = register("hello_service", "127.0.0.1", *port, "http://127.0.0.1:2379", time.Second*10, 15)
//	if err != nil {
//		panic(err)
//	}
//
//	s := grpc.NewServer()
//	proto.RegisterHelloWorldServer(s, &server{})
//	s.Serve(lis)
//}
//
//// register
//func register(name string, host string, port int, etcdaddr string, interval time.Duration, ttl int) error {
//	service := naming.Update{
//		Addr:     fmt.Sprintf("%s:%d", host, port),
//		Metadata: "...",
//	}
//	bts, err := json.Marshal(service)
//	if err != nil {
//		return err
//	}
//	serviceValue := string(bts)
//	serviceKey := fmt.Sprintf("/%s/%s", name, serviceValue)
//
//	client, err := clientv3.New(clientv3.Config{
//		Endpoints: []string{etcdaddr},
//	})
//	if err != nil {
//		return fmt.Errorf("grpclb: create etcd3 client failed: %v", err)
//	}
//
//	go func() {
//		// invoke self-register with ticker
//		ticker := time.NewTicker(interval)
//		for {
//			// minimum lease TTL is ttl-second
//			resp, _ := client.Grant(context.TODO(), int64(ttl))
//			// should get first, if not exist, set it
//			_, err := client.Get(context.Background(), serviceKey)
//			if err != nil {
//				if err == rpctypes.ErrKeyNotFound {
//					if _, err := client.Put(context.TODO(), serviceKey, serviceValue, clientv3.WithLease(resp.ID)); err != nil {
//						log.Printf("grpclb: set service '%s' with ttl to etcd3 failed: %s", name, err.Error())
//					}
//				} else {
//					log.Printf("grpclb: service '%s' connect to etcd3 failed: %s", name, err.Error())
//				}
//			} else {
//				// refresh set to true for not notifying the watcher
//				if _, err := client.Put(context.Background(), serviceKey, serviceValue, clientv3.WithLease(resp.ID)); err != nil {
//					log.Printf("grpclb: refresh service '%s' with ttl to etcd3 failed: %s", name, err.Error())
//				}
//			}
//			select {
//			case <-ticker.C:
//			}
//		}
//	}()
//
//	return nil
//}
