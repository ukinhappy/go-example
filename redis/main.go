package main

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

// dial dials connections then returns pool.
func dial() *redis.Pool {
	dialOptions := []redis.DialOption{
		redis.DialConnectTimeout(time.Second),
		redis.DialReadTimeout(time.Second),
		redis.DialWriteTimeout(time.Second),
	}
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:6379", dialOptions...)
			if err != nil {
				return nil, err
			}

			return c, nil
		},
		MaxIdle:     1000,
		MaxActive:   1000,
		IdleTimeout: time.Second * 100,
	}
}


func main() {
	//lua()
	//snowflake()
	incrlua()
}

