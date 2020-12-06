package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "github.com/ukinhappy/go-example/grpc/pkg/proto"
	"google.golang.org/grpc/credentials"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
)

func main() {

	// 获取客户端证书
	cert, _ := tls.LoadX509KeyPair("client.pem", "client.key")

	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("ca.pem")
	certPool.AppendCertsFromPEM(ca)

	t := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "grpc-ca-tls",
		RootCAs:      certPool,
	})


	conn, err := grpc.Dial(":9001", grpc.WithTransportCredentials(t))
	defer conn.Close()
	if err != nil {
		log.Println(err)
	}

	c := pb.NewHelloWorldClient(conn)
	context := context.Background()
	body := &pb.HelloWorldRequest{
		Referer: "uink",
	}

	r, err := c.SayHelloWorld(context, body)
	if err != nil {
		log.Println(err)
	}

	log.Println(r.Message)
}
