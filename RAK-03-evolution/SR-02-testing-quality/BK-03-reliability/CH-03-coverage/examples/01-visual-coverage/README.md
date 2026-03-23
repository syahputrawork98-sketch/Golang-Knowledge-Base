# Lab 01: Visual Coverage Analysis

## Objective
Mempraktikkan cara melihat jalur kode yang belum teruji secara visual.

## Workflow
1. Jalankan test dengan profile: `go test -coverprofile=coverage.out ./...`.
2. Buka laporan visual: `go tool cover -html=coverage.out`.
3. Amati baris berwarna **merah** (belum teruji) vs **hijau** (sudah teruji).

## Execution
```bash
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```
Gunakan informasi ini untuk menambah test case pada jalur kode yang masih berwarna merah.
