# Alur Kerja Penulisan (PPM) V4 - Go Edition

Prosedur Penulisan Materi (PPM) memastikan setiap Bab di Golang Knowledge Base memiliki kualitas yang setara dengan repositori "Gold Standard" lainnya.

## Tahapan PPM (Prosedur Penulisan Materi) V4

Setiap pengerjaan Bab wajib mengikuti 4 Tahapan berikut secara disiplin:

### 1. Tahap 1: Definisi & Konsep ("Apa itu?")
- **Deskripsi**: Menjelaskan definisi materi atau fitur yang dibahas secara gamblang.
- **Landasan**: Memberikan pemahaman awal yang kuat sebelum masuk ke detail teknis.
- **Terminologi Teknis**: Mencantumkan dan menjelaskan istilah-istilah kunci yang digunakan oleh para ahli (*Senior Developers*). Tujuannya agar pembaca (sekalipun awam) dapat berkomunikasi dengan bahasa yang sama di lingkungan profesional.

### 2. Tahap 2: Rasionalitas ("Why & How?")
- **Kenapa**: Menjelaskan alasan fitur tersebut diciptakan dan masalah apa yang ia selesaikan.
- **Bagaimana**: Menjelaskan mekanisme kerja fitur tersebut dari sudut pandang arsitektural.
- **Analogi**: 
    - **Analogi Panjang**: Untuk membongkar mekanisme kompleks (misal: CSP Model, Garbage Collection).
    - **Analogi Pendek**: Untuk ringkasan cepat.
    - *Catatan*: Sesuaikan keberadaan analogi dengan kompleksitas materi.

### 3. Tahap 3: Implementasi (Examples)
- **Komprehensif**: Kode di folder `examples/` harus bersifat **lengkap** dan mampu menggambarkan seluruh materi yang dibahas di Tahap 1 & 2.
- **Validasi**: Pastikan skrip `.go` valid, dapat dikompilasi, dan menunjukkan efek yang diharapkan.

### 4. Tahap 4: Visualisasi (Assets)
- **Sequence**: Diagram (Mermaid/SVG) dibuat **setelah** seluruh narasi (Tahap 1 & 2) selesai.
- **Proses**: Baca kembali narasi, lalu visualisasikan alur eksekusi, struktur data, atau komunikasi antar goroutines ke dalam folder `assets/`.
- **Estetika**: Fokus pada kebersihan visual (Gopher aesthetic) dan gunakan emoji yang tepat untuk estetika dan intuisi.

---
*Target Akhir: Mencapai [Gold Standard](./architecture.md#kriteria-gold-standard-100-complete).*
