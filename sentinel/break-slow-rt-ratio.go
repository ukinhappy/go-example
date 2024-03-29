package main

import (
	"fmt"
	"log"
	"time"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/circuitbreaker"
	"github.com/alibaba/sentinel-golang/core/config"
	"github.com/alibaba/sentinel-golang/logging"
)
//
//type stateChangeTestListener struct {
//}
//
//func (s *stateChangeTestListener) OnTransformToClosed(prev circuitbreaker.State, rule circuitbreaker.Rule) {
//	fmt.Printf("rule.steategy: %+v, From %s to Closed, time: %d\n", rule.Strategy, prev.String(), util.CurrentTimeMillis())
//}
//
//func (s *stateChangeTestListener) OnTransformToOpen(prev circuitbreaker.State, rule circuitbreaker.Rule, snapshot interface{}) {
//	fmt.Printf("rule.steategy: %+v, From %s to Open, snapshot: %.2f, time: %d\n", rule.Strategy, prev.String(), snapshot, util.CurrentTimeMillis())
//}
//
//func (s *stateChangeTestListener) OnTransformToHalfOpen(prev circuitbreaker.State, rule circuitbreaker.Rule) {
//	fmt.Printf("rule.steategy: %+v, From %s to Half-Open, time: %d\n", rule.Strategy, prev.String(), util.CurrentTimeMillis())
//}

func main() {
	conf := config.NewDefaultConfig()
	// for testing, logging output to console
	conf.Sentinel.Log.Logger = logging.NewConsoleLogger()
	err := sentinel.InitWithConfig(conf)
	if err != nil {
		log.Fatal(err)
	}
	ch := make(chan struct{})
	// Register a state change listener so that we could observer the state change of the internal circuit breaker.
	//circuitbreaker.RegisterStateChangeListeners(&stateChangeTestListener{})

	_, err = circuitbreaker.LoadRules([]*circuitbreaker.Rule{
		// Statistic time span=5s, recoveryTimeout=3s, slowRtUpperBound=50ms, maxSlowRequestRatio=50%
		{
			Resource:                     "abc",
			Strategy:                     circuitbreaker.SlowRequestRatio,
			RetryTimeoutMs:               3000,
			MinRequestAmount:             5,
			StatIntervalMs:               1000,
			MaxAllowedRtMs:               50,
			Threshold:                    0.5,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	logging.Info("[CircuitBreaker SlowRtRatio] Sentinel Go circuit breaking demo is running. You may see the pass/block metric in the metric log.")
	go func() {
		for {
			e, b := sentinel.Entry("abc")
			if b != nil {
				fmt.Println("g1 blocker")
				// g1 blocked
				time.Sleep(time.Duration(100) * time.Millisecond)
			} else {
				// g1 passed
				fmt.Println("g1 passed")
				time.Sleep(time.Duration(100) * time.Millisecond)
				e.Exit()
			}
		}
	}()

	go func() {
		for {
			e, b := sentinel.Entry("abc")
			if b != nil {
				// g2 blocked
				fmt.Println(" g2 blocked")
				time.Sleep(time.Duration(100) * time.Millisecond)
			} else {
				// g2 passed
				fmt.Println("g2 passed")
				time.Sleep(time.Duration(100) * time.Millisecond)
				e.Exit()
			}
		}
	}()
	<-ch
}