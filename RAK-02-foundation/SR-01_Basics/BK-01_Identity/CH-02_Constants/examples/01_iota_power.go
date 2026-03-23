package main

import "fmt"

// Menggunakan iota untuk membuat urutan otomatis
const (
	StatusPending = iota // 0
	StatusActive         // 1
	StatusClosed         // 2
)

// Iota dengan pola bitwise (Shifting)
const (
	Read = 1 << iota // 1 (0001)
	Write           // 2 (0010)
	Execute         // 4 (0100)
)

func main() {
	fmt.Println("--- Constants & Iota in Go ---")
	fmt.Printf("Pending: %d, Active: %d, Closed: %d\n", StatusPending, StatusActive, StatusClosed)
	fmt.Printf("Permissions: Read(%d), Write(%d), Execute(%d)\n", Read, Write, Execute)
}
