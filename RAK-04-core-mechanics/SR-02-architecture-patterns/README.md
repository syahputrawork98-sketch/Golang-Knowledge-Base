# SR-02: Architecture Patterns

Sub-rak ini membahas pola arsitektur yang tumbuh alami dari desain Go: middleware, concurrency pattern, dan desain interface yang robust.

## Struktur

### [BK-01-middleware-decorators](./BK-01-middleware-decorators/)
Functional options, middleware chaining, dan pola dekorasi yang lazim di Go.

### [BK-02-advanced-concurrency](./BK-02-advanced-concurrency/)
Worker pool, fan-in, fan-out, pipeline, dan context sebagai pola arsitektur concurrent.

### [BK-03-robust-interface-design](./BK-03-robust-interface-design/)
Interface composition, dependency inversion, dan desain kontrak yang tahan perubahan.

## Boundary

- fokus pada pattern yang muncul dari cara Go membentuk sistem;
- cocok untuk pembaca yang ingin merancang kode Go yang stabil dan mudah berkembang;
- bukan tempat utama untuk package standard library atau source-level runtime internals.

---
*Status: [/] Partial*
