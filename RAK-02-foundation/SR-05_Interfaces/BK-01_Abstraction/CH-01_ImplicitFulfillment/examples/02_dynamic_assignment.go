package main

import (
	"fmt"
	"reflect"
)

type Worker interface {
	DoWork()
}

type Engineer struct{}
func (Engineer) DoWork() { fmt.Println("Writing code...") }

type Manager struct{}
func (Manager) DoWork() { fmt.Println("Meeting...") }

func main() {
	var w Worker

	// Engineer dipasang ke Worker
	w = Engineer{}
	fmt.Printf("Worker is currently: %v\n", reflect.TypeOf(w))
	w.DoWork()

	fmt.Println("--- Switching ---")

	// Manager dipasang ke Worker yang sama
	w = Manager{}
	fmt.Printf("Worker is currently: %v\n", reflect.TypeOf(w))
	w.DoWork()
}
