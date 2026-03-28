# SR-01: Design Philosophy

Sub-rak ini membahas alasan dasar di balik pilihan desain Go: cara bahasa ini menangani error, memandang nilai dan memori, serta mendorong komposisi dan concurrency yang khas.

## Struktur

### [BK-01-error-handling-strategy](./BK-01-error-handling-strategy/)
Strategi error handling eksplisit, wrapping, dan inspection.

### [BK-02-memory-semantics](./BK-02-memory-semantics/)
Value vs pointer, layout memori, dan implikasinya terhadap desain kode.

### [BK-03-compositional-logic](./BK-03-compositional-logic/)
Komposisi, implicit satisfaction, dan pola berpikir non-inheritance di Go.

### [BK-04-concurrency-philosophy](./BK-04-concurrency-philosophy/)
Filosofi komunikasi, ownership, dan alasan di balik model concurrency Go.

## Boundary

- fokus pada alasan desain yang membentuk karakter Go;
- membantu pembaca memahami kenapa idiom Go terlihat seperti sekarang;
- bukan tempat utama untuk tutorial sintaks dasar atau bedah source code runtime.

---
*Status: [/] Partial*
