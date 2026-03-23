# Lab 01: Multi-Module Workspace Sync

## Objective
Mempraktikkan penggunaan `go.work` untuk mengelola aplikasi dan library lokal secara simultan tanpa `replace`.

## Struktur Labs
- `app/`: Modul aplikasi utama.
- `lib/`: Modul library pendukung.
- `go.work`: File kontrol workspace.

## Workflow
1. Inisialisasi workspace: `go work init ./app ./lib`.
2. Edit `lib/lib.go`.
3. Jalankan `app/main.go` dan lihat perubahannya langsung teraplikasi tanpa update versi di `go.mod`.

## Execution
```bash
go work init ./app ./lib
cd app
go run main.go
```
Go akan mendeteksi `example.com/lib` ada di direktori `./lib` melalui konfigurasi `go.work`.
