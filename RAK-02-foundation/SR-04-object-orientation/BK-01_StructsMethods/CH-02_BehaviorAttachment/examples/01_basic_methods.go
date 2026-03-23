package main

import "fmt"

type Rectangle struct {
	Width, Height float64
}

// Method dengan Value Receiver
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Method dengan Pointer Receiver
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}

	fmt.Println("Initial Area:", rect.Area())

	// Memanggil pointer receiver pada variable value
	// Go otomatis melakukan (&rect).Scale(2)
	rect.Scale(2)

	fmt.Println("New Width:", rect.Width)
	fmt.Println("New Area:", rect.Area())
}
