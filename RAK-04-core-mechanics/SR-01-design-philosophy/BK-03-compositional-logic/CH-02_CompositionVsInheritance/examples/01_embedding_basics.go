package main

import "fmt"

// 01_embedding_basics.go - Struct Embedding (Composition)
// Analogi: Memasang komponen mesin ke dalam rangka mobil.

type Engine struct {
	HorsePower int
}

func (e Engine) Start() {
	fmt.Printf("Engine starts with %d HP!\n", e.HorsePower)
}

// Car "mempunyai" Engine, bukan "adalah" Engine.
type Car struct {
	Engine // Embedded Type
	Model  string
}

func main() {
	myCar := Car{
		Engine: Engine{HorsePower: 300},
		Model:  "Tesla Model S",
	}

	// Method Promotion: Kita bisa memanggil Start langsung dari Car
	myCar.Start()
	fmt.Printf("My %s is ready to go!\n", myCar.Model)
}
