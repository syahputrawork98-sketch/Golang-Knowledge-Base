package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // Aktifkan pprof via HTTP
	"time"
)

// leakedWorker simulates a worker that leaks
func leakedWorker(ch <-chan int) {
	// Goroutine ini akan menunggu selamanya karena ch tidak pernah dikirim
	// dan tidak pernah ditutup dalam simulasi ini.
	for val := range ch {
		fmt.Println(val)
	}
}

func main() {
	// Jalankan pprof server di background
	go func() {
		log.Println("Pprof server started at http://localhost:6060/debug/pprof")
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	fmt.Println("Simulating Goroutine leaks...")
	
	for i := 0; i < 100; i++ {
		ch := make(chan int)
		go leakedWorker(ch) // Membuat 100 goroutine yang bocor
		time.Sleep(10 * time.Millisecond)
	}

	fmt.Println("100 leaked goroutines created.")
	fmt.Println("Please open: http://localhost:6060/debug/pprof/goroutine?debug=1")
	fmt.Println("Or run: go tool pprof http://localhost:6060/debug/pprof/goroutine")
	
	// Tunggu agar user bisa inspect
	select {}
}
