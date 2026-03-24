package main

/*
#include <stdio.h>

// Deklarasi fungsi C yang akan memanggil fungsi Go
extern void go_callback();

static inline void call_go() {
    printf("C: Memanggil fungsi Go...\n");
    go_callback();
}
*/
import "C"
import "fmt"

//export go_callback
func go_callback() {
	fmt.Println("Go: Halo dari dunia Go! (Callback berhasil)")
}

func main() {
	fmt.Println("Main: Memulai eksekusi Cgo...")
	C.call_go()
	fmt.Println("Main: Eksekusi selesai.")
}
