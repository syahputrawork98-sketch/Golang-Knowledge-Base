# Lab 01: Minimal Version Selection (MVS) Simulation

## Objective
Membuktikan bahwa Go memilih versi dependensi yang paling 'tua' namun memenuhi syarat minimum (MVS), bukan versi terbaru (Latest).

## Struktur Labs
- `app/`: Aplikasi utama yang membutuhkan `pkg-a`.
- `pkg-a-v1.1.0/`: Mock modul versi 1.1.0.
- `pkg-a-v1.2.0/`: Mock modul versi 1.2.0.

## Workflow
1. Inisialisasi dependensi dengan `go mod tidy`.
2. Lihat bagaimana Go memilih versi 1.2.0 jika diminta, namun tetap menjaga stabilitas.
3. Gunakan `go list -m all` untuk melihat hasil final resolusi.

## Execution
Jalankan di terminal:
```bash
cd app
go mod tidy
go list -m all
```
Anda akan melihat versi yang dipilih adalah versi terendah yang masih memenuhi semua batasan (`require`).
