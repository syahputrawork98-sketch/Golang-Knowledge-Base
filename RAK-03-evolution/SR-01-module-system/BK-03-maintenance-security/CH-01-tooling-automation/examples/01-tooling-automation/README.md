# Lab 01: go list & go mod edit Automation

## Objective
Menguasai penggunaan CLI untuk inspeksi dan manipulasi modul tanpa editor visual.

## Workflow
1. Gunakan `go list -m -json all` untuk melihat metadata dependensi.
2. Gunakan `go mod why -m [module-name]` untuk mencari tahu alasan sebuah modul ada di proyek.
3. Gunakan `go mod edit -require=[module@version]` untuk menambah/mengubah dependensi secara paksa.

## Execution
```bash
# Inspeksi dalam format JSON
go list -m -json all | jq .

# Kenapa ada uuid?
go mod why -m github.com/google/uuid

# Hapus require (hati-hati, gunakan tidy setelahnya)
go mod edit -droprequire=github.com/google/uuid
```
Penting untuk administrator sistem dan insinyur devops.
