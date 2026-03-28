# BK-02: Performance Benchmarking

Buku ini membahas alat ukur performa di Go: benchmark, profiling, dan fuzz testing sebagai cara membuktikan kualitas teknis dengan data, bukan tebakan.

## Struktur

### [CH-01_MicroBenchmarking](./CH-01_MicroBenchmarking/)
Microbenchmark untuk mengukur throughput, latency, dan alokasi pada unit kode kecil.

### [CH-02_ProfilingStrategies](./CH-02_ProfilingStrategies/)
Profiling dan tracing untuk menemukan bottleneck nyata di program Go.

### [CH-03_FuzzTesting](./CH-03_FuzzTesting/)
Fuzz testing untuk menemukan input aneh, crash, dan edge case yang sering lolos dari test biasa.

## Boundary

- fokus pada pengukuran, pembuktian, dan workflow evaluasi performa;
- membantu pembaca memilih alat ukur yang tepat untuk masalah yang sedang dihadapi;
- bukan tempat utama untuk observability produksi penuh atau runtime internals tingkat rendah.

---
*Status: [x] Complete*
