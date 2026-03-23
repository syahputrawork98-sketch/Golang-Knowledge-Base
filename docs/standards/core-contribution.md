# Core Contribution Standards

Dokumen ini merinci standar teknis untuk kontribusi pada materi inti (Core) di ekosistem Knowledge Base.

## 1. Akurasi Sumber (Source Alignment)
- Setiap Bab wajib mencantumkan referensi ke dokumentasi resmi (**The Go Programming Language Specification** atau **Go Blog**).
- Perubahan pada spesifikasi bahasa (misal: penambahan Generics, Iterators) harus segera direfleksikan di dalam repository ini.

## 2. Standar Contoh Kode
- Kode harus mandiri (*self-contained*) dan dapat di-compile (`go build` atau `go run`).
- Gunakan komentar untuk menjelaskan bagian yang kompleks.
- Pastikan kode mengikuti *best practices* terbaru dan idiomatik (Standard `gofmt` & `go vet`).

## 3. Standar Visualisasi
- Gunakan Mermaid.js untuk diagram alur eksekusi (Goroutine, Channels).
- Gunakan SVG untuk diagram arsitektur Runtime (Scheduler, Garbage Collector).
- Skema warna harus konsisten dengan tema Go (Gopher Blue).
