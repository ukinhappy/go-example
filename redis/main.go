package main

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

// dial dials connections then returns pool.
func dial(addr, pw string) *redis.Pool {
	dialOptions := []redis.DialOption{
		redis.DialConnectTimeout(time.Second),
		redis.DialReadTimeout(time.Second),
		redis.DialWriteTimeout(time.Second),
	}
	if pw != "" {
		dialOptions = append(dialOptions, redis.DialPassword(pw))
	}
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr, dialOptions...)
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
	//incrlua()

}


