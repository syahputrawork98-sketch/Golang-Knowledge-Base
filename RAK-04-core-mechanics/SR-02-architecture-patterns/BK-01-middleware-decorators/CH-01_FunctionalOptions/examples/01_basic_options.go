package main

import "fmt"

// 01_basic_options.go - Pola Functional Options
// Analogi: Memesan Kopi Custom. Cukup sebutkan tambahan yang diinginkan.

type Server struct {
	Addr string
	Port int
}

// Option adalah fungsi yang memodifikasi Server
type Option func(*Server)

func WithPort(p int) Option {
	return func(s *Server) {
		s.Port = p
	}
}

func NewServer(addr string, opts ...Option) *Server {
	// Nilai default
	srv := &Server{
		Addr: addr,
		Port: 8080,
	}

	// Terapkan semua opsi yang diberikan user
	for _, opt := range opts {
		opt(srv)
	}

	return srv
}

func main() {
	// Server dengan port default
	s1 := NewServer("localhost")
	fmt.Printf("Server 1: %s:%d\n", s1.Addr, s1.Port)

	// Server dengan port custom via functional option
	s2 := NewServer("localhost", WithPort(9000))
	fmt.Printf("Server 2: %s:%d\n", s2.Addr, s2.Port)
}
