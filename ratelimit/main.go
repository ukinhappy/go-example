package main

import (
	"fmt"
	"github.com/juju/ratelimit"
	"time"
)

func main() {
	b := ratelimit.NewBucketWithRate(10, 10)

	time.Sleep(time.Second)
	fmt.Println("----------")
	fmt.Println(b.TakeAvailable(1))
	fmt.Println(b.TakeAvailable(8))
	fmt.Println(b.TakeAvailable(1))
	fmt.Println(b.TakeAvailable(1))
	fmt.Println(b.Available())
	fmt.Println(b.Rate())
	fmt.Println(b.Capacity())

	fmt.Println(b.Take(1))
	fmt.Println(b.Wait )

}
