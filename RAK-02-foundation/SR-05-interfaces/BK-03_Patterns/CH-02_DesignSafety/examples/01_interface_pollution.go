package main

import "fmt"

// ANTI-PATTERN: Membuat interface untuk setiap struct secara prematur
type Greeter interface {
	SayHello()
}

type MyGreeter struct{}
func (MyGreeter) SayHello() { fmt.Println("Hello World") }

func main() {
	// Jika Anda hanya punya satu implementasi, interface ini "Pollution"
	var g Greeter = MyGreeter{}
	g.SayHello()

	fmt.Println("Gunakan struct langsung saja jika tidak butuh polimorfisme/mocking.")
}
