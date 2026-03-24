package main

import "fmt"

// 01_io_composition.go - Komposisi Interface (Swiss Army Knife)
// Analogi: Fitting Lampu yang bisa menerima berbagai jenis bohlam.

type Reader interface {
	Read() string
}

type Writer interface {
	Write(string)
}

// Komposisi Interface
type ReadWriter interface {
	Reader
	Writer
}

type Device struct{}

func (d Device) Read() string   { return "Data ditarik" }
func (d Device) Write(s string) { fmt.Printf("Menulis: %s\n", s) }

func process(rw ReadWriter) {
	msg := rw.Read()
	rw.Write(msg + " (Processed)")
}

func main() {
	d := Device{}
	process(d) // Device memenuhi ReadWriter secara otomatis
}
