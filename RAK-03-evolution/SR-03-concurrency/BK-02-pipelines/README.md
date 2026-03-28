# BK-02: Pipelines and Flow Patterns

Buku ini membahas pola aliran kerja concurrent di Go: distribusi kerja, pipeline bertahap, dan generator lazy yang menjaga aliran data tetap efisien.

## Struktur

### [CH-01-fan-out-in](./CH-01-fan-out-in/)
Fan-out dan fan-in untuk membagi beban kerja ke banyak worker lalu menggabungkan hasilnya kembali.

### [CH-02-orchestration](./CH-02-orchestration/)
Pipeline bertahap dengan `context.Context` untuk cancellation, teardown, dan koordinasi lintas stage.

### [CH-03-generators](./CH-03-generators/)
Generator berbasis goroutine dan channel untuk aliran data lazy dan memory-efficient.

## Boundary

- fokus pada pola aliran kerja concurrent yang dipakai engineer aplikasi;
- membantu pembaca merancang stage, cancellation, dan konsolidasi hasil dengan rapi;
- bukan tempat utama untuk membahas scheduler runtime atau observability pipeline produksi.

---
*Status: [x] Complete*
