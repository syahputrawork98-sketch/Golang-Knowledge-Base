# Example 01: Corporate GOPRIVATE Setup

## Tujuan

Memberi panduan dasar untuk konfigurasi `GOPRIVATE` dan kredensial HTTPS saat mengakses modul privat di lingkungan perusahaan.

## Isi Folder

- `README.md` - panduan konfigurasi konseptual

## Cara Coba

1. Setel domain privat dengan `go env -w GOPRIVATE=<domain-privat>`.
2. Siapkan `.netrc` atau `_netrc` sesuai sistem operasi.
3. Isi kredensial yang dibutuhkan oleh host Git privat.
4. Coba jalankan perintah seperti `go get` atau `go mod tidy` pada modul privat yang relevan.

## Catatan

- Folder ini tidak membawa token, host nyata, atau modul privat sungguhan.
- Perlakukan lab ini sebagai checklist setup, bukan contoh runnable penuh.
