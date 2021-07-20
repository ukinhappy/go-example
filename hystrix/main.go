package main

import (
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	metricCollector "github.com/afex/hystrix-go/hystrix/metric_collector"
	"github.com/afex/hystrix-go/plugins"
	"net"
	"net/http"
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
	collector, err := plugins.NewDatadogCollector("localhost:8125", "/metric")
	if err != nil {
		panic(err)
	}
	metricCollector.Registry.Register(collector)
	select {

	}
	return
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	go http.ListenAndServe(net.JoinHostPort("0.0.0.0", "8081"), hystrixStreamHandler)

	select {}
}
