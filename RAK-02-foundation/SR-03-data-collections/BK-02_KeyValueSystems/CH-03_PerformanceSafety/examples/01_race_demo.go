package main

import (
	"fmt"
	"time"
)

func main() {
	m := make(map[int]int)

	fmt.Println("Memulai penulisan bersamaan tanpa pengaman...")
	
	// Thread 1: Menulis terus menerus
	go func() {
		for {
			m[1] = 1
		}
	}()

	// Thread 2: Menulis terus menerus
	go func() {
		for {
			m[2] = 2
		}
	}()

	// Tunggu sebentar hingga crash terjadi
	time.Sleep(1 * time.Second)
	fmt.Println("Jika Anda melihat ini, berarti Anda beruntung (atau CPU Anda lambat).")
}
