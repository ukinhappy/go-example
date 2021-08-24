package main

import (
	"fmt"
	"github.com/juju/ratelimit"
	"time"
)

func main() {
	b := ratelimit.NewBucketWithRate(1, 1)
	fmt.Println(b.TakeAvailable(1))
	time.Sleep(time.Second)
	fmt.Println(b.TakeAvailable(1))


	return
	b.Take(1)
	fmt.Println(b.Available())
	b.TakeAvailable(1)
	b.TakeMaxDuration(1, time.Second)
	b.Wait(1)
	b.WaitMaxDuration(1, time.Second)
	b.Available()

	fmt.Println(b.TakeAvailable(11))
	fmt.Println(b.TakeAvailable(1))
	fmt.Println(b.TakeAvailable(1))
	fmt.Println(b.Available())
	fmt.Println(b.Rate())
	fmt.Println(b.Capacity())

	fmt.Println(b.Take(1))

}
