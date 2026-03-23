package main

import (
	"fmt"
	"time"
)

func main() {
	// Unbuffered channel
	ch := make(chan string)

	go func() {
		fmt.Println("Goroutine: Menyiapkan data...")
		time.Sleep(2 * time.Second)
		fmt.Println("Goroutine: Mencoba mengirim data ke channel...")
		ch <- "RAHASIA NEGARA" // Akan BLOCKING di sini sampai main menerima
		fmt.Println("Goroutine: Data berhasil diterima oleh main!")
	}()

	fmt.Println("Main: Melakukan pekerjaan lain...")
	time.Sleep(4 * time.Second)
	
	fmt.Println("Main: Siap menerima data dari channel...")
	msg := <-ch // Main akan BLOCKING di sini sampai goroutine mengirim
	
	fmt.Printf("Main: Pesan diterima: %s\n", msg)
}
