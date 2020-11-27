package main

import "fmt"
import "github.com/gomodule/redigo/redis"

func lua() {
	p := dial()
	// --------
	fmt.Println(redis.String(p.Get().Do("get", "name")))

	p.Get().Do("set", "name", "happy")
	fmt.Println(redis.String(p.Get().Do("get", "name")))

	// --------
	var luatxt = `return redis.call('set',KEYS[1],'happylua')`
	p.Get().Do("eval", luatxt, 1, "name")
	fmt.Println(redis.String(p.Get().Do("get", "name")))
}
