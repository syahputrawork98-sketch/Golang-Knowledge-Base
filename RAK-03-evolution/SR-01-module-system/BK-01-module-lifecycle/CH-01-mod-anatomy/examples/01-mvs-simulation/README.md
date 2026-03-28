# Example 01: MVS Simulation

## Tujuan

Menunjukkan cara Go menyelesaikan versi modul saat beberapa dependency meminta versi minimum yang berbeda.

## Isi Folder

- `app/` - modul aplikasi utama
- `lib-a-v11/` - mock modul `lib-a` versi lebih lama
- `lib-a-v12/` - mock modul `lib-a` versi lebih baru
- `lib-b/` - modul lain yang ikut memengaruhi resolusi versi

## Cara Coba

1. Masuk ke folder `app/`.
2. Jalankan `go mod tidy`.
3. Jalankan `go list -m all`.
4. Amati versi modul yang akhirnya dipilih oleh resolver.

## Catatan

- Lab ini bersifat runnable dan cocok dipakai untuk membaca hasil resolusi modul secara langsung.
- Fokusnya bukan pada fetch dari internet, tetapi pada hubungan versi antar modul lokal.
