package main

import "fmt"

// Antarmuka yang didefinisikan oleh sistem
type Transporter interface {
	Move(distance int) string
}

// Struct 1: Car
type Car struct {
	Brand string
}
func (c Car) Move(d int) string {
	return fmt.Sprintf("Car %s driving %d km", c.Brand, d)
}

// Struct 2: Airplane
type Airplane struct {
	Model string
}
func (a Airplane) Move(d int) string {
	return fmt.Sprintf("Airplane %s flying %d km", a.Model, d)
}

func main() {
	// 1. Variabel Interface
	var transport Transporter

	// Car memenuhi Transporter secara implisit
	transport = Car{"Toyota"}
	fmt.Println(transport.Move(50))

	// Airplane memenuhi Transporter secara implisit
	transport = Airplane{"Boeing 747"}
	fmt.Println(transport.Move(1000))

	fmt.Println("\nKesimpulan: Interface tidak peduli APA tipenya, tapi BISA APA tipenya.")
}
