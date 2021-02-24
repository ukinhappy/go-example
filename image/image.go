package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

func main() {
	f, err := os.Open("/Users/ukinhappy/code/goproject/src/go-example/image/baidu.jpg")
	if err != nil {
		log.Fatal(err)
	}
	im, s, err := image.DecodeConfig(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)
	fmt.Println(im)
}
