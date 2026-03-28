# BK-02: Mocking Patterns

Buku ini membahas cara memisahkan logika dari dependency eksternal lewat interface injection, function injection, dan integration testing dengan container nyata saat mock tidak lagi cukup.

## Struktur

### [CH-01-interface-mock](./CH-01-interface-mock/)
Mocking berbasis interface kecil untuk menjaga isolasi dan testability.

### [CH-02-func-mock](./CH-02-func-mock/)
Mocking berbasis function injection untuk kasus yang lebih ringan dan ringkas.

### [CH-03-testcontainers](./CH-03-testcontainers/)
Integration testing dengan container ephemeral saat kita perlu menguji sistem nyata, bukan tiruannya.

## Boundary

- fokus pada teknik memisahkan logika dari dependency untuk kebutuhan testing;
- membantu pembaca memilih kapan cukup memakai mock, dan kapan perlu integration test nyata;
- bukan tempat utama untuk reliability tooling seperti fuzzing, race detector, atau coverage.

---
*Status: [x] Complete*
