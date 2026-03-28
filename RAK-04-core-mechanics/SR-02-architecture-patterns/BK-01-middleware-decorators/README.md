# BK-01: Middleware and Decorators

Buku ini membahas pola middleware dan decorator yang tumbuh alami dalam ekosistem Go: bagaimana perilaku bisa disisipkan, dirangkai, dan divalidasi tanpa membuat API menjadi rumit.

## Struktur

### [CH-01_FunctionalOptions](./CH-01_FunctionalOptions/)
Pola functional options untuk membangun API yang fleksibel dan tetap mudah dirawat.

### [CH-02_MiddlewareChaining](./CH-02_MiddlewareChaining/)
Pola merangkai middleware agar perilaku lintas fungsi bisa ditambahkan secara modular.

## Boundary

- fokus pada pola penyusunan perilaku dan konfigurasi di level arsitektur;
- membantu pembaca melihat kenapa Go nyaman dengan fungsi kecil yang bisa dirangkai;
- bukan tempat utama untuk tutorial `net/http` dasar atau framework-specific middleware.

---
*Status: [x] Complete*
