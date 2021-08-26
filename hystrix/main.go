package main

import (
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"time"
)

func main() {

	hystrix.ConfigureCommand("my_command", hystrix.CommandConfig{
		Timeout:               900,
		MaxConcurrentRequests: 100,
		ErrorPercentThreshold: 25,
	})

	for i := 0; i < 120; i++ {
		go func() {
			hystrix.Go("my_command", func() error {
				time.Sleep(time.Second)
				return nil
			}, func(err error) error {
				fmt.Println(err)
				return err
			})

		}()
	}

	select {}
}
