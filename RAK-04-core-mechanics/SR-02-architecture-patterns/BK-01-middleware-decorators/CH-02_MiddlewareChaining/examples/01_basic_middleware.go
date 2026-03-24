package main

import (
	"fmt"
	"net/http"
	"time"
)

// 01_basic_middleware.go - Interceptor Pattern (Middleware)
// Analogi: Pos pemeriksaan keamanan di Bandara.

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		// Lanjutkan ke handler berikutnya
		next.ServeHTTP(w, r)
		
		fmt.Printf("Logged: %s %s took %v\n", r.Method, r.URL.Path, time.Since(start))
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Gopher Hub!")
}

func main() {
	mux := http.NewServeMux()
	
	// Membungkus handler asli dengan middleware
	finalHandler := loggerMiddleware(http.HandlerFunc(helloHandler))
	
	mux.Handle("/", finalHandler)
	fmt.Println("Server running at :8080...")
	// http.ListenAndServe(":8080", mux) // Uncomment to run
}
