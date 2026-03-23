package main

import "fmt"

type NormalMover struct{}
func (NormalMover) Move() { fmt.Println("Walking slowly...") }

type FastMover struct{}
func (FastMover) Move() { fmt.Println("Running fast!") }

// Robot bisa memiliki strategi pergerakan yang berbeda
type Robot struct {
	Name string
	// Kita bisa menanamkan "Strategy" pergerakan
	NormalMover 
}

func main() {
	r := Robot{Name: "R2D2"}
	r.Move() // Menggunakan NormalMover secara default

	fmt.Println("Kesimpulan: Dengan komposisi, kita bisa merakit fungsionalitas " +
		"layaknya potongan Lego.")
}
