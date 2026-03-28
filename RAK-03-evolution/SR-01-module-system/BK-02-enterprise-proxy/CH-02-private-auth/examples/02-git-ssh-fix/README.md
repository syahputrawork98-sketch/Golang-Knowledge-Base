# Example 02: Git SSH Rewrite

## Tujuan

Menunjukkan pendekatan umum untuk memaksa Git memakai SSH saat Go mencoba mengambil modul privat.

## Isi Folder

- `README.md` - panduan konfigurasi konseptual

## Cara Coba

1. Atur rewrite URL Git dari HTTPS ke SSH dengan `git config`.
2. Setel `GOPRIVATE` sesuai organisasi atau domain privat.
3. Pastikan SSH key lokal memang sudah valid untuk host target.
4. Jalankan ulang operasi modul yang sebelumnya gagal lewat HTTPS.

## Catatan

- Folder ini sengaja tidak menyertakan source code tambahan.
- Lab ini cocok dipakai saat masalah utamanya ada di transport Git, bukan di `go.mod`.
