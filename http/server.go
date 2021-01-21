package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("test")
		writer.WriteHeader(http.StatusAccepted)
		writer.WriteHeader(http.StatusOK)
		return
	})
	http.ListenAndServe("127.0.0.1:8080", nil)
}
