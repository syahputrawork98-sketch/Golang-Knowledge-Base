# Example 01: Tooling Automation Commands

## Tujuan

Merangkum perintah CLI yang sering dipakai untuk menginspeksi dan mengubah metadata modul tanpa editor manual.

## Isi Folder

- `README.md` - panduan command-only

## Cara Coba

1. Jalankan `go list -m -json all` untuk melihat metadata dependency.
2. Gunakan `go mod why -m <module>` untuk mengetahui alasan sebuah modul masuk ke graph.
3. Gunakan `go mod edit` untuk menambah atau menghapus entri tertentu secara eksplisit.

## Catatan

- Folder ini tidak menyertakan proyek demo sendiri.
- Perintah di sini paling berguna saat dijalankan pada modul nyata yang sedang Anda audit.
