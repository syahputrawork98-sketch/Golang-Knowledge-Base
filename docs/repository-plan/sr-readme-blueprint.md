# SR README Work Blueprint

Blueprint ini memandu fase audit dan refactor `README.md` di level `SR` sebelum kita masuk ke eksekusi vertikal pada `RAK-04`.

## 1. Tujuan Fase Ini

Tujuan utama fase `SR README` adalah:
- membuat layer navigasi tengah repo menjadi stabil;
- menyelaraskan semua `SR README` dengan docs inti yang baru;
- menyiapkan jalur turun yang rapi dari `RAK` ke `BK`;
- mengurangi kebingungan sebelum kita membenahi isi `BK`, `CH`, dan `SEC`.

Dengan kata lain, fase ini adalah jembatan antara:
- navigasi global di level root dan `RAK`;
- eksekusi vertikal per rak.

## 2. Prinsip Kerja

Fase ini mengikuti prinsip:
- sederhana dulu, lengkap kemudian;
- navigasi dulu, materi inti kemudian;
- akurat terhadap struktur folder nyata;
- jangan memaksakan keseragaman yang melanggar pengecualian struktur.

## 3. Scope

### In Scope
- seluruh file `README.md` yang hidup di level `SR`;
- pembuatan `SR README` yang memang belum ada;
- penyelarasan gaya bahasa, status, dan boundary di level `SR`.

### Out of Scope
- perubahan isi `BK`, `CH`, atau `SEC`;
- penulisan ulang chapter;
- revisi status detail di level bawah;
- perubahan struktur folder.

## 4. Kondisi Saat Ini

Hasil audit awal:

### RAK-01
- tidak memiliki `SR`;
- ini benar dan sesuai pengecualian struktur.

### RAK-02
- memiliki `7` `SR README`;
- sebagian besar sudah ada, tetapi masih memakai gaya lama dan status lama.

### RAK-03
- memiliki `5` `SR README`;
- sudah ada, tetapi gaya dan fungsinya belum sepenuhnya konsisten.

### RAK-04
- memiliki `0` `SR README`;
- ini adalah gap paling penting sebelum vertical execution.

### RAK-05
- memiliki `0` `SR README`;
- perlu dibuat agar navigasi middle layer lengkap.

### RAK-06
- memiliki `0` `SR README`;
- perlu dibuat agar jalur dari `RAK` ke `BK` lebih rapi.

## 5. Diagnosis Masalah

Masalah yang terlihat saat ini:

1. Sebagian `SR README` lama masih memakai gaya branding lama.
2. Fungsi navigasi belum konsisten.
3. Sebagian `SR README` masih terlalu tipis atau terlalu lama gayanya.
4. `RAK-04`, `RAK-05`, dan `RAK-06` belum punya `SR README`, padahal justru tiga area ini penting untuk navigasi bertingkat.

## 6. Bentuk Target SR README

Setiap `SR README` sebaiknya memiliki bentuk sederhana seperti ini:

1. Judul `SR`
2. Deskripsi singkat fungsi sub-rak
3. Daftar `BK` di dalamnya
4. Boundary singkat
5. Status

Karakter target:
- lebih mirip hub navigasi;
- lebih pendek dari `RAK README`;
- tidak berubah menjadi roadmap besar;
- tidak berubah menjadi manifesto branding.

## 7. Klasifikasi Kerja

Fase ini akan berisi tiga jenis pekerjaan:

### A. Normalize Existing SR README
Untuk `RAK-02` dan `RAK-03`:
- sederhanakan bahasa;
- hilangkan status dan branding lama;
- cocokkan dengan struktur `BK` nyata.

### B. Create Missing SR README
Untuk `RAK-04`, `RAK-05`, dan `RAK-06`:
- buat `README.md` baru di setiap folder `SR`;
- gunakan template yang seragam;
- isi hanya fungsi navigasi dan boundary.

### C. Preserve Structural Exceptions
Untuk `RAK-01`:
- tidak perlu menambah `SR`;
- jangan memaksakan pola yang tidak sesuai struktur.

## 8. Urutan Eksekusi

Urutan yang direkomendasikan:

1. Rapikan dulu `RAK-02` dan `RAK-03`
   Tujuannya: kita punya contoh `SR README` yang sudah sehat.
2. Buat `SR README` untuk `RAK-04`
   Ini paling penting karena akan jadi pintu masuk vertical execution.
3. Buat `SR README` untuk `RAK-05`
4. Buat `SR README` untuk `RAK-06`
5. Audit cepat semua hasil agar gaya bahasanya seragam.

## 9. Kenapa RAK-04 Setelah Fase Ini

`RAK-04` paling cocok jadi rak pertama untuk vertical execution karena:
- status `RAK`-nya masih `Partial`;
- posisinya ada di tengah antara pemakaian praktis dan internals rendah;
- boundary-nya penting untuk menjaga repo tetap tidak tumpang tindih;
- saat `SR README`-nya sudah ada, kita bisa turun ke `BK` dengan peta yang jauh lebih jelas.

## 10. Deliverables

Deliverables fase ini:
- seluruh `SR README` yang perlu ada sudah tersedia;
- semua `SR README` memakai gaya bahasa dan format yang lebih seragam;
- layer navigasi `RAK -> SR -> BK` menjadi jelas;
- `RAK-04` siap dipakai untuk vertical execution.

## 11. Acceptance Criteria

Fase `SR README` dianggap selesai jika:
- semua `SR` yang memang ada foldernya sudah punya `README.md`;
- tidak ada lagi `SR README` yang memakai gaya status lama;
- semua `SR README` cocok dengan struktur `BK` aktual;
- `RAK-04` punya semua `SR README` yang dibutuhkan untuk masuk fase vertikal.

## 12. Langkah Setelah Fase Ini

Setelah fase `SR README` selesai, langkah berikutnya adalah:

1. vertical execution `RAK-04`;
2. mulai dari `SR-01-design-philosophy`;
3. lanjut ke `BK` di bawahnya;
4. baru turun ke `CH` dan `SEC` sesuai prioritas.
