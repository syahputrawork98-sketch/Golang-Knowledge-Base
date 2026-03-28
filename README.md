# Golang Knowledge Base

`Golang Knowledge Base` adalah monorepository materi belajar Go yang disusun dengan struktur perpustakaan bertingkat: dari anatomi bahasa, fondasi penggunaan, evolusi ekosistem, mekanika inti, standard library, hingga compiler dan runtime.

Repo ini dipakai sebagai:
- knowledge base terstruktur untuk pembelajaran Go;
- basis dokumentasi yang sinkron dengan sumber resmi;
- workspace penulisan materi teknis berbasis `PPM V4`.

## Struktur Utama

### `RAK-01-anatomy`
Narasi asal-usul, filosofi, dan konteks ekosistem Go.

### `RAK-02-foundation`
Fondasi bahasa Go dari sudut pandang engineer yang menulis kode sehari-hari.

### `RAK-03-evolution`
Evolusi tooling, module system, testing, concurrency patterns, dan observability.

### `RAK-04-core-mechanics`
Rasionalitas desain dan pola sistemik Go.

### `RAK-05-standard-library`
Eksplorasi paket bawaan dan utilitas utama di standard library.

### `RAK-06-compiler-runtime`
Bedah compiler, scheduler, memory management, dan batas Go dengan sistem luar.

## Cara Membaca Repo Ini

Jika kamu baru masuk ke repo:
1. baca [docs/standards/README.md](./docs/standards/README.md);
2. lanjut ke [docs/standards/authoring.md](./docs/standards/authoring.md);
3. lihat [docs/repository-plan/README.md](./docs/repository-plan/README.md);
4. baru navigasi ke `RAK-*` yang relevan.

## Dokumen Penting

- [docs/standards/README.md](./docs/standards/README.md): panduan struktur repo dan aturan inti.
- [docs/standards/authoring.md](./docs/standards/authoring.md): panduan penulisan materi dan Gold Standard.
- [docs/repository-plan/README.md](./docs/repository-plan/README.md): blueprint besar repositori.
- [status.md](./status.md): status operasional repo.

## Catatan Penting

- Aturan formal repo diringkas di `docs/standards/README.md` dan `docs/standards/authoring.md`.
- `.cursorrules` adalah ringkasan operasional AI, bukan sumber standar utama.
- `README.md` ini sengaja dibuat ringkas agar menjadi landing page, bukan spesifikasi penuh.
