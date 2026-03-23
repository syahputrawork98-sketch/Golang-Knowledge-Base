package main

import (
	"fmt"
	"slices"
)

type Data struct {
	ID int
}

func main() {
	// Slice berisi pointer ke objek
	items := []*Data{
		{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4},
	}

	fmt.Println("Initial size:", len(items))

	// Hapus ID 2 dan 3
	items = slices.DeleteFunc(items, func(d *Data) bool {
		return d.ID == 2 || d.ID == 3
	})

	fmt.Println("Final size:", len(items))
	
	// Mekanisme Internal: slices.DeleteFunc secara otomatis men-set 
	// elemen sisa di backing array menjadi nil agar objek ID 2 & 3 
	// bisa di-GC meskipun backing array slice 'items' masih ada.
	
	fmt.Println("Selesai. Go modern menjamin keamanan memori pada operasi fungsional.")
}
