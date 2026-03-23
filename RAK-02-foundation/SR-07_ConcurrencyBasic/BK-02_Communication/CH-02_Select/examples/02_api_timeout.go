package main

import (
	"fmt"
	"time"
)

func CallSlowAPI(ch chan string) {
	// Simulasi API lambat yang butuh 5 detik
	time.Sleep(5 * time.Second)
	ch <- "API RESPONSE: DATA USER OK"
}

func main() {
	resCh := make(chan string)
	go CallSlowAPI(resCh)

	fmt.Println("Memanggil API... Menunggu maksimal 2 detik.")

	select {
	case res := <-resCh:
		fmt.Println("Suksess:", res)
	case <-time.After(2 * time.Second):
		fmt.Println("FAILED: API terlalu lambat (Timeout exceeded!)")
	}
}
