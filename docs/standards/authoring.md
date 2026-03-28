# Authoring Guide

Panduan ini menjelaskan bagaimana materi ditulis, kapan sebuah unit dianggap selesai, dan standar minimum untuk narasi, contoh kode, dan visual.

## 1. PPM V4

Setiap unit `CH` atau `SEC` normal mengikuti 5 tahap:

### Tahap 1. Source Alignment dan Judul
- judul unit;
- tautan ke sumber resmi;
- framing atau analogi singkat bila membantu.

### Tahap 2. Konsep dan Rasionalitas
- definisi formal;
- alasan fitur itu penting;
- analogi model mental;
- terminologi teknis yang tepat.

### Tahap 3. Visualisasi Sistem
- Mermaid inline di dalam `README.md`.

### Tahap 4. Mekanisme Pembuktian
- penjelasan under-the-hood;
- perilaku runtime, compiler, memory, atau algoritma yang relevan.

### Tahap 5. Lab Praktis
- rujukan ke `examples/`;
- contoh yang runnable atau testable;
- file bernomor urut.

## 2. Gold Standard

Unit `CH` atau `SEC` hanya boleh ditandai `Complete` jika:
- kelima tahap PPM V4 terpenuhi; atau
- unit tersebut sah sebagai `Nil Content`.

`assets/` boleh dipakai sebagai pelengkap, tetapi tidak menggantikan kebutuhan Mermaid inline.

## 3. Nil Content

Jika unit murni historis, filosofis, atau naratif:
- tulis penafian eksplisit di `README.md`;
- jangan buat `examples/` kosong;
- jangan buat `assets/` kosong.

Contoh penafian:

`Unit ini tidak membutuhkan Lab Praktis atau visualisasi tambahan karena bersifat penjelasan sejarah atau konsep naratif.`

## 4. Standar Narasi

- akurat terhadap sumber resmi;
- teknis, tetapi tetap bisa dipahami;
- analogi membantu menjelaskan, bukan mengganti definisi;
- jangan membuat istilah baru jika istilah resmi Go sudah ada.

## 5. Standar Kode

- contoh kode harus self-contained sejauh memungkinkan;
- contoh harus bisa dijalankan atau diuji;
- gunakan `gofmt` dan idiom Go;
- komentar dipakai seperlunya saja.

File di `examples/` wajib bernama seperti:
- `01_basic.go`
- `02_context_timeout.go`
- `03_edge_case_test.go`

## 6. Standar Visual

- Mermaid inline adalah default;
- SVG atau PNG hanya dipakai jika visual eksternal memang dibutuhkan;
- gaya visual harus tetap sederhana, jelas, dan konsisten.

## 7. Kontribusi dan Refactor

Saat menulis atau merefaktor materi:
1. tentukan dulu lokasi unit dalam hirarki repo;
2. baca `docs/standards/README.md`;
3. ikuti PPM V4;
4. jangan meninggalkan folder kosong;
5. perbarui status hanya jika ada bukti yang cukup.

## 8. Template

Template yang tersedia:
- [templates/CHAPTER_README.md](./templates/CHAPTER_README.md)
- [templates/BOOK_STATUS.md](./templates/BOOK_STATUS.md)
- [templates/RAK_README.md](./templates/RAK_README.md)
