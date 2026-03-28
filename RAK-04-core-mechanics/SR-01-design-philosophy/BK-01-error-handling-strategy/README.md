# BK-01: Error Handling Strategy

Buku ini membahas cara Go memandang error sebagai bagian dari desain sistem: bukan sekadar cara menangkap kegagalan, tetapi cara menjaga alur program tetap jelas, terukur, dan mudah dirawat.

## Struktur

### [CH-01_ExplicitErrorHandling](./CH-01_ExplicitErrorHandling/)
Dasar error handling eksplisit dan alasan kenapa Go tidak memakai exception sebagai jalur utama.

### [CH-02_ErrorWrappingAndInspection](./CH-02_ErrorWrappingAndInspection/)
Wrapping, inspection, dan cara melacak error tanpa kehilangan konteks.

## Boundary

- fokus pada strategi desain error di Go, bukan sekadar pola `if err != nil`;
- membantu pembaca memahami kenapa error handling Go terasa eksplisit dan sengaja dibuat seperti itu;
- bukan tempat utama untuk panic/recover sebagai mekanisme ketahanan runtime penuh.

---
*Status: [x] Complete*
