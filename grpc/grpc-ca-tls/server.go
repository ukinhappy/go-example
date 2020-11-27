package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"go-example/grpc/pkg/proto"
	"google.golang.org/grpc/credentials"
	"crypto/x509"
	"io/ioutil"
	"crypto/tls"
)

type HelloWorldService struct{}

func (s *HelloWorldService) SayHelloWorld(ctx context.Context, r *proto.HelloWorldRequest) (*proto.HelloWorldResponse, error) {
	return &proto.HelloWorldResponse{Message: r.GetReferer() + " hello"}, nil
}

const PORT = "9001"

func main() {

	// 获取server 公钥 和私钥
	cert, _ := tls.LoadX509KeyPair("server.pem", "server.key")

	// 新建cert pool
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("ca.pem")
	// 将ca证书添加进去 用来校验client的身份
	certPool.AppendCertsFromPEM(ca)
	t := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})

	server := grpc.NewServer(grpc.Creds(t))
	proto.RegisterHelloWorldServer(server, &HelloWorldService{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	server.Serve(lis)
}
