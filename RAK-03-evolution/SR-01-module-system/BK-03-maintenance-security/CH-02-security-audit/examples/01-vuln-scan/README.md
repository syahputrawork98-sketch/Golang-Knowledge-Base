# Lab 01: Vulnerability Scan Simulation

## Objective
Mempraktikkan cara mendeteksi kerentanan pada modul menggunakan `govulncheck`.

## Struktur Labs
- `demo/`: Proyek kecil yang sengaja menggunakan modul lama dengan kerentanan.

## Workflow
1. Instal `govulncheck`: `go install golang.org/x/vuln/cmd/govulncheck@latest`.
2. Jalankan pemindaian: `govulncheck ./...`.
3. Analisis laporan: cari apakah kerentanan tersebut "Reachable" (terjangkau oleh kode).

## Execution
```bash
# Instal tool
go install golang.org/x/vuln/cmd/govulncheck@latest

# Jalankan audit
govulncheck ./...
```
Anda akan melihat daftar CVE yang relevan dan anjuran upgrade versi.
