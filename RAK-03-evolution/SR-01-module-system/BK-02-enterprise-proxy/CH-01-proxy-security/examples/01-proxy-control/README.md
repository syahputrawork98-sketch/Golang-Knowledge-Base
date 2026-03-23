# Lab 01: GOPROXY Control & Inspection

## Objective
Memahami cara Go mencari modul dan bagaimana memaksa build offline atau menggunakan proxy spesifik.

## Struktur Labs
- `demo/`: Proyek kecil untuk testing.

## Workflow
1. Cek konfigurasi saat ini: `go env GOPROXY`.
2. Gunakan `GOPROXY=off` untuk melihat kegagalan build saat dependensi belum ada di cache.
3. Gunakan `GOPROXY=direct` untuk bypass proxy dan ambil langsung dari source.

## Execution
```bash
# Lihat env
go env GOPROXY GOSUMDB

# Simulasi build tanpa internet/proxy
export GOPROXY=off 
go mod tidy # Akan error jika dependensi tidak ada di cache

# Reset kembali ke default
export GOPROXY=https://proxy.golang.org,direct
```
Lab ini penting untuk debugging masalah koneksi di lingkungan CI/CD perusahaan.
