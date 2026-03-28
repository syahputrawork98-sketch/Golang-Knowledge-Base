# Example 01: Vulnerability Scan

## Tujuan

Menunjukkan bentuk minimal proyek yang bisa dipakai untuk mencoba `govulncheck` pada dependency yang diketahui punya riwayat kerentanan.

## Isi Folder

- `go.mod` - mendeklarasikan dependency demo
- `main.go` - memakai paket dari dependency tersebut

## Cara Coba

1. Instal `govulncheck` bila belum ada.
2. Jalankan `govulncheck ./...` dari folder example ini.
3. Bandingkan hasilnya dengan `go list -m all` atau upgrade versi dependency lalu scan ulang.

## Catatan

- Lab ini memerlukan dependency eksternal dan biasanya butuh akses jaringan saat setup awal.
- Tujuannya adalah belajar membaca hasil audit, bukan mengeksploitasi kerentanan.
