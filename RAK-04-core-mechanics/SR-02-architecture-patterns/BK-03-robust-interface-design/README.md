# BK-03: Robust Interface Design

Buku ini membahas bagaimana interface dipakai untuk membangun kontrak yang kecil, stabil, dan tahan perubahan dalam desain sistem Go.

## Struktur

### [CH-01_InterfaceComposition](./CH-01_InterfaceComposition/)
Komposisi interface kecil sebagai fondasi abstraksi yang lebih kuat.

### [CH-02_AcceptInterfacesReturnStructs](./CH-02_AcceptInterfacesReturnStructs/)
Prinsip menerima interface dan mengembalikan struct untuk menjaga fleksibilitas tanpa kehilangan kejelasan.

### [CH-03_DependencyInversion](./CH-03_DependencyInversion/)
Dependency inversion sebagai cara memisahkan kebijakan dari detail implementasi.

## Boundary

- fokus pada desain interface yang bertahan baik saat sistem berkembang;
- membantu pembaca membedakan abstraksi yang sehat dari interface yang terlalu besar;
- bukan tempat utama untuk penjelasan dasar apa itu interface atau type assertion.

---
*Status: [x] Complete*
