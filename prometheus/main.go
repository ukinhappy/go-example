package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shirou/gopsutil/mem"
)

func main() {

	http.Handle("/metrics", promhttp.Handler())

	memoryPercent := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "memory_percent",
		Help: "memory use percent",
	},
		[]string{"percent"},
	)

	counter := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "counter",
		Help: "counter of ",
	}, []string{"incry"})

	prometheus.MustRegister(memoryPercent)
	prometheus.MustRegister(counter)

	go func() {
		for {
			time.Sleep(time.Second * 1)
			//counter
			counter.WithLabelValues("useCounter").Inc()
			// gauge
			v, err := mem.VirtualMemory()
			if err != nil {
				panic(err)
			}
			memoryPercent.WithLabelValues("usedMemory").Set(v.UsedPercent)
			memoryPercent.WithLabelValues("usedMemory_new").Set(v.UsedPercent)
			fmt.Println(v.UsedPercent)
		}
	}()

	err := http.ListenAndServe("0.0.0.0:9100", nil)
	if err != nil {
		panic(err)
	}

}
