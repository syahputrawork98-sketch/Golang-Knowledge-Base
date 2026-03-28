# BK-04: Concurrency Philosophy

Buku ini membahas filosofi concurrency Go dari sisi desain: kenapa komunikasi lebih diutamakan daripada berbagi memori, dan bagaimana cara pandang ini membentuk idiom Go.

## Struktur

### [CH-01_ShareByCommunicating](./CH-01_ShareByCommunicating/)
Prinsip komunikasi antar goroutine sebagai dasar model concurrent Go.

### [CH-02_Don'tCommunicateBySharing](./CH-02_Don'tCommunicateBySharing/)
Alasan desain di balik semboyan Go tentang komunikasi dan ownership data.

## Boundary

- fokus pada filosofi concurrency dan alasan desain yang mendasarinya;
- membantu pembaca memahami pola pikir Go sebelum masuk ke pattern concurrent yang lebih kompleks;
- bukan tempat utama untuk scheduler, work stealing, atau internals runtime concurrency.

---
*Status: [/] Partial*
