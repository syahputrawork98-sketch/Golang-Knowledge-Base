# Rak 11: Distributed Systems

Menguasai Go dalam arsitektur sistem terdistribusi yang tangguh dan berperforma tinggi.

## Struktur Sub-Rak
- **[SR-01-GRPC](./SR-01-GRPC/README.md)**: Komunikasi antar layanan menggunakan Protocol Buffers.
- **[SR-02-MICROSERVICES-PATTERNS](./SR-02-MICROSERVICES-PATTERNS/README.md)**: Circuit Breaker, Service Discovery, dan API Gateway dengan Go.

## The Power of Go in Systems
Go adalah bahasa yang "didikte oleh kebutuhan cloud":
1. **CSP Model**: Penggunaan Channel menyederhanakan koordinasi data antar layanan.
2. **Built-in Testing**: Tooling internal Go memudahkan pembuatan unit test untuk layanan skala besar.
3. **Observation**: Integrasi mudah dengan Prometheus & OpenTelemetry (Observability).

---

## 🌐 Architectural Nexus (Jembatan Master Plan)
Materi di Rak ini terhubung langsung dengan komponen lain dalam ekosistem [Master Plan Senior](../../catatan/Master-Plan-Senior.md):
- **Microservices Design**: Lihat [System Design Architect Hub](../../catatan/05-Architecture-Hubs/System-Design-Architect-Hub.md) untuk pola implementasi.
- **Server-Side Rendering**: Integrasi gRPC dengan [Golang Frontend Hub](../../catatan/03-Digital-UI-Hubs/Golang-Frontend-WASM-Hub.md).

---
*Kembali ke [Documentation Hub](../docs/README.md)*
