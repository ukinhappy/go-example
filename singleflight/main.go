package main

import (
	"fmt"
	"golang.org/x/sync/singleflight"
)

func main() {
	sg := singleflight.Group{}
	sg.Do("1", func() (interface{}, error) {
		fmt.Println(111)
		return nil, nil
	})
	sg.Do("1", func() (interface{}, error) {
		fmt.Println(111)
		return nil, nil
	})
}
