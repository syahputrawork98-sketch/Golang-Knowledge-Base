# Protokol Pembaruan Status (Go Edition)

Progress pengerjaan proyek Golang Knowledge Base dihitung secara otomatis (bubbling up) dari unit terkecil.

## 1. Level Bab (Chapter)
Status dicatat di `docs/status.md` pada level **Buku**.
- `Draft`: Narasi sedang ditulis.
- `Partial`: Kode atau Diagram belum lengkap.
- `Sync`: Mencapai Gold Standard.

## 2. Level Buku
Setiap Buku wajib memiliki `docs/status.md` yang merangkum progress seluruh Bab di dalamnya.
Format: `(Σ Bab Sync) / (Total Bab)`.

## 3. Level Rak & Global
Progress dari level Buku dilaporkan ke:
1. `RAK-XX/README.md`
2. `docs/README.md` (Global Status Table)

---
*Peringatan: Dokumentasi dianggap "Published" hanya jika statusnya `Sync`.*
