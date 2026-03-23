package main

import (
	"fmt"
	"strings"
)

// MyString adalah alias tipe (User-defined type)
type MyString string

// Kita bisa menempelkan method pada tipe bentukan apa saja 
// asalkan berada dalam package yang sama.
func (s MyString) IsUpper() bool {
	return string(s) == strings.ToUpper(string(s))
}

func main() {
	s1 := MyString("HELLO")
	s2 := MyString("Hello")

	fmt.Printf("Is '%s' upper? %v\n", s1, s1.IsUpper())
	fmt.Printf("Is '%s' upper? %v\n", s2, s2.IsUpper())
}
