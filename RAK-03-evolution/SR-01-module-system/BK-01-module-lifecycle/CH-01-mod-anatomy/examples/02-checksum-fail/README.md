# Lab 02: Checksum Verification Failure Simulation

## Objective
Membuktikan perlindungan Go terhadap modifikasi paket dengan sengaja merusak file `go.sum`.

## Workflow
1. Inisialisasi proyek dengan modul valid.
2. Edit `go.sum` secara manual (ubah hash).
3. Jalankan `go build`.
4. Amati pesan error `checksum mismatch`.

## Execution
Jalankan di terminal:
```bash
go mod tidy
# Edit go.sum, ubah satu karakter di hash
go build
```
Go akan menolak melakukan build karena integritas paket tidak lagi terjamin.
