package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Definisikan metrik custom
var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})

	responseTime = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "myapp_response_time_seconds",
		Help:    "Response time in seconds",
		Buckets: prometheus.DefBuckets,
	})
)

func main() {
	// Simulasi event secara konkuren
	go func() {
		for {
			opsProcessed.Inc()
			
			// Simulasi latensi acak
			start := time.Now()
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			responseTime.Observe(time.Since(start).Seconds())
			
			time.Sleep(1 * time.Second)
		}
	}()

	// 1. Ekspos endpoint /metrics standar Prometheus
	http.Handle("/metrics", promhttp.Handler())

	log.Println("Prometheus Exposer started at :2112/metrics")
	log.Fatal(http.ListenAndServe(":2112", nil))
}
