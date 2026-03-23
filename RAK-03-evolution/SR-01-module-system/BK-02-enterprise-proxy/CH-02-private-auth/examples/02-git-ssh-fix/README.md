# Lab 02: SSH Instead of HTTPS for Private Modules

## Objective
Mensimulasikan penggunaan SSH Key untuk modul privat tanpa menyentuh `.netrc`.

## Workflow
1. Gunakan konfigurasi global Git untuk memaksa SSH:
   ```bash
   git config --global url."git@github.com:".insteadOf "https://github.com/"
   ```
2. Setel `GOPRIVATE=github.com/company-org/*`.
3. Jalankan `go get`.

## Result
Setiap kali Go mencoba melakukan `clone` via HTTPS, Git akan secara otomatis mengubahnya menjadi SSH (`git@github.com:...`), sehingga menggunakan SSH Key lokal Anda.
