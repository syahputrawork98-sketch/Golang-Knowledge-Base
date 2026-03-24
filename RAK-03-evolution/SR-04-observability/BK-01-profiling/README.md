# BK-01: Profiling & Deep Inspection

Buku ini membahas teknik inspeksi mendalam terhadap performa aplikasi Go menggunakan `pprof`. Fokusnya adalah menemukan bottleneck CPU, pemborosan memori (Heap), kebocoran Goroutine, hingga hambatan sinkronisasi (Mutex Contention).

## Chapters

| Chapter | Topik | Konsep Inti |
|---------|-------|------------|
| [CH-01](./CH-01-pprof/README.md) | CPU & Heap Analysis | Flat/Cum Time, Flamegraphs, Allocation sampling |
| [CH-02](./CH-02-goroutines/README.md) | Goroutine & Thread | Leak detection, Stack grouping, Thread profiling |
| [CH-03](./CH-03-contention/README.md) | Block & Mutex Contention | Lock wait identification, SetMutexProfileFraction |

## Key Visual Assets

| Asset | Description |
|-------|-------------|
| `CH-01/assets/pprof-sampling.svg` | CPU/Heap sampling and pprof data pipeline |
| `CH-02/assets/goroutine-leak.svg` | Leak detection via stack trace grouping |
| `CH-03/assets/mutex-contention.svg` | Mutex wait queue and accumulated wait visualization |

---
*Back to [SR-04 Page](../README.md)*
