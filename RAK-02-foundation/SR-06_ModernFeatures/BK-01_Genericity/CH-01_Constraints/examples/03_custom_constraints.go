package main

import "fmt"

type MyString string

// Constraint dengan ~ memungkinkan tipe alias masuk
type Stringish interface {
	~string
}

func PrintWrapped[T Stringish](val T) {
	fmt.Printf("<<< %s >>>\n", val)
}

func main() {
	// 1. String asli
	PrintWrapped("Hello Go")

	// 2. Tipe alias (MyString)
	// Jika tanpa ~, baris di bawah akan error
	var s MyString = "Alias Power"
	PrintWrapped(s)
}
