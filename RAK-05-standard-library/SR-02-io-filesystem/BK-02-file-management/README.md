# BK-02: File Management and Embedded Assets

Buku ini membahas interaksi praktis dengan filesystem dan asset lokal melalui package bawaan Go, mulai dari operasi file dasar sampai penyertaan file statis ke dalam binary.

## Struktur

### [CH-01_OS](./CH-01_OS/)
Operasi file dasar, metadata, dan interaksi umum dengan environment lewat package `os`.

### [CH-02_Embed](./CH-02_Embed/)
`embed` untuk menyertakan file statis ke dalam binary agar distribusi aplikasi lebih sederhana.

## Boundary

- fokus pada file handling dan asset lokal yang dibawa standard library;
- membantu pembaca memahami hubungan antara code, filesystem, dan binary final;
- bukan tempat utama untuk networking, observability, atau konfigurasi pihak ketiga.

---
*Status: [x] Complete*
