# BK-03: Maintenance and Security

Buku ini membahas alat dan praktik untuk memelihara kesehatan dependency graph, mengaudit kerentanan, dan menangani kesalahan publikasi versi modul secara lebih aman.

## Struktur

### [CH-01-tooling-automation](./CH-01-tooling-automation/)
Inspeksi graph modul dan otomasi perubahan metadata dengan toolchain bawaan.

### [CH-02-security-audit](./CH-02-security-audit/)
Audit kerentanan modul dengan `govulncheck` dan pendekatan yang lebih minim noise.

### [CH-03-retraction](./CH-03-retraction/)
Penarikan kembali versi modul yang salah rilis melalui mekanisme retraction.

## Boundary

- fokus pada maintenance dan keamanan supply chain modul;
- membantu pembaca mengelola dependency graph secara lebih aman dan terukur;
- bukan tempat utama untuk proxy privat atau lifecycle dasar modul.

---
*Status: [x] Complete*
