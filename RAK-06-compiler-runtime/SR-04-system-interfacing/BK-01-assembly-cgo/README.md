# BK-01: Assembly and Cgo

Buku ini membahas batas interaksi antara Go dengan C dan assembly, termasuk biaya transisi dan area yang perlu diperlakukan lebih hati-hati.

## Chapter

### [CH-01_CgoBarrier](./CH-01_CgoBarrier/)
Memahami transisi eksekusi dari runtime Go ke fungsi C dan implikasinya terhadap scheduler, stack, dan pointer safety.

### [CH-02_GoAssembly](./CH-02_GoAssembly/)
Memahami bagaimana membaca assembly hasil kompilasi Go untuk melihat konsekuensi instruksi mesin dari kode Go biasa.

## Boundary

- fokus pada perbatasan Go dengan lapisan sistem rendah;
- menjelaskan kapan cgo dan assembly relevan, serta biaya mentalnya;
- bukan tempat utama untuk tutorial optimisasi platform-spesifik yang lebih dalam.

---
*Status: [x] Complete*
