package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:56022")
	if err != nil {
		log.Fatal(err)
	}

	var i int
	for {
		_, err := conn.Write([]byte(fmt.Sprintf("hello %d", i)))
		if err != nil {
			log.Fatal("write failed ",err)
		}
		i++
		time.Sleep(time.Second)
	}

}
