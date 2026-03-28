# Example 01: Postgres Integration with Testcontainers

## Tujuan

Menunjukkan pola dasar pengujian integrasi yang menyalakan PostgreSQL sementara lewat Testcontainers.

## Isi Folder

- `postgres_test.go` - contoh test integrasi yang menyiapkan container Postgres

## Cara Coba

1. Pastikan Docker berjalan.
2. Pastikan dependency Go untuk testcontainers dan driver database yang dibutuhkan sudah tersedia di module context Anda.
3. Jalankan `go test` dari folder yang sesuai.
4. Amati lifecycle container dan koneksi database selama test berlangsung.

## Catatan

- Sidecar ini menjelaskan snippet integrasi, bukan setup module lengkap yang sepenuhnya mandiri.
- Jika dependency belum tersedia, Anda perlu menambahkannya terlebih dahulu sebelum test dapat berjalan penuh.
