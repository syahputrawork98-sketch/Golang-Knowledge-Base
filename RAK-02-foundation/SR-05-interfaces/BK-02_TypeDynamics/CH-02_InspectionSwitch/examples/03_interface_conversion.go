package main

import "fmt"

type Reader interface {
	Read()
}

type Closer interface {
	Close()
}

type File struct{}
func (File) Read()  { fmt.Println("Reading file...") }
func (File) Close() { fmt.Println("Closing file...") }

func main() {
	var r Reader = File{}

	// Assertion dari Interface Reader ke Interface Closer!
	// Ini membuktikan bahwa interface ke interface bisa dilakukan
	// selama tipe konkrit aslinya memenuhi kedua kontrak.
	if c, ok := r.(Closer); ok {
		fmt.Println("Objek ini juga bisa ditutup (memenuhi Closer)")
		c.Close()
	} else {
		fmt.Println("Objek ini tidak punya method Close()")
	}
}
