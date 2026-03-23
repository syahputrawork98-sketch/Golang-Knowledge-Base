package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// Non-blocking Receive
	select {
	case msg := <-messages:
		fmt.Println("Pesan diterima:", msg)
	default:
		fmt.Println("Tidak ada pesan masuk (Main melanjutkan pekerjaan lain)")
	}

	// Non-blocking Send
	msg := "HI"
	select {
	case messages <- msg:
		fmt.Println("Pesan terkirim:", msg)
	default:
		fmt.Println("Pesan TIDAK terkirim (Tidak ada penerima)")
	}

	// Multi-way Non-blocking
	select {
	case msg := <-messages:
		fmt.Println("Pesan diterima:", msg)
	case sig := <-signals:
		fmt.Println("Sinyal diterima:", sig)
	default:
		fmt.Println("Eksplorasi selesai: Tidak ada aktivitas.")
	}
}
