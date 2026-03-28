# BK-02: Scheduler and Runtime Flow

Buku ini membahas perilaku scheduler Go yang paling penting dipahami engineer: model G-M-P, penanganan blocking syscall, dan preemption yang menjaga distribusi CPU tetap sehat.

## Struktur

### [CH-01-gmp-model](./CH-01-gmp-model/)
Model G-M-P sebagai fondasi penjadwalan goroutine di atas thread OS.

### [CH-02-syscall-poller](./CH-02-syscall-poller/)
Blocking syscall, netpoller, dan dampaknya pada aliran kerja scheduler.

### [CH-03-preemption](./CH-03-preemption/)
Preemption untuk mencegah goroutine rakus memonopoli CPU terlalu lama.

## Boundary

- fokus pada perilaku scheduler yang berguna untuk engineer aplikasi dan performance tuning;
- menjadi jembatan menuju runtime internals yang lebih dalam tanpa pindah penuh ke `RAK-06`;
- bukan tempat utama untuk bedah source scheduler Go secara lengkap.

---
*Status: [x] Complete*
