package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
)

func main() {
	sm := semaphore.NewWeighted(2)
	fmt.Println(sm.TryAcquire(1))
	fmt.Println(sm.TryAcquire(1))
	fmt.Println(sm.Acquire(context.TODO(),1))
}
