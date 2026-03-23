# Lab 01: v2 Module Path Implementation

## Objective
Mempraktikkan cara mendefinisikan dan mengonsumsi modul versi major `v2` secara idiomatik.

## Struktur Labs
- `lib/`: Modul yang sudah naik kelas ke `v2`.
- `app/`: Aplikasi yang mengonsumsi `lib/v2`.

## Workflow
1. Inisialisasi library dengan module path `example.com/lib/v2`.
2. Gunakan `import "example.com/lib/v2"` di aplikasi.
3. Jalankan aplikasi.

## Execution
```bash
cd lib
go mod init example.com/lib/v2
cd ../app
go mod init example.com/app
# Tambahkan require di go.mod app
go build
```
Perhatikan bahwa meskipun lokasinya di folder `lib/`, identitas modulnya sudah mengandung suffix `/v2`.
