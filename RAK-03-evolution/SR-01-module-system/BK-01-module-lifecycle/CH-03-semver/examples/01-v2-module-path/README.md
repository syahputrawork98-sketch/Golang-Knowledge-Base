# Example 01: v2 Module Path

## Tujuan

Menunjukkan bentuk idiomatik modul major version `v2`, baik di sisi producer maupun consumer.

## Isi Folder

- `lib/` - modul dengan path `example.com/lib/v2`
- `app/` - consumer yang mengimpor `example.com/lib/v2`

## Cara Coba

1. Lihat `lib/go.mod` untuk memastikan module path memakai suffix `/v2`.
2. Lihat `app/go.mod` untuk melihat `require` dan `replace` yang menunjuk ke modul lokal.
3. Masuk ke folder `app/`.
4. Jalankan `go run .`.

## Catatan

- Lab ini runnable karena consumer sudah memakai `replace example.com/lib/v2 => ../lib`.
- Poin utamanya adalah identitas modul, bukan nama folder semata.
