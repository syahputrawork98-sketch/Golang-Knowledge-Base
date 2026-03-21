# Alur Kerja Penulisan (PPM) V4 - JS Edition

Prosedur Penulisan Materi (PPM) memastikan setiap Bab (**CH**) atau Section (**SEC**) di JavaScript Knowledge Base memiliki kualitas yang setara dengan repositori "Gold Standard" lainnya.

## Tahapan PPM (Prosedur Penulisan Materi) V4

Setiap penyusunan `README.md` unit materi (CH/SEC) wajib disusun dengan **urutan representasi** 5 Tahapan (Gold Standard) berikut secara disiplin:

### 1. Tahap 1: Source Alignments & Judul
- **Header**: Judul materi yang diiringi dengan metafora/analogi singkat.
- **Source Link**: Tautan spesifik langsung ke MDN atau Spesifikasi ECMA-262 yang relevan sebagai jaminan akurasi.

### 2. Tahap 2: Konsep & Esensi (Definisi & Rasionalitas)
- **Definisi ("Apa itu?")**: Menjelaskan konsep dasar secara formal dan teknis.
- **Rasionalitas ("Why & How?")**: Membedah alasan fitur diciptakan dan masalah apa yang ia selesaikan.
- **Analogi Model Mental**: Menjabarkannya dalam contoh dunia nyata.
- **Terminologi Teknis**: Istilah senior yang digunakan oleh arsitek.

### 3. Tahap 3: Visualisasi Sistem (Mermaid)
- **Inline Diagram**: Diagram Mermaid (seperti *Call Stack*, *Event Loop*, *State Flow*) diletakkan secara **inline** (````mermaid````) tepat setelah penjelasan konsep. Tidak perlu memisahkan kode Mermaid ke `.mmd` fisik di `assets/`.
- **Fokus Visual**: Memfokuskan pada arus eksekusi atau manajemen memori berdasar konsep yang dibahas.

### 4. Tahap 4: Mekanisme Pembuktian (Algoritma Detil)
- **Detail Balik Layar**: Menguliti bagaimana kode dieksekusi oleh mesin V8 secara *Low-Level* (Bitecode, Engine behavior, Runtime logic).

### 5. Tahap 5: Multi-file Lab Praktis (Examples)
- **Kewajiban Referensi**: Setiap teori yang dikemukakan di atas diakhiri dengan referensi pembelajaran praktik.
- **Struktur Folder `examples/`**: Harus memiliki beberapa skrip `.js` berurutan dengan prefix numerik (misal: `01_dasar.js`, `02_edge_cases.js`). Pastikan setiap skrip valid, dapat dijalankan, dan memiliki komentar bedah kode yang mantap.

### Pengecualian Eksplisit "Nil Content" (Narasi Sejarah)
Jika suatu unit (terutama di RAK-01) murni bersifat historis atau konseptual non-teknis:
- **Tulis Penafian**: Wajib menyertakan teks *"Unit ini tidak membutuhkan Lab Praktis/Visualisasi karena bersifat penjelasan sejarah/konsep naratif."* pada bagian `README.md` terkait.
- **Pelarangan Direktori**: **JANGAN** membuat direktori `assets/` maupun `examples/` yang pada akhirnya akan dibiarkan kosong, demi mencegah penumpukan struktur kosong.

---
*Target Akhir: Mencapai [Gold Standard](./architecture.md#kriteria-gold-standard-100-complete).*
