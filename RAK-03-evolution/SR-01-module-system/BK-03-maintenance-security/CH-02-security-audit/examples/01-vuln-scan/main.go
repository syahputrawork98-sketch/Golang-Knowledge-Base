package main

import (
	"fmt"
	"golang.org/x/text/language"
)

func main() {
	tag := language.MustParse("id")
	fmt.Printf("Language Tag: %s\n", tag)
}
