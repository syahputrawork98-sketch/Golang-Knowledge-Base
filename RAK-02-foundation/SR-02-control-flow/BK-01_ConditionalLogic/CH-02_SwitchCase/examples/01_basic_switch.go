package main

import "fmt"

func main() {
	finger := 3

	fmt.Print("Jari ke-", finger, " adalah: ")
	
	// Switch di Go memiliki implisit break.
	// Tidak perlu menulis 'break' di setiap akhir case.
	switch finger {
	case 1:
		fmt.Println("Jempol")
	case 2:
		fmt.Println("Telunjuk")
	case 3:
		fmt.Println("Tengah")
	case 4:
		fmt.Println("Manis")
	case 5:
		fmt.Println("Kelingking")
	default:
		fmt.Println("Bukan jari manusia")
	}
}
