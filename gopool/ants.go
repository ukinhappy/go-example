package gopool

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"time"
)

func TestAnts() {
	fn := func(i interface{}) {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
	p, _ := ants.NewPoolWithFunc(1, fn)
	p.Invoke(1)
	p.Release()

}
