# SR-03: Concurrency

Sub-rak ini membahas concurrency Go di level yang lebih lanjut: primitive sinkronisasi, pipeline, dan pola ketahanan sistem concurrent.

## Struktur

### [BK-01-low-level](./BK-01-low-level/)
Primitive sinkronisasi seperti `sync.Once`, `sync.Pool`, semaphore, dan `sync.Cond`.

### [BK-02-pipelines](./BK-02-pipelines/)
Fan-in, fan-out, pipeline, dan aliran kerja concurrent yang terstruktur.

### [BK-03-resilience](./BK-03-resilience/)
Rate limiting, `errgroup`, `singleflight`, dan kontrol ketahanan di sistem concurrent.

## Boundary

- fokus pada concurrency pattern yang lebih lanjut dari fondasi dasar;
- menjembatani teori dasar concurrency dengan kebutuhan sistem produksi;
- bukan tempat utama untuk filosofi concurrency Go atau scheduler runtime level rendah.

---
*Status: [x] Complete*
