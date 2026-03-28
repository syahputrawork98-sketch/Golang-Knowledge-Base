# BK-02: Enterprise Proxy

Buku ini membahas bagaimana Go mengelola distribusi modul lewat proxy publik, checksum database, dan workflow akses modul privat di lingkungan korporat.

## Struktur

### [CH-01-proxy-security](./CH-01-proxy-security/)
Peran `GOPROXY`, `GOSUMDB`, dan integritas distribusi modul.

### [CH-02-private-auth](./CH-02-private-auth/)
Konfigurasi akses modul privat dengan `GOPRIVATE`, autentikasi Git, dan workflow korporat.

## Boundary

- fokus pada distribusi modul dan autentikasi dependency di lingkungan nyata;
- membantu pembaca memahami perbedaan alur modul publik dan modul privat;
- bukan tempat utama untuk lifecycle versi modul atau audit keamanan lanjutan.

---
*Status: [x] Complete*
