package util

import (
	"crypto/tls"
	"golang.org/x/net/http2"
)

func GetTLSConfig(certPemPath, certKeyPath string) *tls.Config {
	cert, _ := tls.LoadX509KeyPair(certPemPath, certKeyPath)
	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		NextProtos:   []string{http2.NextProtoTLS},
	}
}
