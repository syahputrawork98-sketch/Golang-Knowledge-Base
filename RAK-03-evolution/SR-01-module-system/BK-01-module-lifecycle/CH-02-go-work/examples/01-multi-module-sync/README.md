# Example 01: Multi-Module Workspace Sync

## Tujuan

Menunjukkan bagaimana `go.work` menyatukan beberapa modul lokal tanpa perlu `replace` di `go.mod` aplikasi.

## Isi Folder

- `app/` - modul aplikasi
- `lib/` - modul library lokal
- `go.work` - konfigurasi workspace

## Cara Coba

1. Pastikan posisi terminal berada di folder example ini.
2. Jalankan `go work sync` bila perlu untuk menyelaraskan workspace.
3. Masuk ke folder `app/`.
4. Jalankan `go run main.go`.
5. Ubah isi `lib/` lalu jalankan lagi aplikasi untuk melihat perubahan lokal langsung terbaca.

## Catatan

- Lab ini runnable, tetapi perilaku persisnya bergantung pada keadaan workspace lokal dan akses dependency yang dibutuhkan.
- Fokus utamanya adalah hubungan `app`, `lib`, dan `go.work`.
