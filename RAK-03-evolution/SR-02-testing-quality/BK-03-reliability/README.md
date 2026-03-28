# BK-03: Reliability

Buku ini membahas teknik untuk meningkatkan keandalan kode Go melalui fuzz testing, race detection, dan coverage analysis sebagai alat ukur confidence yang lebih tajam.

## Struktur

### [CH-01-fuzz-testing](./CH-01-fuzz-testing/)
Fuzz testing untuk menemukan input aneh dan bug yang tidak terpikir oleh test biasa.

### [CH-02-race-detector](./CH-02-race-detector/)
Race detector untuk menemukan konflik akses memory pada kode konkuren.

### [CH-03-coverage](./CH-03-coverage/)
Coverage analysis untuk melihat bagian kode mana yang benar-benar terjamah oleh test.

## Boundary

- fokus pada alat dan teknik yang meningkatkan confidence terhadap perilaku kode;
- membantu pembaca memahami batas dan manfaat tiap alat reliability;
- bukan tempat utama untuk mocking pattern atau observability runtime mendalam.

---
*Status: [x] Complete*
