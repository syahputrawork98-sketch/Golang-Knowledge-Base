package main

import "fmt"

type Shouter interface {
	Shout()
}

type Person struct{ Name string }

// Method ini menggunakan Pointer Receiver
func (p *Person) Shout() {
	fmt.Println("WAAAAAH!", p.Name)
}

func main() {
	var s Shouter
	p := Person{"Budi"}

	// s = p // ERROR: Person tidak memenuhi Shouter (karena Shout pakai pointer)
	
	s = &p // BENAR: Pointer ke Person memenuhi Shouter
	s.Shout()
}
