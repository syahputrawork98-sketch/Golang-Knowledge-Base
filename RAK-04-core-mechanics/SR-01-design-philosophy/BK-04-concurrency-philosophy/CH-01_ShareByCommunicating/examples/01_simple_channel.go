package main

import (
	"fmt"
	"time"
)

// 01_simple_channel.go - Dasar Komunikasi Channel (CSP)
// Analogi: Ban berjalan di dapur restoran.

func main() {
	// Membuat channel untuk mengirim pesan string
	messages := make(chan string)

	// Goroutine 1: Pengirim (Producer)
	go func() {
		fmt.Println("Producer: Menyiapkan laporan...")
		time.Sleep(1 * time.Second)
		messages <- "Laporan Bulanan Siap!" // Mengirim ke channel
	}()

	// Goroutine Utama: Penerima (Consumer)
	fmt.Println("Main: Menunggu laporan dari channel...")
	msg := <-messages // Menunggu hingga data tersedia
	fmt.Printf("Main: Pesan diterima: %s\n", msg)
}
