package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
	"golang.org/x/sync/singleflight"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	mutex.Lock()
	mutex.Unlock()
	// mutex.Unlock() 不可多次解锁

	//写时，等待写锁释放，等待读锁释放
	//读时，等待写锁释放

	var rwmutex sync.RWMutex
	rwmutex.Lock()
	rwmutex.Unlock()
	rwmutex.RLock()
	rwmutex.RUnlock()

	//wait 可以同时多个goroutine等待同一个wait.wait() 等待结束
	var wait sync.WaitGroup
	wait.Add(2)
	go func() {
		defer wait.Done()
		time.Sleep(time.Second)
	}()

	go func() {
		defer wait.Done()
		time.Sleep(time.Second)
	}()

	go func() {
		wait.Wait()
		fmt.Println("wait 1111")
	}()
	go func() {
		wait.Wait()
		fmt.Println("wait 2222")
	}()

	time.Sleep(time.Second * 2)

	// once 保证只执行一次
	var once sync.Once
	once.Do(func() {
		fmt.Println("once1")
	})
	once.Do(func() {
		fmt.Println("once2")
	})

	//cond cond
	cond := sync.NewCond(&sync.Mutex{})
	for i := 0; i < 10; i++ {
		go func(id int) {
			cond.L.Lock()
			defer cond.L.Unlock()
			cond.Wait()
			fmt.Println("wait ", id)
		}(i)
	}
	time.Sleep(time.Second)

	cond.L.Lock()
	cond.Broadcast()
	cond.L.Unlock()
	time.Sleep(time.Second)

	//errgroup
	var eg errgroup.Group

	for i := 0; i < 10; i++ {

		eg.Go(func() error {
			id := rand.Intn(1)
			if id == 0 {
				return errors.New(fmt.Sprint("%d", id))
			}
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Println(err.Error())
	}

	//semaphore
	sem := semaphore.NewWeighted(10)
	sem.Acquire(context.Background(), 10)

	go func() {
		//优先请求5个，释放的时候先给这次请求，，这次请求的n不够的话，也不会给第二个请求的
		fmt.Println("请求5个", sem.Acquire(context.Background(), 5))
	}()

	time.Sleep(time.Second)
	go func() {
		fmt.Println("请求3个", sem.Acquire(context.Background(), 2))
	}()
	//释放4个
	sem.Release(4)
	time.Sleep(time.Second)
	sem.Release(3)

	time.Sleep(time.Second)

	//singleflight 并发多个请求只执行一次，
	g := singleflight.Group{}
	for i := 0; i < 100; i++ {
		go func(id int) {
			fmt.Println(g.Do("result", func() (interface{}, error) {
				return 1, nil
			}))
		}(i)

	}
	time.Sleep(time.Second)
}
