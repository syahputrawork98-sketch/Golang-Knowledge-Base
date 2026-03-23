package main

import "fmt"

func main() {
	// Buffered channel dengan kapasitas 2
	ch := make(chan int, 2)

	// Kita bisa mengirim 2 data tanpa ada penerima yang stand-by
	ch <- 10
	ch <- 20
	fmt.Println("Buffer terisi: 2/2")

	// ch <- 30 // Ini akan DEADLOCK karena buffer penuh dan tidak ada penerima

	fmt.Println("Menerima dari buffer:")
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	
	fmt.Println("\nKesimpulan: Buffered channel mengurangi 'Tight Coupling' antara pengirim dan penerima.")
}
