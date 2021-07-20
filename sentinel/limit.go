package main

import (
	"fmt"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/flow"
	"time"
)

func initSentinel() {
	err := sentinel.InitWithConfigFile("/Users/ukinhappy/code/goproject/src/go-example/sentinel/cfg.yml")
	if err != nil {
		panic(err)
	}
}

func main() {
	initSentinel()
	qpsLimit()
	time.Sleep(time.Minute)
}

func qpsLimit() {

	_, err := flow.LoadRules([]*flow.Rule{
		{
			Resource:               "sentinel-go-demo",
			Threshold:              5,
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject,
		},
	})
	if err != nil {
		panic(err)
	}


	for i := 0; i < 6; i++ {
		go func() {
			// Entry 方法用于埋点
			e, b := sentinel.Entry("sentinel-go-demo", sentinel.WithTrafficType(base.Inbound))
			if b != nil {
				fmt.Println(b.Error())
			} else {
				fmt.Println("success")
				e.Exit()
			}
		}()
	}

}