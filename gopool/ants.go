package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
	"time"
)

func main() {
	fn := func(i interface{}) {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
	p, _ := ants.NewPoolWithFunc(1, fn)
	p.Release()

	wait := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wait.Add(1)
		go func(idx int) {
			defer wait.Done()
			p.Invoke(idx)
		}(i)
	}
	wait.Wait()


}
