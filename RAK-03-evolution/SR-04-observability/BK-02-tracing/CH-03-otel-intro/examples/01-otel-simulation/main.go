package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
)

// TraceHeaderName standar W3C
const TraceHeaderName = "traceparent"

// generateTraceID membuat ID trace ala OTel
func generateTraceID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

// ServiceB bertindak sebagai downstream service
func serviceB(ctx context.Context) {
	traceID := ctx.Value("trace_id")
	fmt.Printf("[Service B] Menerima request dengan TraceID: %v\n", traceID)
	fmt.Println("[Service B] Memproses data ke DB...")
}

// ServiceA bertindak sebagai upstream service (Gateway)
func serviceA() {
	// 1. Generate ID baru di entry point
	traceID := generateTraceID()
	fmt.Printf("[Service A] Memulai request baru. TraceID: %s\n", traceID)

	// 2. Simpan ke context (simulasi propagasi)
	ctx := context.WithValue(context.Background(), "trace_id", traceID)

	// 3. Panggil downstream
	fmt.Println("[Service A] Memanggil Service B...")
	serviceB(ctx)
}

func main() {
	log.Println("Simulasi Distributed Tracing (OpenTelemetry Concept)")
	log.Println("-----------------------------------------------------")
	
	serviceA()

	log.Println("-----------------------------------------------------")
	log.Println("Konsep: TraceID diteruskan melalui context (lokal) ")
	log.Println("atau melalui HTTP Header 'traceparent' (antar jaringan).")
}
