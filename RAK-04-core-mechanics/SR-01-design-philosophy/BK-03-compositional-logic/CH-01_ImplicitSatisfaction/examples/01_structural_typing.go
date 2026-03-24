package main

import "fmt"

// 01_structural_typing.go - Pemenuhan Interface Implisit
// Analogi: Stopkontak dan steker. Asal kaki tiga, pasti masuk.

// Kontrak (Interface)
type Speaker interface {
	Speak() string
}

// Implementasi 1: Robot
type Robot struct{}
func (r Robot) Speak() string { return "Beep Boop" }

// Implementasi 2: Human (Tanpa mendaftarkan diri ke Speaker)
type Human struct{}
func (h Human) Speak() string { return "Halo Dunia" }

func introduce(s Speaker) {
	fmt.Printf("Suara: %s\n", s.Speak())
}

func main() {
	r := Robot{}
	h := Human{}

	// Keduanya otomatis dianggap Speaker karena memiliki metode Speak()
	introduce(r)
	introduce(h)
}
