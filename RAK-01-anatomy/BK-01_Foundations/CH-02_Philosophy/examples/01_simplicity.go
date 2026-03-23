package main

import "fmt"

// 1. Definisikan Interface kecil (Sub-Rak 6 Level Standard)
type Speaker interface {
	Speak() string
}

// 2. Struct dasar (Basis Komposisi)
type Animal struct {
	Name string
}

func (a Animal) Intro() string {
	return fmt.Sprintf("I am an animal named %s", a.Name)
}

// 3. Embedding: Dog "memiliki" Animal, bukan "keturunan" Animal secara hirarkis C++
type Dog struct {
	Animal // Anonymous embedding
	Breed  string
}

// Implementasi interface Speaker
func (d Dog) Speak() string {
	return "Woof! Woof!"
}

func main() {
	// Inisialisasi menggunakan komposisi
	myDog := Dog{
		Animal: Animal{Name: "Buddy"},
		Breed:  "Golden Retriever",
	}

	// Akses method dari Animal (delegasi otomatis)
	fmt.Println(myDog.Intro())

	// Akses method dari Dog itu sendiri
	fmt.Println(myDog.Speak())

	// Membuktikan fleksibilitas interface (Polimorfisme tanpa inheritance)
	var s Speaker = myDog
	fmt.Printf("The Speaker says: %s\n", s.Speak())
}
