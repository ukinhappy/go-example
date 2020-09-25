package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,
		"Hi, This is an example of https service in golang!")
}
func main() {
	http.HandleFunc("/", handler)
	fmt.Println(http.ListenAndServeTLS("127.0.0.1:8081", "server.crt", "server.key", nil))
}
