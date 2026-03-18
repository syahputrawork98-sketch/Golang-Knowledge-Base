# Golang Knowledge Base: Documentation Hub

Pusat dokumentasi ini menyimpan seluruh cetak biru, standar kualitas, dan panduan meta untuk membangun repository ini secara konsisten mengikuti standar *go.dev/doc*.

## Daftar Dokumen Utama

### 1. Standar & Protokol
- **[Architecture Standards](./standards/architecture.md)**: Definisi analogi "The Gopher Factory" dan aturan pewajiban file per tingkatan.
- **[Naming Conventions](./standards/conventions.md)**: Aturan penamaan RAK, SR, BK, hingga Level Bab.
- **[PPM Workflow](./standards/workflow.md)**: 4 Tahapan Penulisan Materi (PPM) V4.
- **[Status Protocol](./standards/status-protocol.md)**: Cara melaporkan dan menghitung progress pengerjaan.
- **[Contribution Guide](./standards/contribution.md)**: Panduan kontribusi untuk Gophers.
- **[Core Contribution](./standards/core-contribution.md)**: Standar teknis materi inti Go.

### 2. Cetak Biru (Blueprints)
- **[Repository Plan](./repository-plan/README.md)**: Dekomposisi total dari dokumentasi resmi Go ke dalam struktur 9-Rack.

### 3. Narasi & Esensi
- **[Golang Origins](./golang-origins.md)**: Kisah di balik layar mengapa Google menciptakan bahasa Go.
- **[Golang History](./golang-history.md)**: Sejarah kelahiran Go dari rasa frustrasi di Google hingga era cloud-native.
- **[Philosophy & Essence](./golang-philosophy.md)**: Membedah filosofi "Less is More" dan model CSP.
- **[Why Golang?](./why-golang.md)**: Rasionalitas teknis kapan harus memilih Go untuk sistem skala besar.

## Struktur Direktori `docs/`

```text
/docs
├── standards/             # Standar & Protokol (Architecture, Conventions, etc.)
├── repository-plan/       # Cetak biru 9-Rack (Level 2)
├── golang-history.md      # Narasi Sejarah
├── golang-philosophy.md   # Filosofi Bahasa
├── golang-origins.md      # Legenda Inisiasi
├── why-golang.md          # Rasionalitas Penggunaan
└── README.md              # File ini (Hub navigasi)
```
