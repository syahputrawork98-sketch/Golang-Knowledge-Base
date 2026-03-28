# RAK-03 Vertical Execution Blueprint

Blueprint ini memandu vertical execution untuk `RAK-03-evolution` setelah governance layer dan `RAK-04` selesai direkonsiliasi.

## 1. Tujuan

Tujuan fase ini adalah menormalkan `RAK-03` dari atas ke bawah tanpa membongkar struktur yang sebenarnya sudah lengkap:

1. `RAK`
2. `SR`
3. `BK`
4. `CH`
5. `example-sidecars`
6. `status reconciliation`

Fokus utamanya bukan membangun navigasi dari nol, tetapi menyelaraskan gaya, status, boundary, dan referensi lab ke standar repo yang sekarang.

## 2. Kenapa RAK-03 Dikerjakan Berikutnya

`RAK-03` layak menjadi target berikutnya karena:
- topiknya langsung menyambung dengan area yang baru kita kerjakan di `RAK-04`;
- ia punya banyak materi bernilai tinggi: modules, testing, concurrency, observability, dan internals;
- strukturnya sudah penuh, jadi bottleneck utamanya bukan folder kosong, melainkan inkonsistensi isi;
- merapikan `RAK-03` akan membantu menjaga boundary dengan `RAK-04` dan `RAK-06`.

## 3. Kondisi Aktual RAK-03

### Struktur

`RAK-03` saat ini memiliki:
- `5` sub-rak (`SR`)
- `15` buku (`BK`)
- `44` chapter (`CH`)

Rinciannya:

### SR-01-module-system
- `BK-01-module-lifecycle`
- `BK-02-enterprise-proxy`
- `BK-03-maintenance-security`

### SR-02-testing-quality
- `BK-01-unit-testing`
- `BK-02-mocking-patterns`
- `BK-03-reliability`

### SR-03-concurrency
- `BK-01-low-level`
- `BK-02-pipelines`
- `BK-03-resilience`

### SR-04-observability
- `BK-01-profiling`
- `BK-02-tracing`
- `BK-03-telemetry`

### SR-05-internals
- `BK-01-memory`
- `BK-02-scheduler`
- `BK-03-optimization`

## 4. Temuan Audit Penting

Temuan utama dari audit awal:

1. Semua `SR` dan semua `BK` sudah punya `README.md`.
2. Semua `CH` utama juga sudah punya `README.md`.
3. Berbeda dengan `RAK-04`, masalah utama `RAK-03` bukan kekurangan layer navigasi, tetapi warisan gaya lama.
4. Setidaknya `59` README di `RAK-03` masih kuat memakai pola legacy seperti:
   - `*Target: ...*`
   - `Back to ...`
   - heading seperti `The Logic`, `The Lab`, `Senior Terms`
5. Ada `12` README tambahan di dalam folder `examples/`, jadi `RAK-03` punya sidecar docs lab yang lebih banyak daripada `RAK-04`.
6. Status level `RAK` dan `SR` sudah tertulis `Complete`, tetapi itu lebih mencerminkan kelengkapan struktur daripada hasil normalisasi penuh ke standar baru.

Implikasinya:
- vertical execution `RAK-03` harus diperlakukan sebagai **large-scale normalization**, bukan **structure build**;
- kita tidak perlu membuat banyak `BK README`, tetapi perlu audit isi secara disiplin;
- README di dalam `examples/` perlu diperlakukan sebagai scope sekunder, bukan target utama pada pass pertama.

## 5. Karakter Masalah RAK-03

Secara umum, `RAK-03` berada di tengah antara dua kondisi:

- lebih rapi daripada `RAK-04` saat belum dieksekusi, karena struktur dan README utamanya sudah lengkap;
- tetapi lebih berisiko dari `RAK-04`, karena materi lamanya jauh lebih banyak dan overlap boundary antartopik lebih sensitif.

Artinya, strategi terbaik bukan membuat ulang semuanya, tetapi:
- bekukan struktur;
- audit per `SR`;
- refactor per `BK`;
- rapikan contoh dan sidecar lab setelah inti `CH` stabil.

## 6. Prinsip Eksekusi

Fase ini mengikuti prinsip:
- pertahankan struktur yang sudah sehat;
- refactor isi utama dulu, sidecar belakangan;
- selesaikan satu `SR` penuh sebelum pindah ke `SR` berikutnya;
- jaga boundary dengan ketat agar `RAK-03` tidak bocor ke `RAK-04` atau `RAK-06`;
- perlakukan status `Complete` lama sebagai provisional sampai unit selesai diaudit.

## 7. Urutan Eksekusi Besar

### Phase A: Freeze Structure dan Audit Baseline

Tidak ada pembangunan struktur besar-besaran.

Target:
- petakan semua `SR/BK/CH`;
- tandai pola legacy yang berulang;
- pisahkan unit utama dari README sidecar di `examples/`.

### Phase B: Audit SR-01 sampai CH

Kerjakan `SR-01-module-system` dulu.

Alasan:
- paling dekat dengan konteks kerja terakhir;
- contoh lab modul di repo relatif matang;
- boundary-nya cukup jelas dan risikonya rendah dibanding concurrency atau internals.

Fokus:
- `go.mod`, `go.sum`, `go.work`, semver, proxy, private auth, security audit, retraction.

### Phase C: Audit SR-02 sampai CH

Lanjut ke `SR-02-testing-quality`.

Alasan:
- masih cukup aman secara boundary;
- kontribusi ke kualitas repo sangat besar;
- pola testing lebih mudah dinormalisasi setelah module-system rapi.

### Phase D: Audit SR-04 sampai CH

Lanjut ke `SR-04-observability`.

Alasan:
- observability punya relasi kuat ke workflow engineering;
- lebih aman dikerjakan sebelum concurrency dan internals agar tidak kebablasan ke runtime forensics.

### Phase E: Audit SR-03 sampai CH

Baru masuk `SR-03-concurrency`.

Alasan:
- ini area paling rawan overlap dengan `RAK-04`;
- banyak pola di sini harus dibedakan dari "design philosophy" dan "architecture patterns" di rak lain.

### Phase F: Audit SR-05 sampai CH

Terakhir kerjakan `SR-05-internals`.

Alasan:
- area ini paling rawan overlap dengan `RAK-06`;
- butuh boundary yang sangat disiplin agar tidak berubah menjadi bedah runtime/compiler mendalam.

### Phase G: Reconcile Status dan Sidecars

Setelah seluruh `SR` utama selesai:
- selaraskan status `BK`, `SR`, dan `RAK`;
- audit README di `examples/` yang masih penting dipertahankan;
- hapus atau sederhanakan sidecar docs yang ternyata tidak perlu.

## 8. Prioritas Detail

Urutan `SR` yang direkomendasikan:

1. `SR-01-module-system`
2. `SR-02-testing-quality`
3. `SR-04-observability`
4. `SR-03-concurrency`
5. `SR-05-internals`

Urutan ini sengaja tidak mengikuti nomor folder secara kaku.

Pertimbangannya:
- modules dan testing paling aman untuk menetapkan pola normalisasi;
- observability lebih aman sebelum concurrency dan internals;
- concurrency dan internals ditaruh belakangan karena risiko overlap lebih tinggi.

## 9. Risiko Utama

### Risiko 1: Overlap dengan RAK-04

Topik seperti fuzzing, profiling, pipeline, atau context mudah bertabrakan dengan pola di `RAK-04`.

Mitigasi:
- `RAK-03` harus tetap fokus pada evolusi workflow dan praktik engineer modern;
- hindari framing yang terlalu menekankan filosofi desain jika yang dibutuhkan adalah ekosistem dan workflow.

### Risiko 2: Overlap dengan RAK-06

`SR-05-internals` bisa terlalu cepat berubah menjadi pembahasan runtime atau compiler level rendah.

Mitigasi:
- cukup sampai level pemahaman engineer dan implikasi performa;
- dorong source-level internals yang lebih dalam ke `RAK-06`.

### Risiko 3: Legacy Drift

Karena struktur `RAK-03` sudah lengkap, ada godaan untuk menganggap semuanya sudah baik-baik saja.

Mitigasi:
- treat all current `Complete` statuses as provisional until audited;
- audit per `BK`, bukan hanya membaca status di `README`.

### Risiko 4: Sidecar README Sprawl

README kecil di dalam `examples/` bisa membuat scope meledak jika disentuh terlalu awal.

Mitigasi:
- pass pertama fokus pada `RAK -> SR -> BK -> CH`;
- sidecar lab README hanya disentuh jika broken, misleading, atau benar-benar mengganggu pembacaan utama.

## 10. Acceptance Criteria

Fase vertical execution `RAK-03` dianggap selesai jika:
- semua `SR`, `BK`, dan `CH` utama sudah selaras dengan gaya repo baru;
- referensi examples dan assets di README utama hanya menunjuk file yang benar-benar ada;
- status tidak lagi mengandalkan klaim lama, tetapi sesuai bukti aktual;
- boundary `RAK-03` terhadap `RAK-04` dan `RAK-06` terasa jelas;
- sidecar docs di `examples/` sudah diputuskan: dipertahankan, disederhanakan, atau dihapus secara sadar.

## 11. Unit Kerja yang Direkomendasikan

Supaya progres tetap terasa jelas, gunakan unit kerja berikut:

1. satu `BK`
2. audit semua `CH` dalam `BK` tersebut
3. revisi `BK README`
4. tentukan apakah sidecar `examples/README.md` perlu disentuh
5. tutup satu `BK`
6. pindah ke `BK` berikutnya dalam `SR` yang sama

Jangan berpindah antar `SR` sebelum satu `SR` benar-benar stabil.

## 12. Langkah Pertama yang Disarankan

Langkah implementasi pertama yang paling sehat adalah:

1. audit `SR-01-module-system` secara penuh;
2. mulai dari `BK-01-module-lifecycle`;
3. refactor `CH-01-mod-anatomy`, `CH-02-go-work`, dan `CH-03-semver`;
4. baru lanjut ke `BK-02-enterprise-proxy` dan `BK-03-maintenance-security`.
