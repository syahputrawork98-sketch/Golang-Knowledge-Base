# Lab 01: PostgreSQL Integration with Testcontainers

## Objective
Mempraktikkan cara setup infrastruktur database nyata secara otomatis untuk pengujian integrasi.

## Struktur Labs
- `postgres_test.go`: File pengujian yang memicu Docker container.

## Pre-requisites
- Docker Desktop / Engine harus berjalan.
- Library `testcontainers-go` harus terinstal (`go get github.com/testcontainers/testcontainers-go`).

## Workflow
1. Inisialisasi Postgres Container dengan image resmi.
2. Dapatkan connection string (IP & Port dinamis).
3. Buat koneksi `sql.DB`.
4. Jalankan query nyata.
5. `Terminate` container saat selesai.

## Execution
```go
// Inti dari setup:
postgresContainer, err := postgres.RunContainer(ctx,
    testcontainers.WithImage("postgres:16-alpine"),
    postgres.WithDatabase("testdb"),
    postgres.WithUsername("user"),
    postgres.WithPassword("password"),
)
```
Metode ini menjamin pengujian integrasi yang bersih dan terisolasi.
