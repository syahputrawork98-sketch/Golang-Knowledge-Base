package main

import (
	"fmt"
	"runtime"
	"runtime/metrics"
)

func main() {
	// 1. Definisikan metrik yang ingin dibaca
	const (
		memStats = "/memory/classes/heap/objects:bytes"
		goroutines = "/sched/goroutines:goroutines"
	)

	// Persiapkan slice sample
	samples := make([]metrics.Sample, 2)
	samples[0].Name = memStats
	samples[1].Name = goroutines

	// 2. Baca metrik dari runtime
	metrics.Read(samples)

	// 3. Proses hasil
	for _, sample := range samples {
		name := sample.Name
		value := sample.Value

		switch value.Kind() {
		case metrics.KindUint64:
			fmt.Printf("%s: %d\n", name, value.Uint64())
		case metrics.KindFloat64:
			fmt.Printf("%s: %f\n", name, value.Float64())
		case metrics.KindFloat64Histogram:
			fmt.Printf("%s: (Histogram data)\n", name)
		default:
			fmt.Printf("%s: unknown kind\n", name)
		}
	}

	// Simulasi pembuatan goroutine dan baca ulang
	fmt.Println("\nCreating 10 goroutines...")
	for i := 0; i < 10; i++ {
		go func() { select {} }()
	}
	runtime.Gosched()

	metrics.Read(samples)
	fmt.Printf("Updated %s: %d\n", goroutines, samples[1].Value.Uint64())
}
