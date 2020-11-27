package main

import (
	"log"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"

	"crypto/tls"
	pb "go-example/grpc/pkg/proto"
	"go-example/grpc/pkg/util"
	"net/http"
)

type helloService struct{}

func NewHelloService() *helloService {
	return &helloService{}
}

func (h helloService) SayHelloWorld(ctx context.Context, r *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
	return &pb.HelloWorldResponse{
		Message: "happy",
	}, nil
}

func createInternalServer(tlsConfig *tls.Config) (*http.Server) {

	// grpc
	grpcServer := grpc.NewServer(grpc.Creds(credentials.NewTLS(tlsConfig)))

	// register grpc pb
	pb.RegisterHelloWorldServer(grpcServer, NewHelloService())

	// gw server
	ctx := context.Background()
	dcreds, err := credentials.NewClientTLSFromFile(CertPemPath, CertName)
	if err != nil {
		log.Printf("Failed to create client TLS credentials %v", err)
	}
	dopts := []grpc.DialOption{grpc.WithTransportCredentials(dcreds)}
	gwmux := runtime.NewServeMux()

	// register grpc-gateway pb
	if err := pb.RegisterHelloWorldHandlerFromEndpoint(ctx, gwmux, EndPoint, dopts); err != nil {
		log.Printf("Failed to register gw server: %v\n", err)
	}

	// http服务
	mux := http.NewServeMux()
	mux.Handle("/", gwmux)

	return &http.Server{
		Addr:      EndPoint,
		Handler:   util.GrpcHandlerFunc(grpcServer, mux),
		TLSConfig: tlsConfig,
	}
}

var (
	CertName    string = "grpc-gateway"
	CertPemPath string = "server.pem"
	CertKeyPath string = "server.key"
	EndPoint    string = ":8000"
)

// grpc 服务和http服务共用一个端口号 必须使用https
func main() {
	conn, err := net.Listen("tcp", EndPoint)
	if err != nil {
		log.Printf("TCP Listen err:%v\n", err)
	}

	// 获取服务端的公钥和私钥
	tlsConfig := util.GetTLSConfig(CertPemPath, CertKeyPath)

	//
	srv := createInternalServer(tlsConfig)
	if err = srv.Serve(tls.NewListener(conn, tlsConfig)); err != nil {
		log.Printf("ListenAndServe: %v\n", err)
	}
}
