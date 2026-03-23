package main

import "fmt"

// Antarmuka Kontrak
type Speaker interface {
	Speak() string
}

type Human struct{ Name string }
func (h Human) Speak() string { return "Hello, I am " + h.Name }

type Robot struct{ Model string }
func (r Robot) Speak() string { return "Beep Boop, " + r.Model + " online" }

func main() {
	var s Speaker

	s = Human{"Antigravity"}
	fmt.Println(s.Speak())

	s = Robot{"T-800"}
	fmt.Println(s.Speak())
}
