package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "PESAN DARI TIM A"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "PESAN DARI TIM B"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Diterima:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Diterima:", msg2)
		}
	}
}
