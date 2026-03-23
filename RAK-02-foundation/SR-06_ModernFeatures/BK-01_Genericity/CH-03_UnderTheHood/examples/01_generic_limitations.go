package main

import "fmt"

type User struct {
	Name string
}
func (u User) Greet() { fmt.Println("Hello", u.Name) }

// ERROR: T tidak dijamin punya method Greet()
// func SayHi[T any](val T) {
//     val.Greet() 
// }

// SOLUSI: Gunakan Constraint Interface
type Greeter interface {
	Greet()
}

func SayHiSafe[T Greeter](val T) {
	val.Greet()
}

func main() {
	u := User{Name: "Antigravity"}
	SayHiSafe(u)
	
	fmt.Println("\nKesimpulan: Generics tetap butuh Interface jika ingin memanggil method.")
}
