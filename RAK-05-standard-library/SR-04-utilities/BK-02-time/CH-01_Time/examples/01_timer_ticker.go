package main

import (
	"fmt"
	"time"
)

// 01_timer_ticker.go - Penjadwalan Tugas (Async Time)
// Analogi: Alarm (Timer) vs Lonceng Gereja (Ticker).

func main() {
	// 1. Timer: Alarm satu kali (One-off)
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("Timer dimulai (tunggu 2 detik)...")
	<-timer.C
	fmt.Println("Timer berbunyi!")

	// 2. Ticker: Lonceng berulang (Recurring)
	ticker := time.NewTicker(500 * time.Millisecond)
	fmt.Println("Ticker dimulai (berbunyi setiap 0.5 detik)...")
	
	done := make(chan bool)
	go func() {
		time.Sleep(2 * time.Second)
		done <- true
	}()

	for {
		select {
		case t := <-ticker.C:
			fmt.Println("Ticker at:", t.Format("15:04:05.000"))
		case <-done:
			ticker.Stop()
			fmt.Println("Ticker dihentikan.")
			return
		}
	}
}
