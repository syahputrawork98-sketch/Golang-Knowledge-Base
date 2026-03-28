# RAK-03: Evolution

Rak ini membahas bagaimana ekosistem Go berkembang setelah fondasi bahasanya dipahami: module system, testing, concurrency tingkat lanjut, observability, dan internals yang relevan untuk engineer modern.

## Struktur

### [SR-01-module-system](./SR-01-module-system/)
Manajemen dependency, workspace, proxy, keamanan supply chain, dan lifecycle modul Go.

### [SR-02-testing-quality](./SR-02-testing-quality/)
Strategi testing, mocking, fuzzing, coverage, dan reliabilitas kode.

### [SR-03-concurrency](./SR-03-concurrency/)
Primitive sinkronisasi dan pola concurrency yang lebih lanjut dari fondasi dasar.

### [SR-04-observability](./SR-04-observability/)
Profiling, tracing, metrics, dan cara membaca perilaku program Go di runtime.

### [SR-05-internals](./SR-05-internals/)
Area optimisasi dan perilaku internal yang mulai menjembatani praktik engineer dengan pemahaman yang lebih dalam.

## Boundary

- fokus pada evolusi tooling, kualitas, dan pola kerja Go modern;
- cocok setelah fondasi bahasa sudah kuat;
- bukan tempat utama untuk sejarah lahirnya Go atau bedah detail runtime/compiler tingkat paling rendah.

---
*Status: [x] Complete*
