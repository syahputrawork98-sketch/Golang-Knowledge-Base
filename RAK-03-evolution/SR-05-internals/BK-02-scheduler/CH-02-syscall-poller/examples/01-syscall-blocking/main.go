package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

// blockingSyscall simulasi pembacaan file yang memblokir M
func blockingSyscall(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	
	// Open file (syscall sinkron)
	f, err := os.Open("main.go")
	if err != nil {
		return
	}
	defer f.Close()

	// Simulasi kerja lama di level kernel
	time.Sleep(2 * time.Second)
	
	if id%50 == 0 {
		fmt.Printf("[G-%d] Finished blocking syscall\n", id)
	}
}

func main() {
	runtime.GOMAXPROCS(1) // Paksa hanya 1 Processor untuk melihat handoff
	fmt.Println("GOMAXPROCS set to 1")

	var wg sync.WaitGroup
	numG := 200 // 200 Blocking Goroutines

	fmt.Printf("Launching %d Goroutines performing BLOCKING syscalls...\n", numG)

	// Monitor thread OS secara real-time
	stopMonitor := make(chan bool)
	go func() {
		for {
			select {
			case <-stopMonitor:
				return
			default:
				numM, _ := runtime.ThreadCreateProfile(nil)
				fmt.Printf("--> Monitoring: OS Threads (M) = %d\n", numM)
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	for i := 0; i < numG; i++ {
		wg.Add(1)
		go blockingSyscall(i, &wg)
	}

	wg.Wait()
	stopMonitor <- true
	
	fmt.Println("\nLab Analysis:")
	fmt.Println("1. Perhatikan kenaikan OS Threads (M) meskipun P=1.")
	fmt.Println("2. Runtime (sysmon) mendeteksi M yang memblokir dan melepas P.")
	fmt.Println("3. Ini disebut 'Syscall Handoff'.")
}
