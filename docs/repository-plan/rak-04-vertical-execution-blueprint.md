# RAK-04 Vertical Execution Blueprint

Blueprint ini memandu eksekusi vertikal untuk `RAK-04-core-mechanics` setelah layer `RAK` dan `SR` selesai dirapikan.

## 1. Tujuan

Tujuan fase ini adalah membenahi `RAK-04` dari atas ke bawah secara terstruktur:

1. `RAK`
2. `SR`
3. `BK`
4. `CH`
5. `status reconciliation`

Artinya, kita tidak lagi berpindah-pindah antar rak. Kita fokus penuh pada satu rak sampai bentuknya stabil.

## 2. Kenapa RAK-04 Dipilih Dulu

`RAK-04` paling layak dikerjakan lebih dulu karena:
- statusnya masih `Partial`;
- posisinya strategis di antara `RAK-02` dan `RAK-06`;
- boundary-nya penting untuk mencegah tumpang tindih materi;
- layer `SR`-nya baru saja selesai dibangun dan siap dijadikan peta turun.

## 3. Kondisi Aktual RAK-04

### Struktur

`RAK-04` saat ini memiliki:
- `3` sub-rak (`SR`)
- `9` buku (`BK`)
- `21` chapter (`CH`)

Rinciannya:

### SR-01-design-philosophy
- `BK-01-error-handling-strategy`
- `BK-02-memory-semantics`
- `BK-03-compositional-logic`
- `BK-04-concurrency-philosophy`

### SR-02-architecture-patterns
- `BK-01-middleware-decorators`
- `BK-02-advanced-concurrency`
- `BK-03-robust-interface-design`

### SR-03-tooling-excellence
- `BK-01-dependency-management`
- `BK-02-performance-benchmarking`

## 4. Temuan Audit Penting

Temuan paling penting sebelum eksekusi:

1. `RAK README` sudah dirapikan.
2. `SR README` sudah tersedia dan cukup stabil.
3. Semua `BK` di `RAK-04` belum memiliki `README.md`.
4. `CH README` sudah ada, tetapi masih kuat memakai gaya lama.

Implikasinya:
- eksekusi vertikal tidak boleh langsung dimulai dari polishing chapter;
- kita perlu membangun layer `BK README` dulu supaya jalur `SR -> BK -> CH` lengkap.

## 5. Prinsip Eksekusi

Fase ini mengikuti prinsip:
- bangun kerangka dulu, baru audit isi;
- selesaikan satu cabang penuh sebelum pindah terlalu jauh;
- utamakan stabilitas boundary dan navigasi;
- jangan ubah banyak chapter sekaligus tanpa pola yang jelas.

## 6. Urutan Eksekusi Besar

### Phase A: Build BK README Layer

Buat `README.md` untuk semua `BK` di `RAK-04`.

Target:
- seluruh `BK` punya fungsi sebagai hub navigasi;
- semua `BK README` menampilkan daftar `CH` aktual;
- boundary `BK` menjadi jelas.

Ini adalah fondasi minimum sebelum audit chapter.

### Phase B: Audit SR-01 sampai CH

Kerjakan `SR-01-design-philosophy` secara vertikal penuh:
- buat dan rapikan semua `BK README`;
- audit seluruh `CH` di bawahnya;
- selaraskan gaya, struktur, dan status.

Alasan memilih `SR-01` dulu:
- paling dekat dengan identitas desain `RAK-04`;
- membantu menetapkan tone dan boundary rak ini;
- materinya relatif paling konseptual dan cocok jadi acuan.

### Phase C: Audit SR-02 sampai CH

Lanjut ke `SR-02-architecture-patterns` setelah `SR-01` stabil.

Fokus:
- pattern middleware;
- pattern concurrency;
- desain interface yang robust.

Ini area yang rawan tumpang tindih dengan `RAK-02` dan `RAK-03`, jadi boundary harus dijaga ketat.

### Phase D: Audit SR-03 sampai CH

Lanjut ke `SR-03-tooling-excellence`.

Fokus:
- dependency management dari sudut pandang `core mechanics`;
- benchmarking, profiling strategies, dan fuzzing.

Ini area paling rawan overlap dengan `RAK-03`, jadi perlu kehati-hatian ekstra.

### Phase E: Reconcile Status

Setelah seluruh cabang `RAK-04` selesai:
- revisi status di level `BK`;
- revisi status di level `SR`;
- revisi status di level `RAK`;
- sinkronkan ke `status.md` root.

## 7. Unit Kerja yang Direkomendasikan

Supaya progres terasa jelas, gunakan unit kerja berikut:

1. satu `BK README`
2. audit semua `CH` dalam satu `BK`
3. tutup satu `BK`
4. pindah ke `BK` berikutnya

Jangan lompat antar `BK` terlalu cepat.

## 8. Prioritas Detail

Prioritas eksekusi yang direkomendasikan:

1. `SR-01-design-philosophy`
2. `SR-02-architecture-patterns`
3. `SR-03-tooling-excellence`

Di dalam `SR-01`, urutan yang saya sarankan:
1. `BK-01-error-handling-strategy`
2. `BK-02-memory-semantics`
3. `BK-03-compositional-logic`
4. `BK-04-concurrency-philosophy`

Alasan:
- error handling dan memory semantics paling mendefinisikan karakter `RAK-04`;
- compositional logic dan concurrency philosophy lebih aman dikerjakan setelah dua fondasi itu jelas.

## 9. Risiko Utama

### Risiko 1: Overlap dengan RAK-02

Beberapa topik seperti pointer, defer, atau context bisa mudah kembali ke mode tutorial.

Mitigasi:
- selalu jaga POV `architectural rationale`;
- jangan menjelaskan sekadar cara pakai jika yang dibutuhkan adalah alasan desain.

### Risiko 2: Overlap dengan RAK-06

Saat membahas mekanisme internal, penjelasan bisa terlalu jauh masuk ke source-level runtime.

Mitigasi:
- cukup sampai level rasionalitas dan implikasi desain;
- dorong bedah runtime penuh ke `RAK-06`.

### Risiko 3: Chapter Drift

Karena chapter sudah ada lebih dulu, gaya lama bisa menarik kita untuk patch acak tanpa pola.

Mitigasi:
- bangun `BK README` dulu;
- audit per `BK`, bukan per file acak.

## 10. Acceptance Criteria

Fase vertical execution `RAK-04` dianggap selesai jika:
- semua `BK` di `RAK-04` punya `README.md`;
- seluruh layer `RAK -> SR -> BK -> CH` terbaca jelas;
- gaya bahasa dan status lebih konsisten;
- overlap utama dengan `RAK-02` dan `RAK-06` berhasil ditekan;
- `RAK-04` siap dipertimbangkan naik status setelah audit selesai.

## 11. Langkah Pertama yang Disarankan

Langkah implementasi pertama yang paling sehat adalah:

1. buat `README.md` untuk semua `BK` di `SR-01-design-philosophy`;
2. audit `CH-01_ExplicitErrorHandling`;
3. audit `CH-02_ErrorWrappingAndInspection`;
4. tutup `BK-01-error-handling-strategy` sebagai cabang pertama.
