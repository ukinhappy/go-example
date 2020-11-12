package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"strconv"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{Endpoints: []string{"127.0.0.1:2379"}, DialTimeout: time.Second * 5})
	if err != nil {
		panic(err)
	}
	kv := clientv3.NewKV(cli)
	_, err = kv.Put(context.Background(), "/happy/go-example/etcd1", "second")
	if err != nil {
		panic(err)
	}

	resp, err := kv.Get(context.Background(), "/happy/go-example/etcd", clientv3.WithPrefix())
	if err != nil {
		panic(err)
	}

	for _, v := range resp.Kvs {
		fmt.Println(string(v.Key), string(v.Value))
	}

	lease := clientv3.NewLease(cli)

	grant, err := lease.Grant(context.Background(), 10)
	if err != nil {
		panic(err)
	}

	kv.Put(context.Background(), "/happy/go-example/etcd/lease", "lease", clientv3.WithLease(grant.ID))

	go func() {
		tm := time.NewTicker(time.Second * 9)
		for true {
			select {
			case <-tm.C:
				lease.KeepAliveOnce(context.Background(), grant.ID)

			}
		}
	}()

	go func() {
		tm := time.NewTicker(time.Second * 1)
		for true {
			select {
			case t := <-tm.C:
				resp, err := kv.Get(context.Background(), "/happy/go-example/etcd/lease")
				if err != nil {
					panic(err)
				}
				for _, v := range resp.Kvs {
					fmt.Println(t.Second(), string(v.Value), string(v.Key))
				}
			}
		}
	}()

	go func() {
		tm := time.NewTicker(time.Second * 10)
		for true {
			select {
			case t := <-tm.C:
				_, err := kv.Put(context.Background(), "/happy/go-example/etcd/watch", strconv.Itoa(t.Second()))
				if err != nil {
					panic(err)
				}

			}
		}
	}()

	go func() {
		watch := clientv3.NewWatcher(cli).Watch(context.Background(), "/happy/go-example/etcd/watch")

		for true{
			select {
			case r := <-watch:
				for _, w := range r.Events {
					fmt.Println("watch update", string(w.Kv.Key), string(w.Kv.Value))
				}
			}
		}

	}()

	select {}

}
