# RAK-05 Vertical Execution Blueprint

Blueprint ini memandu vertical execution untuk `RAK-05-standard-library` setelah `RAK-03` dan `RAK-04` selesai dinormalisasi.

## Status

- Status blueprint: `Executed`
- Tanggal penyelesaian pass utama: `2026-03-29`
- Cakupan selesai: `SR -> BK -> CH` utama
- Catatan examples: `referensi example utama sudah diselaraskan; tidak ada sidecar README di RAK-05`

## 1. Tujuan

Tujuan fase ini adalah menormalkan `RAK-05` dari atas ke bawah sambil melengkapi layer navigasi yang belum ada:

1. `RAK`
2. `SR`
3. `BK`
4. `CH`
5. `examples`
6. `status reconciliation`

Fokus utamanya bukan membangun ulang struktur besar, tetapi:

- membuat `BK README` yang saat ini belum ada;
- memindahkan chapter ke format repo yang sekarang;
- menyelaraskan referensi `examples/` dengan file yang benar-benar ada.

## 2. Kenapa RAK-05 Dikerjakan Berikutnya

`RAK-05` layak menjadi target berikutnya karena:

- ia masih dekat secara boundary dengan `RAK-03` dan `RAK-04`, jadi transisinya alami;
- topiknya sangat sentral untuk pembaca Go sehari-hari;
- strukturnya kecil dan terkontrol, sehingga cocok untuk vertical execution yang cepat dan bersih;
- ia lebih aman dikerjakan sebelum `RAK-06`, yang jauh lebih berat dan sensitif boundary-nya.

## 3. Kondisi Aktual RAK-05

### Struktur

`RAK-05` saat ini memiliki:

- `4` sub-rak (`SR`)
- `7` buku (`BK`)
- `12` chapter (`CH`)

Rinciannya:

### SR-01-data-processing

- `BK-01-string-text`
- `BK-02-serialization`

### SR-02-io-filesystem

- `BK-01-io-abstractions`
- `BK-02-file-management`

### SR-03-networking

- `BK-01-web-api`

### SR-04-utilities

- `BK-01-security`
- `BK-02-time`

## 4. Temuan Audit Penting

Temuan utama dari audit awal:

1. `RAK README` dan semua `SR README` sudah ada dan cukup sehat.
2. Semua `CH README` utama juga sudah ada.
3. Semua `BK README` belum ada sama sekali.
4. Chapter masih kuat memakai format legacy dan penutup status lama seperti `Platinum Gold`.
5. Tidak ada sidecar `README` di dalam `examples/`, jadi scope contoh lebih kecil daripada `RAK-03`.
6. Setidaknya `8` chapter menyebut file example yang tidak benar-benar ada di folder `examples/`.

Implikasinya:

- `RAK-05` adalah kombinasi **structure completion** dan **content normalization**;
- kita harus membangun layer `BK` dulu sebelum menutup tiap cabang dengan rapi;
- akurasi referensi lab harus menjadi fokus sejak awal, bukan ditunda di akhir.

## 5. Karakter Masalah RAK-05

Secara umum, `RAK-05` berada di tengah antara dua pola:

- seperti `RAK-04` pada tahap awal, karena layer `BK` belum ada;
- seperti `RAK-03`, karena chapter-chapternya mostly butuh normalisasi isi, bukan rekonstruksi total.

Artinya, strategi terbaik adalah:

- bangun semua `BK README` terlebih dahulu;
- refactor `CH` per `BK`;
- rapikan mismatch `examples/` sambil jalan;
- tutup satu `SR` penuh sebelum pindah ke `SR` berikutnya.

## 6. Prinsip Eksekusi

Fase ini mengikuti prinsip:

- lengkapi layer `BK` sebelum audit chapter mendalam;
- jaga `RAK-05` tetap praktis dan engineer-facing;
- refactor isi utama sambil membersihkan referensi lab yang misleading;
- jangan melebar ke framework eksternal atau internals runtime;
- perlakukan status `Complete` saat ini sebagai provisional sampai unit diaudit.

## 7. Urutan Eksekusi Besar

### Phase A: Build Missing BK Layer

Buat semua `BK README` yang belum ada.

Target:

- `7` `BK README` baru;
- boundary tiap `BK` jelas;
- jalur `RAK -> SR -> BK -> CH` menjadi lengkap.

### Phase B: Audit SR-01 sampai CH

Mulai dari `SR-01-data-processing`.

Alasan:

- paling aman untuk menetapkan baseline gaya;
- chapter-nya cukup banyak, tetapi boundary-nya jelas;
- mismatch example di area ini mudah dipetakan dan dibersihkan.

Fokus:

- `strings`, `strconv`, `bytes`, `encoding/json`, `encoding/binary`

### Phase C: Audit SR-02 sampai CH

Lanjut ke `SR-02-io-filesystem`.

Alasan:

- masih aman secara boundary;
- penting sebagai fondasi standard library;
- beberapa referensi example di area ini jelas tidak sinkron.

Fokus:

- `io`, `bufio`, `os`, `embed`

### Phase D: Audit SR-04 sampai CH

Lanjut ke `SR-04-utilities`.

Alasan:

- ukurannya kecil;
- topik `crypto` dan `time` cukup sentral, tapi masih aman dikerjakan sebelum networking.

### Phase E: Audit SR-03 sampai CH

Terakhir masuk `SR-03-networking`.

Alasan:

- paling mudah melebar ke internet, API eksternal, dan perilaku runtime jaringan;
- salah satu lab client saat ini masih menunjuk ke endpoint eksternal sehingga perlu framing yang hati-hati.

## 8. Prioritas Detail

Urutan `SR` yang direkomendasikan:

1. `SR-01-data-processing`
2. `SR-02-io-filesystem`
3. `SR-04-utilities`
4. `SR-03-networking`

Pertimbangannya:

- `SR-01` paling aman untuk membentuk pola kerja;
- `SR-02` sangat dekat secara konsep dan manfaat;
- `SR-04` kecil dan cepat ditutup;
- `SR-03` ditaruh terakhir karena networking paling mudah melebar.

## 9. Risiko Utama

### Risiko 1: Example Drift

README menyebut file example yang tidak ada.

Mitigasi:

- cek isi folder `examples/` sebelum menulis ulang daftar lab;
- hapus referensi fiktif, jangan menebak file yang seharusnya ada.

### Risiko 2: Boundary Drift ke Area Non-Standard Library

Topik seperti HTTP, crypto, atau time mudah melebar ke framework, devops, atau internals.

Mitigasi:

- `RAK-05` harus tetap fokus pada package bawaan Go;
- hindari framing yang terlalu berat ke design pattern atau runtime internals.

### Risiko 3: Lab Verification Bias

Beberapa example tampak sederhana, tetapi bisa bergantung pada jaringan atau environment tertentu.

Mitigasi:

- tandai dengan jujur mana yang runnable lokal dan mana yang bersifat demonstratif;
- jangan memaksakan verifikasi penuh jika butuh dependency atau endpoint eksternal.

## 10. Acceptance Criteria

Fase vertical execution `RAK-05` dianggap selesai jika:

- semua `BK README` sudah ada;
- semua `CH README` utama sudah selaras dengan format repo baru;
- referensi `examples/` hanya menunjuk file yang benar-benar ada;
- status `RAK`, `SR`, `BK`, dan `CH` terasa konsisten secara bukti;
- boundary `RAK-05` sebagai standard library tetap jelas dan praktis.

## 11. Unit Kerja yang Direkomendasikan

Supaya progres tetap terasa jelas, gunakan unit kerja berikut:

1. satu `BK`
2. buat atau revisi `BK README`
3. audit semua `CH` dalam `BK` tersebut
4. cek sinkronisasi `examples/`
5. tutup satu `BK`
6. pindah ke `BK` berikutnya dalam `SR` yang sama

## 12. Langkah Pertama yang Disarankan

Langkah implementasi pertama yang paling sehat adalah:

1. buat `BK README` untuk semua buku di `SR-01-data-processing`;
2. audit `BK-01-string-text`;
3. refactor `CH-01_Strings`, `CH-02_Strconv`, dan `CH-03_Bytes`;
4. rapikan mismatch example di area itu;
5. baru lanjut ke `BK-02-serialization`.

## 13. Hasil Eksekusi

Vertical execution `RAK-05` sudah dijalankan sesuai urutan besar yang ditetapkan:

1. `SR-01-data-processing`
2. `SR-02-io-filesystem`
3. `SR-04-utilities`
4. `SR-03-networking`

Hasil pass utama:

- seluruh `BK README` yang sebelumnya belum ada kini sudah dibuat;
- semua chapter utama `RAK-05` sudah dinormalisasi ke format repo baru;
- referensi `examples/` di chapter utama sudah diselaraskan dengan file yang benar-benar ada;
- mismatch example yang sebelumnya berjumlah `8` chapter sudah dibersihkan di area utama;
- satu example HTTP client diubah agar memakai server lokal `httptest` sehingga tidak tergantung endpoint eksternal;
- satu example `io` yang keliru juga sudah dirapikan agar lebih sehat secara statik.

Residual risk yang masih tersisa:

- verifikasi runtime penuh di environment ini tidak selalu berhasil karena beberapa `go run` lokal mengalami timeout;
- beberapa example tetap bersifat demonstratif ringan, meskipun referensinya sekarang sudah jujur dan sinkron.

Kesimpulan:

- blueprint ini sudah berstatus selesai untuk **pass utama `RAK -> SR -> BK -> CH`**;
- jika nanti ada fase lanjutan, fokusnya bukan lagi normalisasi inti, tetapi hardening examples dan verifikasi runtime yang lebih menyeluruh.
