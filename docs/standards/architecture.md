# Arsitektur & Hierarki Struktur (Go Edition)

Proyek **Golang Knowledge Base** disusun dengan satu standar tunggal: **Unified Gold Standard**. Sistem ini memastikan keseimbangan antara narasi yang manusiawi dengan ketajaman teknis spesifikasi.

## 1. Hirarki 6-Level (Universe Standard)

Sistem ini mengikuti hirarki kedalaman yang konsisten untuk seluruh repositori:

| Tingkatan | Nama | Analogi | Prefix | Contoh Direktori |
| :--- | :--- | :--- | :--- | :--- |
| **Level 1** | **Root** | Library Hub | `/` | `/` |
| **Level 2** | **Rak** | Domain Utama | `RAK-` | `RAK-01-anatomy/` |
| **Level 3** | **Sub-Rak** | Track Spesifik | `SR-` | `SR-01-history-origins/` |
| **Level 4** | **Buku** | Koleksi Terpadu | `BK-` | `BK-01_Basics/` |
| **Level 5** | **Bab** | Materi Inti | `CH-` | `CH-01_GoFirstSteps/` |
| **Level 6** | **Section** | Detail Halaman | `SEC-` | `SEC-01_Introduction/` |
 
 > [!IMPORTANT]
 > **Pengecualian RAK-01 (Anatomy)**: 
 > Karena sifatnya yang murni naratif dan pengantar, RAK-01 menggunakan **Flat Structure** dengan melompati Level 3 (**SR-**). Struktur RAK-01 adalah: **RAK > BK > CH > SEC**.

---

## 2. Prinsip "Digital Mirroring"
Hierarki ini mencerminkan struktur sumber primer (The Go Programming Language Specification / Go Tour) secara **1:1**. Jika sumber menuntut kedalaman lebih (misal: sub-clause), sistem akan menggunakan Level 6 (**SEC**) untuk menampung detail tersebut tanpa merusak struktur utama.

---

## 3. Karakteristik & Branding
- **Analogi Utama**: **Pusat Skalabilitas Cloud (The Cloud-Native Scale Hub)**.
- **Tone Suara**: **Efisien, Kuat, dan Pragmatis**.
- **Filosofi**: Fokus pada konkurensi, keamanan memori, dan kesederhanaan desain Go.

---

## 4. Kriteria Kelulusan Standar (Gold Standard)
Sebuah unit materi (CH/SEC) dianggap **Complete** secara berurutan dan terstruktur jika menyajikan 5 representasi ini di dalam `README.md`:
1. **Source Link** (Akurasi Spec Go / Go Blog).
2. **Dual Definition** (Definisi Formal + Analogi Model Mental).
3. **Mermaid Diagram inline** (Visualisasi aliran/status yang dirender langsung di markdown, bukan direferensikan dari luar).
4. **Mekanisme Detil** (Algoritma / Runtime Behind-the-Scene).
5. **Lab Praktis** (Penjelasan & relasi ke kode `.go` di direktori `examples/`).

---
*Referensi: [Unified Gold Standard](../../../brain/931398eb-3011-4b69-bb8c-e94cd60f9e78/unified_gold_standard_v1.md)*
