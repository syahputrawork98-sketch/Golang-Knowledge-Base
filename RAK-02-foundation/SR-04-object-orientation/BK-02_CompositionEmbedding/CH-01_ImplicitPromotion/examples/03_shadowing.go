package main

import "fmt"

type Base struct {
	Tag string
}

type Container struct {
	Base
	Tag string // Shadowing: Nama yang sama dengan Base
}

func main() {
	c := Container{
		Base: Base{Tag: "BASE_TAG"},
		Tag:  "CONTAINER_TAG",
	}

	// 1. Selector 'Tag' akan merujuk ke level terdekat (Container)
	fmt.Println("c.Tag:", c.Tag) 

	// 2. Untuk mengakses Tag milik Base, harus eksplisit
	fmt.Println("c.Base.Tag:", c.Base.Tag)
}
