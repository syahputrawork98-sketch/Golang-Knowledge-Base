# Lab 01: Version Retraction Simulation

## Objective
Mempraktikkan cara menarik kembali versi modul yang memiliki bug kritis.

## Workflow
1. Siapkan modul dengan tag `v1.0.0`.
2. Temukan "bug" (simulasi).
3. Tambahkan direktif `retract` di `go.mod`.
4. Publikasikan versi `v1.0.1`.
5. Pengguna yang mencoba mengambil `v1.0.0` akan mendapat peringatan.

## Execution
Contoh penulisan di `go.mod`:
```go
module example.com/my-module

go 1.22

// v1.0.0 ditarik karena mengandung celah keamanan pada parser.
retract v1.0.0
```
Gunakan `go list -m -u all` untuk melihat apakah ada versi yang pernah Anda gunakan telah ditarik oleh maintainer.
