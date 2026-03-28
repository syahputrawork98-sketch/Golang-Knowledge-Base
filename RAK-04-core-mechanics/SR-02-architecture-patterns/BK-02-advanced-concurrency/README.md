# BK-02: Advanced Concurrency

Buku ini membahas pola concurrency yang dipakai untuk membentuk arsitektur sistem Go: bagaimana pekerjaan didistribusikan, digabungkan, dialirkan, dan dikoordinasikan secara lebih terstruktur.

## Struktur

### [CH-01_WorkerPools](./CH-01_WorkerPools/)
Worker pool untuk membatasi dan mengelola beban kerja concurrent.

### [CH-02_FanInOut](./CH-02_FanInOut/)
Fan-in dan fan-out untuk penggabungan serta penyebaran aliran kerja.

### [CH-03_PipelinePattern](./CH-03_PipelinePattern/)
Pipeline sebagai model pemrosesan bertahap yang jelas aliran datanya.

### [CH-04_ContextPattern](./CH-04_ContextPattern/)
Context sebagai alat koordinasi pembatalan, timeout, dan propagation lintas goroutine.

## Boundary

- fokus pada pola concurrency sebagai bentuk arsitektur, bukan sekadar mekanik channel;
- membantu pembaca memilih bentuk alur kerja concurrent yang tepat untuk sistem nyata;
- bukan tempat utama untuk scheduler runtime, race detector, atau concurrency basic.

---
*Status: [x] Complete*
