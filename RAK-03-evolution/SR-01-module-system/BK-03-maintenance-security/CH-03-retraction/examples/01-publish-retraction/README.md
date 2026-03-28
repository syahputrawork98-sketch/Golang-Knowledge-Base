# Example 01: Publish Retraction

## Tujuan

Memberi contoh ringkas bagaimana maintainer menandai versi modul yang sebaiknya tidak lagi dipakai.

## Isi Folder

- `README.md` - catatan prosedural untuk simulasi retraction

## Cara Coba

1. Siapkan modul yang sudah memiliki riwayat versi.
2. Tambahkan direktif `retract` di `go.mod`.
3. Rilis versi perbaikan berikutnya.
4. Gunakan `go list -m -u all` untuk melihat sinyal retraction dari sudut pandang consumer.

## Catatan

- Folder ini bersifat konseptual; tidak ada repo tag sungguhan di dalamnya.
- Cocok dipakai sebagai checklist maintainer, bukan sebagai demo runnable penuh.
