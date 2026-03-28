# RAK-06 Vertical Execution Blueprint

Blueprint ini memandu vertical execution untuk `RAK-06-compiler-runtime` setelah `RAK-03`, `RAK-04`, dan `RAK-05` selesai dinormalisasi.

## Status

- Status: `Executed`
- Tanggal selesai pass utama: `2026-03-29`

## 1. Tujuan

Tujuan fase ini adalah menormalkan `RAK-06` dari atas ke bawah sambil menutup gap struktur yang masih nyata:

1. `RAK`
2. `SR`
3. `BK`
4. `CH`
5. `examples`
6. `status reconciliation`

Fokus utamanya bukan memperbanyak materi, tetapi:

- membangun `BK README` yang saat ini belum ada;
- menutup chapter yang strukturnya belum lengkap;
- menormalkan chapter ke format repo baru;
- merapikan referensi `examples/` agar hanya menunjuk file yang benar-benar ada.

## 2. Kenapa RAK-06 Dikerjakan Berikutnya

`RAK-06` layak menjadi target berikutnya karena:

- ia adalah lapisan terdalam yang secara alami datang setelah `RAK-05`;
- boundary dengan `RAK-03`, `RAK-04`, dan `RAK-05` sekarang sudah lebih jelas;
- banyak konsep penting seperti compiler pipeline, scheduler, allocator, dan cgo sudah punya struktur awal yang bisa diselamatkan;
- menyelesaikan `RAK-06` akan membuat fondasi inti repo Go ini nyaris utuh dari permukaan sampai mesin.

## 3. Kondisi Aktual RAK-06

### Struktur

`RAK-06` saat ini memiliki:

- `4` sub-rak (`SR`)
- `8` buku (`BK`)
- `9` chapter (`CH`)

Rinciannya:

### SR-01-compiler-architecture

- `BK-01-frontend`
- `BK-02-middle-end`
- `BK-03-backend`

### SR-02-runtime-orchestration

- `BK-01-scheduler`
- `BK-02-stack`
- `BK-03-garbage-collection`

### SR-03-memory-management

- `BK-01-allocator`

### SR-04-system-interfacing

- `BK-01-assembly-cgo`

## 4. Temuan Audit Penting

Temuan utama dari audit awal:

1. `RAK README` dan semua `SR README` sudah ada dan cukup sehat.
2. Semua `BK README` belum ada sama sekali.
3. Sebagian besar chapter masih memakai format legacy dan penutup status lama seperti `Platinum Gold`.
4. Ada gap struktur nyata: `CH-02_GoAssembly` ada sebagai folder, tetapi belum punya `README.md`.
5. Setidaknya `7` chapter masih menyebut example yang tidak benar-benar ada.
6. Ada chapter seperti `CH-01_SSA` yang menyebut examples, tetapi folder `examples/`-nya kosong.
7. Folder root `assets/` dan `examples/` di level `RAK-06` ada, tetapi kosong.

Implikasinya:

- `RAK-06` bukan sekadar normalisasi isi; ada **structure repair** yang nyata;
- kita harus lebih hati-hati daripada di `RAK-05`, karena beberapa klaim teknis mudah terdengar terlalu absolut;
- referensi lab perlu diperlakukan dengan sangat jujur agar pembaca tidak mendapat kesan palsu bahwa semua pembuktian sudah lengkap.

## 5. Karakter Masalah RAK-06

Secara umum, `RAK-06` adalah kombinasi tiga jenis pekerjaan:

- **structure completion**: karena `BK README` belum ada dan ada chapter yang belum lengkap;
- **content normalization**: karena chapter masih bergaya legacy;
- **technical hardening**: karena materi di rak ini lebih rentan salah framing jika ditulis terlalu simplistis atau terlalu absolut.

Artinya, strategi terbaik adalah:

- bangun dulu semua `BK README`;
- tutup gap struktur `CH` yang belum punya `README`;
- refactor chapter satu per satu dengan bahasa yang presisi;
- rapikan lab dan klaim teknis sambil jalan.

## 6. Prinsip Eksekusi

Fase ini mengikuti prinsip:

- lengkapi layer `BK` sebelum audit chapter mendalam;
- jangan membuat klaim implementasi runtime/compiler yang lebih tegas daripada bukti yang ada;
- jaga `RAK-06` tetap membahas mesin Go, bukan melompat ke framework atau tooling umum;
- perlakukan example yang tidak lengkap secara jujur;
- prioritaskan presisi teknis di atas gaya bombastis.

## 7. Urutan Eksekusi Besar

### Phase A: Build Missing BK Layer and Repair Gaps

Buat semua `BK README` yang belum ada dan tutup gap struktur yang nyata.

Target:

- `8` `BK README` baru;
- `CH-02_GoAssembly/README.md` dibuat;
- jalur `RAK -> SR -> BK -> CH` menjadi lengkap.

### Phase B: Audit SR-01 sampai CH

Mulai dari `SR-01-compiler-architecture`.

Alasan:

- paling aman untuk menetapkan gaya teknis `RAK-06`;
- alur compiler pipeline relatif linear dan boundary-nya cukup jelas;
- cocok jadi baseline sebelum masuk runtime yang lebih sensitif.

Fokus:

- parser
- AST
- SSA
- code generation

### Phase C: Audit SR-02 sampai CH

Lanjut ke `SR-02-runtime-orchestration`.

Alasan:

- ini inti runtime behavior yang paling sering dirujuk;
- perlu dinormalisasi hati-hati karena mudah tumpang tindih dengan `RAK-03` dan `RAK-04`.

Fokus:

- scheduler
- stack growth
- garbage collection

### Phase D: Audit SR-03 sampai CH

Lanjut ke `SR-03-memory-management`.

Alasan:

- allocator dan spans sangat dekat dengan runtime internals, tetapi masih cukup terlokalisasi;
- lebih aman dikerjakan setelah scheduler dan GC punya framing yang jelas.

### Phase E: Audit SR-04 sampai CH

Terakhir masuk `SR-04-system-interfacing`.

Alasan:

- cgo dan assembly paling sensitif terhadap environment dan correctness claim;
- ada gap struktur nyata di area ini;
- paling mudah meleset kalau dikerjakan terlalu cepat.

## 8. Prioritas Detail

Urutan `SR` yang direkomendasikan:

1. `SR-01-compiler-architecture`
2. `SR-02-runtime-orchestration`
3. `SR-03-memory-management`
4. `SR-04-system-interfacing`

Pertimbangannya:

- compiler architecture paling aman untuk membuka rak ini;
- runtime orchestration adalah inti yang perlu baseline narasi kuat;
- memory management mengikuti secara natural;
- system interfacing paling rawan environment-specific, jadi ditaruh terakhir.

## 9. Risiko Utama

### Risiko 1: Overclaim Teknis

Chapter bisa terdengar terlalu pasti padahal implementasi detail runtime/compiler punya nuansa dan konteks.

Mitigasi:

- tulis dengan presisi dan hindari klaim absolut tanpa kebutuhan;
- fokus pada model mental yang benar, bukan detail berlebihan yang belum dibuktikan di repo.

### Risiko 2: Example Mismatch

README menyebut file example yang tidak ada atau tidak lengkap.

Mitigasi:

- cek folder `examples/` sebelum menulis daftar lab;
- hapus referensi yang tidak punya bukti fisik.

### Risiko 3: Struktur Palsu Terlihat Lengkap

Karena `SR README` sudah ada, rak ini bisa tampak lengkap padahal layer `BK` dan sebagian `CH` belum sehat.

Mitigasi:

- audit dari `BK` ke bawah, bukan hanya percaya status di `SR`.

### Risiko 4: Environment-Specific Drift

Cgo, assembly, dan beberapa lab runtime bisa sangat tergantung environment.

Mitigasi:

- tandai mana yang demonstratif dan mana yang benar-benar runnable;
- jangan memaksa verifikasi penuh jika environment saat ini tidak mendukung.

## 10. Acceptance Criteria

Fase vertical execution `RAK-06` dianggap selesai jika:

- semua `BK README` sudah ada;
- gap struktur chapter yang hilang sudah ditutup;
- semua `CH README` utama sudah selaras dengan format repo baru;
- referensi `examples/` hanya menunjuk file yang benar-benar ada;
- boundary antar `SR` di `RAK-06` terasa jelas dan tidak saling menumpuk secara liar;
- status `RAK`, `SR`, `BK`, dan `CH` terasa konsisten secara bukti.

## 11. Unit Kerja yang Direkomendasikan

Supaya progres tetap terasa jelas, gunakan unit kerja berikut:

1. satu `BK`
2. buat atau revisi `BK README`
3. audit semua `CH` dalam `BK` tersebut
4. cek sinkronisasi `examples/`
5. tutup satu `BK`
6. pindah ke `BK` berikutnya dalam `SR` yang sama

## 12. Hasil Eksekusi

Pass utama `RAK-06` sudah dijalankan dengan hasil berikut:

1. `SR-01-compiler-architecture` selesai dinormalisasi.
   - `BK README` untuk frontend, middle-end, dan backend sudah dibuat.
   - `CH-01_ParserAST`, `CH-01_SSA`, dan `CH-01_CodeGen` sudah dipindah ke format repo baru.
   - `CH-01_SSA` sekarang memiliki example lokal nyata: `01_ssa_candidate.go`.

2. `SR-02-runtime-orchestration` selesai dinormalisasi.
   - `BK README` untuk scheduler, stack, dan garbage collection sudah dibuat.
   - `CH-01_GMPModel`, `CH-01_GrowthCopying`, dan `CH-01_TriColorALgo` sudah dipindah ke format repo baru.
   - Referensi example yang tidak ada sudah dibersihkan.
   - Tiga lab inti berhasil diverifikasi secara lokal.

3. `SR-03-memory-management` selesai dinormalisasi.
   - `BK-01-allocator/README.md` sudah dibuat.
   - `CH-01_SpansArenas` sudah dipindah ke format repo baru.
   - Referensi example yang tidak ada sudah dibersihkan.

4. `SR-04-system-interfacing` selesai dinormalisasi.
   - `BK-01-assembly-cgo/README.md` sudah dibuat.
   - `CH-01_CgoBarrier` sudah dipindah ke format repo baru.
   - Gap struktur `CH-02_GoAssembly/README.md` sudah ditutup.
   - Example baru `01_compile_to_assembly.go` ditambahkan agar chapter assembly punya lab nyata.

5. Layer `RAK -> SR -> BK -> CH` sekarang sudah lengkap di seluruh `RAK-06`.

## 13. Residual Risk

Beberapa catatan yang masih perlu dijaga:

1. verifikasi example allocator di `SR-03` masih parsial karena command timeout di environment saat ini;
2. verifikasi penuh untuk `cgo` belum berhasil karena bergantung pada kondisi toolchain C lokal;
3. verifikasi jalur inspeksi assembly juga belum mulus karena error refresh sandbox, walau example dan dokumentasinya sudah disiapkan;
4. karena topik `RAK-06` sangat sensitif, audit presisi teknis lebih dalam masih bisa dilakukan di masa lanjut tanpa mengubah kesimpulan pass utama ini.
