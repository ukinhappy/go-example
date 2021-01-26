package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	listen, err := net.Listen("tcp", "0.0.0.0:56022")
	if err != nil {
		log.Fatal("listen failed ", err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println("accept failed", err)
			continue
		}

		go func() {
			defer conn.Close()
			for {
				buf := make([]byte, 1028)
				n, err := conn.Read(buf)
				if err != nil {
					log.Fatal("read failed ", err)
				}
				fmt.Printf("读到的数据 %s ", buf[:n])
			}
		}()

	}
}
