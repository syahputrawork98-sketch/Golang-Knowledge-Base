package main

import "fmt"

type Animal interface {
	MakeSound()
}

type Cat struct{}
func (Cat) MakeSound() { fmt.Println("Meow!") }

type Dog struct{}
func (Dog) MakeSound() { fmt.Println("Woof!") }

// Fungsi ini tidak peduli apa tipenya, selama dia adalah Animal
func AppreciateAnimal(a Animal) {
	fmt.Print("Animal says: ")
	a.MakeSound()
}

func main() {
	zoo := []Animal{Cat{}, Dog{}}

	for _, a := range zoo {
		AppreciateAnimal(a)
	}
}
