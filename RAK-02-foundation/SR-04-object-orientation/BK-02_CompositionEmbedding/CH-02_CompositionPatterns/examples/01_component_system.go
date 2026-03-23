package main

import "fmt"

// Komponen 1: Kesehatan
type Health struct {
	MaxHP, CurrentHP int
}

func (h *Health) TakeDamage(amt int) {
	h.CurrentHP -= amt
	if h.CurrentHP < 0 {
		h.CurrentHP = 0
	}
}

// Komponen 2: Posisi
type Position struct {
	X, Y float64
}

// Entitas Utama: Player (Menggabungkan Komponen)
type Player struct {
	Name string
	Health // Embedded Component
	Pos    Position // Regular Component
}

func main() {
	p := Player{
		Name:   "Knight",
		Health: Health{MaxHP: 100, CurrentHP: 100},
		Pos:    Position{X: 0, Y: 0},
	}

	// 1. Memanggil method dari embedded Health
	p.TakeDamage(20)
	
	// 2. Mengakses field regular Position
	p.Pos.X = 10.5

	fmt.Printf("Player %s at (%.1f, %.1f) has HP: %d/%d\n", 
		p.Name, p.Pos.X, p.Pos.Y, p.CurrentHP, p.MaxHP)
}
