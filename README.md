# SECUREPASS

## Sistem Manajemen Akun Digital Berbasis Terminal

SecurePass adalah aplikasi berbasis bahasa pemrograman Go (Golang) yang digunakan untuk mengelola data akun digital melalui terminal. Program ini memungkinkan pengguna menyimpan, mengubah, menghapus, mencari, mengurutkan, dan menganalisis data akun secara sederhana dan terstruktur.

Aplikasi ini dikembangkan sebagai Tugas Besar Mata Kuliah Algoritma dan Pemrograman.

---

# Daftar Isi

- [Deskripsi Program](#deskripsi-program)
- [Tujuan Program](#tujuan-program)
- [Fitur Aplikasi](#fitur-aplikasi)
- [Struktur Data](#struktur-data)
- [Struktur Penyimpanan Data](#struktur-penyimpanan-data)
- [Cara Menjalankan Program](#cara-menjalankan-program)
- [Panduan Penggunaan](#panduan-penggunaan)
- [Fungsi Pendukung](#fungsi-pendukung)
- [Algoritma yang Digunakan](#algoritma-yang-digunakan)
- [Struktur Repository](#struktur-repository)
- [Contoh Penggunaan](#contoh-penggunaan)
- [Catatan Teknis](#catatan-teknis)
- [Teknologi](#teknologi)
- [Anggota Tim](#anggota-tim)
- [Dokumentasi](#dokumentasi)

---

# Deskripsi Program

SecurePass merupakan aplikasi manajemen akun digital yang digunakan untuk menyimpan berbagai informasi akun dalam satu sistem. Setiap akun memiliki data berupa nama layanan, email, kata sandi, tanggal pembaruan, dan ID unik yang dibuat secara otomatis.

Program menerapkan beberapa konsep dasar Algoritma dan Pemrograman, seperti:

- Array
- Struct
- Function
- Percabangan (if-else)
- Perulangan (for)
- Sequential Search
- Binary Search
- Selection Sort
- Insertion Sort

---

# Tujuan Program

Program SecurePass dibuat untuk membantu pengguna mengelola data akun digital secara lebih terorganisir.

Tujuan utama aplikasi ini adalah:

- Menyimpan data akun digital dalam satu tempat.
- Mempermudah pencarian akun berdasarkan nama layanan.
- Mempermudah pengurutan data akun.
- Menampilkan statistik data akun yang tersimpan.
- Memberikan informasi mengenai kekuatan kata sandi yang digunakan.

---

# Fitur Aplikasi

| No | Fitur | Deskripsi |
|----|--------|-----------|
| 1 | Tambah Akun | Menambahkan data akun baru ke sistem |
| 2 | Ubah Akun | Mengubah data akun berdasarkan ID |
| 3 | Hapus Akun | Menghapus akun berdasarkan ID |
| 4 | Tampilkan Semua Akun | Menampilkan seluruh akun yang tersimpan |
| 5 | Cari Akun | Mencari akun berdasarkan nama layanan |
| 6 | Urutkan Data | Mengurutkan data akun menggunakan algoritma sorting |
| 7 | Statistik | Menampilkan statistik akun yang tersimpan |
| 8 | Keluar | Mengakhiri program |

---

# Struktur Data

Program menggunakan struct berikut:

```go
type dataAkun struct {
    nama    string
    email   string
    sandi   string
    tanggal string
    id      int
}
```

Struct tersebut digunakan untuk menyimpan data satu akun.

Data akun disimpan dalam array menggunakan struct:

```go
type tabAkun struct {
    data [NMAX]dataAkun
}
```

---

# Struktur Penyimpanan Data

Program menggunakan array statis sebagai media penyimpanan data akun.

```go
const NMAX = 100
```

Artinya program dapat menyimpan maksimal 100 akun.

Setiap data akun disimpan dalam array dan diakses menggunakan indeks array.

Pendekatan ini dipilih karena sesuai dengan materi Array dan Struct pada mata kuliah Algoritma dan Pemrograman.

---

# Cara Menjalankan Program

## Prasyarat

Pastikan Go (Golang) telah terinstal pada perangkat.

Cek versi Go dengan perintah:

```bash
go version
```

## Menjalankan Program

Buka terminal pada folder project kemudian jalankan:

```bash
go run main.go
```

Program akan menampilkan menu utama SecurePass.

---

# Panduan Penggunaan

## Menu Utama

```text
=== SECUREPASS ===
1. Tambah Akun
2. Ubah Akun
3. Hapus Akun
4. Tampil Semua
5. Cari Akun
6. Urutkan Data
7. Statistik
8. Keluar
```

---

## 1. Tambah Akun

Digunakan untuk menambahkan akun baru.

Data yang dimasukkan:

- Nama layanan
- Email
- Kata sandi
- Tanggal pembaruan

Program akan membuat ID secara otomatis dan menyimpan data ke dalam array.

---

## 2. Ubah Akun

Digunakan untuk mengubah data akun yang sudah tersimpan.

Langkah-langkah:

1. Masukkan ID akun.
2. Sistem mencari akun berdasarkan ID.
3. Jika ditemukan, pengguna memasukkan data baru.
4. Data akun diperbarui.

---

## 3. Hapus Akun

Digunakan untuk menghapus akun berdasarkan ID.

Langkah-langkah:

1. Masukkan ID akun.
2. Sistem mencari akun berdasarkan ID.
3. Jika ditemukan, data dihapus.
4. Data setelahnya digeser satu posisi ke depan.

---

## 4. Tampilkan Semua Akun

Menampilkan seluruh akun yang tersimpan beserta informasi:

- ID akun
- Nama layanan
- Email
- Kata sandi
- Kategori kekuatan sandi
- Tanggal pembaruan

---

## 5. Cari Akun

Pencarian dilakukan berdasarkan nama layanan.

Tersedia dua metode:

### Sequential Search

Melakukan pencarian dengan memeriksa setiap data satu per satu.

### Binary Search

Melakukan pencarian dengan membagi area pencarian menjadi dua bagian secara berulang.

**Catatan:** Data harus diurutkan berdasarkan nama layanan terlebih dahulu sebelum menggunakan Binary Search.

---

## 6. Urutkan Data

Tersedia dua metode pengurutan:

### Selection Sort

Mengurutkan data berdasarkan nama layanan dari A sampai Z.

### Insertion Sort

Mengurutkan data berdasarkan ID akun (urutan input).

---

## 7. Statistik

Menampilkan informasi:

- Total akun
- Jumlah akun dengan sandi kuat
- Jumlah akun dengan sandi sedang
- Jumlah akun dengan sandi lemah
- Sandi terpanjang
- Sandi terpendek

---

## 8. Keluar

Mengakhiri penggunaan aplikasi dan menutup program.

---

# Fungsi Pendukung

Selain fitur utama, program menggunakan beberapa fungsi pendukung.

| Fungsi | Kegunaan |
|----------|----------|
| `kekuatanSandi()` | Menentukan kategori kekuatan kata sandi |
| `samaStr()` | Membandingkan dua string tanpa membedakan huruf besar dan kecil |
| `lebihKecil()` | Membandingkan urutan alfabet dua string |
| `cariID()` | Mencari posisi akun berdasarkan ID |

### Kategori Kekuatan Kata Sandi

Fungsi `kekuatanSandi()` mengelompokkan kata sandi menjadi:

#### Lemah

Memiliki satu jenis karakter.

Contoh:

```text
abcdef
```

#### Sedang

Memiliki dua jenis karakter.

Contoh:

```text
abc123
```

#### Kuat

Memiliki tiga atau lebih jenis karakter.

Contoh:

```text
Abc123!
```

Fungsi ini digunakan pada fitur:

- Tampilkan Semua Akun
- Cari Akun
- Statistik

---

# Algoritma yang Digunakan

## Sequential Search

Digunakan pada fungsi:

```go
cariSeq()
```

Mencari data dengan memeriksa seluruh elemen array secara berurutan.

---

## Binary Search

Digunakan pada fungsi:

```go
cariBiner()
```

Mencari data pada array yang telah terurut dengan membagi area pencarian menjadi dua bagian secara berulang.

---

## Selection Sort

Digunakan pada fungsi:

```go
selectionSort()
```

Mengurutkan data berdasarkan nama layanan secara alfabetis.

---

## Insertion Sort

Digunakan pada fungsi:

```go
insertionSort()
```

Mengurutkan data berdasarkan ID akun.

---

# Struktur Repository

```text
SecurePass
│
├── README.md
├── main.go
├── Pseudocode.md
├── Flowchart.pdf
└── Laporan_Tubes.pdf
```

Keterangan:

- README.md → Dokumentasi repository
- main.go → Source code program
- Pseudocode.md → Pseudocode seluruh fitur
- Flowchart.pdf → Flowchart aplikasi
- Laporan_Tubes.pdf → Laporan tugas besar

---

# Contoh Penggunaan

### Menambahkan Akun

```text
Nama layanan : Gmail
Email        : user@gmail.com
Kata sandi   : Abc123!
Tgl update   : 01-06-2026
```

### Menampilkan Data

```text
ID  Layanan              Email
1   Gmail                user@gmail.com
```

### Menampilkan Statistik

```text
Total akun      : 1
Kuat            : 1
Sedang          : 0
Lemah           : 0
```

---

# Catatan Teknis

- Data disimpan dalam memori selama program berjalan.
- Data akan hilang setelah program ditutup.
- Kapasitas maksimum penyimpanan adalah 100 akun.
- Binary Search memerlukan data yang sudah terurut.
- Program menggunakan `fmt.Scan()`, sehingga input tidak mendukung spasi dalam satu field.

---

# Teknologi

- Bahasa Pemrograman : Go (Golang)
- Package yang digunakan : fmt
- Tipe Aplikasi : Console Application
- Platform : Windows, Linux, dan macOS

---

# Anggota Tim

### 108072500065

I Nyoman Ari Pradhipta Widnyana Tangkas

### 108072500088

Farrel Ivander Budi Setiawan

---

# Dokumentasi

Repository ini berisi:

- Source code program (`main.go`)
- Dokumentasi pseudocode (`Pseudocode.md`)
- Flowchart aplikasi (`Flowchart.pdf`)
- Laporan tugas besar (`Laporan_Tubes.pdf`)
- Dokumentasi repository (`README.md`)

---

**Dibuat sebagai Tugas Besar Mata Kuliah Algoritma dan Pemrograman.**
