# Example 02: Checksum Failure

## Tujuan

Menunjukkan bahwa integritas dependency di Go dijaga lewat checksum, dan perubahan manual pada hash akan memicu kegagalan verifikasi.

## Isi Folder

- `go.mod` - titik awal modul demo

## Cara Coba

1. Jalankan `go mod tidy` agar `go.sum` terbentuk.
2. Ubah salah satu hash di `go.sum` secara manual.
3. Jalankan `go build`.
4. Amati error `checksum mismatch` atau error verifikasi serupa.

## Catatan

- `go.sum` memang belum ada di awal folder ini; file itu dibuat saat dependency diselesaikan pertama kali.
- Lab ini sengaja kecil karena fokusnya ada pada perilaku verifikasi, bukan pada aplikasi contoh.
