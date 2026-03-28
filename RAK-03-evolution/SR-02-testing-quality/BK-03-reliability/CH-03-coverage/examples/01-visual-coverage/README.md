# Example 01: Visual Coverage

## Tujuan

Memberi langkah singkat untuk membaca hasil coverage secara visual setelah test dijalankan.

## Isi Folder

- `README.md` - panduan command-only untuk analisis coverage

## Cara Coba

1. Jalankan `go test -coverprofile=coverage.out ./...` pada module target.
2. Jalankan `go tool cover -html=coverage.out`.
3. Baca bagian hijau sebagai jalur yang teruji dan bagian merah sebagai jalur yang belum tersentuh test.

## Catatan

- Folder ini memang tidak menyertakan source code demo tambahan.
- Gunakan panduan ini bersama project atau example lain yang memang punya file test nyata.
