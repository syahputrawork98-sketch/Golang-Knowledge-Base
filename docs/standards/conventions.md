# Konvensi Penamaan & Struktur Folder (JS Edition)

Untuk menjaga keteraturan "Perpustakaan Dunia Maya", setiap file dan folder wajib mengikuti konvensi penamaan berikut sesuai dengan **Unified Gold Standard**.

## 1. Penamaan Direktori

| Level | Prefix | Format | Contoh |
| :--- | :--- | :--- | :--- |
| Rak | `RAK-` | `RAK-XX-slug` | `RAK-01-introduction-essence` |
| Sub-Rak | `SR-` | `SR-XX-slug` | `SR-01-get-started` |
| Buku | `BK-` | `BK-XX_Slug` | `BK-01_Basics` |
| Bab | `CH-` | `CH-XX_Slug` | `CH-01_JSFirstSteps` |
| Section | `SEC-` | `SEC-XX_Slug` | `SEC-01_Introduction` |

> [!IMPORTANT]
> **Rak dan Sub-Rak** menggunakan tanda hubung (`-`) untuk slug. 
> **Buku, Bab, dan Section** menggunakan garis bawah (`_`) untuk memisahkan urutan dengan nama slug.

## 2. Struktur Internal Unit (Level 5 & 6)

Setiap folder Bab (**CH**) atau Section (**SEC**) memiliki struktur dasar berikut:
```text
CH- atau SEC-/
├── README.md        <- Materi teks inti berserta diagram Mermaid inline (Target Utama PPM).
├── examples/        <- Kode lab fungsional multi-file (.js). (Opsional jika "Nil Content")
└── assets/          <- Arsip media statis eksternal (.png/.svg). (Opsional)
```

> [!CAUTION]
> **Aturan "Nil Content" (Murni Narasi)**:
> Jika sebuah unit secara spesifik bersifat sejarah/filosofis (Nil Content), Anda **DILARANG BERSERTA-SERTA** membuat direktori `examples/` maupun `assets/`. Biarkan folder unit tersebut hanya berisi `README.md` demi menjaga kebersihan repositori.

## 3. Aturan Penamaan Lab Praktis (`examples/`)
File di dalam direktori `examples/` **WAJIB** menggunakan *prefix* numerik berurutan untuk menjaga alur pembelajaran lab.
**Contoh Benar**: `01_basic.js`, `02_advanced_closures.js`, `03_edge_cases.js`.
**Contoh Salah**: `dasar.js`, `contoh-closure.js`, `test.js`.

## 4. Aturan README.md
Setiap tingkatan dari Level 2 hingga Level 6 wajib memiliki file `README.md` sebagai hub navigasi atau penyaji materi.

---
*Referensi: [Architecture](./architecture.md) | [Unified Gold Standard](../../../brain/931398eb-3011-4b69-bb8c-e94cd60f9e78/unified_gold_standard_v1.md)*
