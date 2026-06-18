# SECUREPASS

## Sistem Manajemen Akun Digital Berbasis Terminal

SecurePass adalah aplikasi berbasis Go (Golang) yang digunakan untuk mengelola data akun digital secara sederhana melalui terminal. Program ini memungkinkan pengguna menyimpan, mengubah, menghapus, mencari, mengurutkan, dan menganalisis data akun beserta informasi kekuatan kata sandinya.

Program ini dikembangkan sebagai Tugas Besar Mata Kuliah Algoritma dan Pemrograman.

---

# Daftar Isi

- Deskripsi Program
- Fitur Aplikasi
- Struktur Data
- Cara Menjalankan Program
- Panduan Penggunaan
- Fungsi Pendukung
- Algoritma yang Digunakan
- Struktur Repository
- Contoh Penggunaan
- Catatan Teknis
- Teknologi
- Anggota Tim

---

# Deskripsi Program

SecurePass merupakan aplikasi manajemen akun digital yang digunakan untuk menyimpan berbagai data akun dalam satu sistem. Setiap akun memiliki informasi berupa nama layanan, email, kata sandi, tanggal pembaruan, dan ID unik.

Program menerapkan konsep-konsep dasar Algoritma dan Pemrograman seperti:

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

# Fitur Aplikasi

| No | Fitur | Deskripsi |
|----|--------|-----------|
| 1 | Tambah Akun | Menambahkan data akun baru ke dalam sistem |
| 2 | Ubah Akun | Mengubah data akun berdasarkan ID |
| 3 | Hapus Akun | Menghapus akun berdasarkan ID |
| 4 | Tampilkan Semua Akun | Menampilkan seluruh akun yang tersimpan |
| 5 | Pencarian Akun | Mencari akun berdasarkan nama layanan |
| 6 | Pengurutan Data Akun | Mengurutkan data akun menggunakan algoritma sorting |
| 7 | Statistik Data Akun | Menampilkan ringkasan statistik akun |
| 8 | Keluar Aplikasi | Mengakhiri program |

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

Struct tersebut digunakan untuk menyimpan informasi satu akun.

Seluruh akun disimpan dalam array melalui struct:

```go
type tabAkun struct {
    data [NMAX]dataAkun
}
```

Kapasitas maksimum penyimpanan:

```go
const NMAX = 100
```

Artinya program dapat menyimpan maksimal 100 akun.

---

# Cara Menjalankan Program

## Prasyarat

Pastikan Go (Golang) telah terinstal pada komputer.

Untuk memeriksa versi Go:

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

Fitur ini digunakan untuk menambahkan akun baru.

Data yang dimasukkan:

- Nama layanan
- Email
- Kata sandi
- Tanggal pembaruan

Program akan membuat ID secara otomatis dan menyimpan data ke dalam array.

---

## 2. Ubah Akun

Digunakan untuk mengubah data akun yang sudah tersimpan.

Langkah:

1. Masukkan ID akun.
2. Sistem mencari akun berdasarkan ID.
3. Jika ditemukan, pengguna memasukkan data baru.
4. Data akun diperbarui.

---

## 3. Hapus Akun

Digunakan untuk menghapus akun berdasarkan ID.

Langkah:

1. Masukkan ID akun.
2. Sistem mencari akun.
3. Jika ditemukan, data dihapus.
4. Data setelahnya digeser satu posisi ke depan.

---

## 4. Tampilkan Semua Akun

Menampilkan seluruh akun yang tersimpan beserta:

- ID
- Nama layanan
- Email
- Kata sandi
- Kategori kekuatan sandi
- Tanggal pembaruan

---

## 5. Pencarian Akun

Pencarian dilakukan berdasarkan nama layanan.

Metode yang tersedia:

### Sequential Search

Menelusuri seluruh data satu per satu hingga data ditemukan.

### Binary Search

Mencari data dengan membagi area pencarian menjadi dua bagian secara berulang.

Catatan:

Data harus sudah diurutkan berdasarkan nama layanan sebelum menggunakan Binary Search.

---

## 6. Pengurutan Data Akun

Tersedia dua metode pengurutan:

### Selection Sort

Mengurutkan data berdasarkan nama layanan dari A-Z.

### Insertion Sort

Mengurutkan data berdasarkan ID akun.

---

## 7. Statistik Data Akun

Menampilkan informasi:

- Total akun
- Jumlah sandi kuat
- Jumlah sandi sedang
- Jumlah sandi lemah
- Sandi terpanjang
- Sandi terpendek

---

## 8. Keluar Aplikasi

Mengakhiri penggunaan aplikasi dan menutup program.

---

# Fungsi Pendukung

## Analisis Kekuatan Kata Sandi

Program menggunakan fungsi:

```go
kekuatanSandi()
```

untuk mengelompokkan kata sandi menjadi:

### Lemah

Memiliki satu jenis karakter.

Contoh:

```text
abcdef
```

### Sedang

Memiliki dua jenis karakter.

Contoh:

```text
abc123
```

### Kuat

Memiliki tiga atau lebih jenis karakter.

Contoh:

```text
Abc123!
```

Fungsi ini digunakan pada fitur:

- Tampilkan Semua Akun
- Pencarian Akun
- Statistik Data Akun

---

# Algoritma yang Digunakan

## Sequential Search

Digunakan pada:

```go
cariSeq()
```

Fungsi mencari data akun dengan memeriksa setiap elemen array satu per satu.

Kompleksitas waktu:

```text
O(n)
```

---

## Binary Search

Digunakan pada:

```go
cariBiner()
```

Fungsi mencari data akun dengan membagi area pencarian menjadi dua bagian secara berulang.

Kompleksitas waktu:

```text
O(log n)
```

---

## Selection Sort

Digunakan pada:

```go
selectionSort()
```

Mengurutkan data berdasarkan nama layanan secara alfabetis.

Kompleksitas waktu:

```text
O(n²)
```

---

## Insertion Sort

Digunakan pada:

```go
insertionSort()
```

Mengurutkan data berdasarkan ID akun.

Kompleksitas waktu:

```text
O(n²)
```

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
- main.go → Source code utama aplikasi
- Pseudocode.md → Pseudocode seluruh fitur
- Flowchart.pdf → Flowchart aplikasi
- Laporan_Tubes.pdf → Laporan tugas besar

---

# Contoh Penggunaan

### Menambah Akun

```text
Nama layanan : Gmail
Email        : user@gmail.com
Kata sandi   : Abc123!
Tgl update   : 01-06-2026
```

### Menampilkan Data

```text
ID  Layanan  Email  Sandi  Kuat  Tgl Update
1   Gmail    user@gmail.com  Abc123!  K  01-06-2026
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

- Data disimpan di memori selama program berjalan.
- Data akan hilang ketika program ditutup.
- Kapasitas maksimum penyimpanan adalah 100 akun.
- Binary Search memerlukan data yang sudah terurut.
- Input menggunakan `fmt.Scan`, sehingga tidak mendukung spasi dalam satu field.

---

# Teknologi

- Bahasa Pemrograman : Go (Golang)
- Package yang digunakan : fmt
- Tipe Aplikasi : Console Application

---

# Anggota Tim

### 108072500065

I Nyoman Ari Pradhipta Widnyana Tangkas

### 108072500088

Farrel Ivander Budi Setiawan

---

Dibuat sebagai Tugas Besar Algoritma dan Pemrograman.
