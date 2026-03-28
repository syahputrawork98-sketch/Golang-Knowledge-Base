# Repository Guide

Panduan ini adalah pintu masuk utama untuk memahami cara kerja `Golang Knowledge Base`. Tujuannya adalah merangkum aturan repositori dalam satu dokumen yang ringkas dan mudah dibaca.

## 1. Cara Baca Dokumen

Urutan baca yang direkomendasikan:
1. file ini, `docs/standards/README.md`
2. [authoring.md](./authoring.md)
3. [../repository-plan/README.md](../repository-plan/README.md)

Jika ada konflik aturan, gunakan urutan ini:
1. `docs/standards/README.md`
2. `docs/standards/authoring.md`
3. `docs/repository-plan/README.md`
4. `.cursorrules`
5. `README.md`
6. `status.md`

## 2. Struktur Repo

Repositori ini memakai hierarki:

1. `Root`
2. `RAK`
3. `SR`
4. `BK`
5. `CH`
6. `SEC`

Pengecualian:
- `RAK-01` boleh melewati `SR` karena sifatnya naratif dan pengantar.

## 3. Fungsi Tiap Level

| Level | Prefix | Fungsi |
| :--- | :--- | :--- |
| Root | `/` | Hub utama repositori |
| RAK | `RAK-` | Domain besar pembelajaran |
| SR | `SR-` | Track spesifik |
| BK | `BK-` | Kumpulan chapter |
| CH | `CH-` | Unit materi utama |
| SEC | `SEC-` | Rincian terdalam bila dibutuhkan |

## 4. Naming Convention

Gunakan format berikut:

| Level | Format | Contoh |
| :--- | :--- | :--- |
| RAK | `RAK-XX-slug` | `RAK-01-anatomy` |
| SR | `SR-XX-slug` | `SR-01-basics` |
| BK | `BK-XX_Slug` | `BK-01_Foundations` |
| CH | `CH-XX_Slug` | `CH-01_Variables` |
| SEC | `SEC-XX_Slug` | `SEC-01_Introduction` |

Aturan tambahan:
- `RAK` dan `SR` memakai tanda hubung `-`;
- `BK`, `CH`, dan `SEC` memakai garis bawah `_` setelah nomor.

## 5. Struktur Minimum Folder

### Untuk `RAK`, `SR`, dan `BK`
- wajib punya `README.md` sebagai navigasi atau ringkasan.

### Untuk `CH` dan `SEC`
- wajib punya `README.md`;
- boleh punya `examples/`;
- boleh punya `assets/`.

Jika unit bersifat naratif murni:
- cukup `README.md`;
- jangan buat `examples/` kosong;
- jangan buat `assets/` kosong.

## 6. Terminologi yang Dipakai

Gunakan istilah Go yang resmi dan presisi, misalnya:
- `goroutine`
- `channel`
- `escape analysis`
- `scheduler`
- `allocator`

Gunakan istilah status operasional berikut saja:
- `Draft`
- `Partial`
- `Complete`

Hindari memakai istilah status bercampur seperti `Sync`, `Premium`, atau `Platinum`.

## 7. Aturan Status

Status dasar dicatat di level `CH` atau `SEC`:
- `[ ] Draft`
- `[/] Partial`
- `[x] Complete`

Status lalu naik ke `BK`, `SR`, `RAK`, lalu ke root `status.md`.

Prinsipnya sederhana:
- status tidak boleh lebih optimistis dari bukti yang tersedia;
- jika audit belum selesai, pakai status yang konservatif.

## 8. Boundary Antar Rak

Untuk menghindari tumpang tindih:

| Rak | POV |
| :--- | :--- |
| `RAK-02` | Engineer intuition |
| `RAK-04` | Architectural rationale |
| `RAK-06` | Low-level implementation |

## 9. Dokumen yang Benar-Benar Penting

Setelah perampingan, dokumen inti repo ini adalah:
- `README.md`: landing page manusia
- `docs/standards/README.md`: panduan struktur dan aturan repo
- `docs/standards/authoring.md`: panduan penulisan materi
- `docs/repository-plan/README.md`: blueprint besar isi repo
- `.cursorrules`: ringkasan operasional untuk AI
- `status.md`: status operasional repo
