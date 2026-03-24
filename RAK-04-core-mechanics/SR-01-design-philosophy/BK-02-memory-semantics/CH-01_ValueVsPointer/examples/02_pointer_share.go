package main

import (
	"fmt"
)

// 02_pointer_share.go - Pointer Semantics (Efficiency & Mutation)
// Analogi: Memberikan alamat rumah. Perubahan warna cat rumah terlihat oleh semua pengirim.

type Player struct {
	HP  int
}

func healing(p *Player) {
	// p di sini mereferensikan alamat memori yang sama
	p.HP += 50
	fmt.Printf("Healing done! Current HP: %d\n", p.HP)
}

func main() {
	p1 := Player{HP: 10}
	
	fmt.Printf("P1 HP Before: %d\n", p1.HP)
	healing(&p1) // Mengirim alamat memori (pointer)
	fmt.Printf("P1 HP After: %d\n", p1.HP) // Nilai berubah
}
