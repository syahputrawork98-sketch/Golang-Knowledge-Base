# Lab 01: Corporate GOPRIVATE & .netrc Setup

## Objective
Mensimulasikan konfigurasi akses ke modul privat menggunakan token personal.

## Workflow (Conceptual)
1. Setel domain privat: `go env -w GOPRIVATE=github.company.com`.
2. Buat file `.netrc` (Linux/macOS) atau `_netrc` (Windows) di Home directory.
3. Tambahkan entry:
   ```text
   machine github.company.com
   login [YOUR_GITHUB_TOKEN]
   password x-oauth-basic
   ```
4. Jalankan `go get github.company.com/org/repo`.

## Result
Go akan mendeteksi `GOPRIVATE`, melewati Proxy/SumDB, dan menggunakan kredensial di `.netrc` untuk melakukan autentikasi HTTPS ke Git.
