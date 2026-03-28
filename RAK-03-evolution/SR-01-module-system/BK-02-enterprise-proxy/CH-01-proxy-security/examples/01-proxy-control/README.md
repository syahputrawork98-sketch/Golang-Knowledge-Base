# Example 01: Proxy Control

## Tujuan

Memberi checklist singkat untuk menginspeksi dan mengubah perilaku `GOPROXY` saat debugging dependency resolution.

## Isi Folder

- `README.md` - panduan command-only untuk inspeksi environment

## Cara Coba

1. Jalankan `go env GOPROXY GOSUMDB`.
2. Uji perilaku dengan mode seperti `GOPROXY=off` atau `GOPROXY=direct` sesuai shell yang dipakai.
3. Bandingkan hasil `go mod tidy` atau perintah modul lain pada tiap mode.

## Catatan

- Folder ini tidak menyertakan source code tambahan; ini memang sidecar konseptual.
- Gunakan lab ini sebagai panduan operasional, bukan sebagai demo aplikasi mandiri.
