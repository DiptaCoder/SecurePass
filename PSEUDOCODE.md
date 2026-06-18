# Pseudocode SecurePass

## Struktur Data

```go
STRUCTURE dataAkun
    nama     : STRING
    email    : STRING
    sandi    : STRING
    tanggal  : STRING
    id       : INTEGER
ENDSTRUCTURE

STRUCTURE tabAkun
    data : ARRAY[100] OF dataAkun
ENDSTRUCTURE

DECLARE idAkun : INTEGER ← 1
DECLARE NMAX   : INTEGER ← 100
```

---

# Main Program

```go
PROCEDURE main()

    DECLARE daftar : tabAkun
    DECLARE n      : INTEGER ← 0
    DECLARE pilih  : INTEGER

    WHILE TRUE DO

        OUTPUT "=== SECUREPASS ==="
        OUTPUT "1. Tambah Akun"
        OUTPUT "2. Ubah Akun"
        OUTPUT "3. Hapus Akun"
        OUTPUT "4. Tampil Semua"
        OUTPUT "5. Cari Akun"
        OUTPUT "6. Urutkan Data"
        OUTPUT "7. Statistik"
        OUTPUT "8. Keluar"

        INPUT pilih

        CASE pilih OF

            1 : CALL tambah(daftar,n)

            2 : CALL ubah(daftar,n)

            3 : CALL hapus(daftar,n)

            4 : CALL tampil(daftar,n)

            5 : CALL cari(daftar,n)

            6 : CALL urutkan(daftar,n)

            7 : CALL statistik(daftar,n)

            8 :
                OUTPUT "Sampai jumpa"
                STOP

            OTHERWISE :
                OUTPUT "Pilihan tidak valid"

        ENDCASE

    ENDWHILE

ENDPROCEDURE
```

---

# 1. Tambah Akun

```go
PROCEDURE tambah(T,n)

    IF n >= NMAX THEN
        OUTPUT "Data penuh"
        RETURN
    ENDIF

    DECLARE d : dataAkun

    INPUT d.nama
    INPUT d.email
    INPUT d.sandi
    INPUT d.tanggal

    d.id ← idAkun
    idAkun ← idAkun + 1

    T.data[n] ← d
    n ← n + 1

    OUTPUT "Berhasil ditambahkan"

ENDPROCEDURE
```

---

# 2. Ubah Akun

```go
PROCEDURE ubah(T,n)

    INPUT id

    idx ← cariID(T,n,id)

    IF idx = -1 THEN
        OUTPUT "Akun tidak ditemukan"
        RETURN
    ENDIF

    INPUT emailBaru
    INPUT sandiBaru
    INPUT tanggalBaru

    T.data[idx].email ← emailBaru
    T.data[idx].sandi ← sandiBaru
    T.data[idx].tanggal ← tanggalBaru

    OUTPUT "Data berhasil diubah"

ENDPROCEDURE
```

---

# 3. Hapus Akun

```go
PROCEDURE hapus(T,n)

    INPUT id

    idx ← cariID(T,n,id)

    IF idx = -1 THEN
        OUTPUT "Akun tidak ditemukan"
        RETURN
    ENDIF

    FOR i ← idx TO n-2 DO
        T.data[i] ← T.data[i+1]
    ENDFOR

    n ← n - 1

    OUTPUT "Data berhasil dihapus"

ENDPROCEDURE
```

---

# 4. Tampilkan Semua Akun

```go
PROCEDURE tampil(T,n)

    IF n = 0 THEN
        OUTPUT "Belum ada data"
        RETURN
    ENDIF

    OUTPUT "Daftar Akun"

    FOR i ← 0 TO n-1 DO

        kekuatan ← kekuatanSandi(T.data[i].sandi)

        OUTPUT T.data[i].id
        OUTPUT T.data[i].nama
        OUTPUT T.data[i].email
        OUTPUT T.data[i].tanggal
        OUTPUT kekuatan

    ENDFOR

ENDPROCEDURE
```

---

# 5. Analisis Kekuatan Sandi

```go
FUNCTION kekuatanSandi(s)

    adaBesar ← FALSE
    adaKecil ← FALSE
    adaAngka ← FALSE
    adaSimbol ← FALSE

    FOR setiap karakter dalam s DO

        IF karakter adalah huruf besar THEN
            adaBesar ← TRUE

        ELSE IF karakter adalah huruf kecil THEN
            adaKecil ← TRUE

        ELSE IF karakter adalah angka THEN
            adaAngka ← TRUE

        ELSE
            adaSimbol ← TRUE

        ENDIF

    ENDFOR

    jumlah ← 0

    IF adaBesar THEN jumlah ← jumlah + 1
    IF adaKecil THEN jumlah ← jumlah + 1
    IF adaAngka THEN jumlah ← jumlah + 1
    IF adaSimbol THEN jumlah ← jumlah + 1

    IF jumlah <= 1 THEN
        RETURN "Lemah"

    ELSE IF jumlah = 2 THEN
        RETURN "Sedang"

    ELSE
        RETURN "Kuat"
    ENDIF

ENDFUNCTION
```

---

# 6. Pencarian Akun

```go
PROCEDURE cari(T,n)

    OUTPUT "1. Sequential Search"
    OUTPUT "2. Binary Search"

    INPUT pilihan
    INPUT nama

    IF pilihan = 1 THEN
        CALL cariSeq(T,n,nama)

    ELSE IF pilihan = 2 THEN
        CALL cariBiner(T,n,nama)

    ELSE
        OUTPUT "Pilihan tidak valid"
    ENDIF

ENDPROCEDURE
```

---

## Sequential Search

```go
PROCEDURE cariSeq(T,n,nama)

    ditemukan ← FALSE

    FOR i ← 0 TO n-1 DO

        IF T.data[i].nama = nama THEN

            OUTPUT data akun

            ditemukan ← TRUE

        ENDIF

    ENDFOR

    IF ditemukan = FALSE THEN
        OUTPUT "Tidak ditemukan"
    ENDIF

ENDPROCEDURE
```

---

## Binary Search

```go
PROCEDURE cariBiner(T,n,nama)

    kiri ← 0
    kanan ← n - 1

    WHILE kiri <= kanan DO

        tengah ← (kiri + kanan) DIV 2

        IF nama = T.data[tengah].nama THEN

            OUTPUT data ditemukan
            RETURN

        ELSE IF nama < T.data[tengah].nama THEN

            kanan ← tengah - 1

        ELSE

            kiri ← tengah + 1

        ENDIF

    ENDWHILE

    OUTPUT "Tidak ditemukan"

ENDPROCEDURE
```

---

# 7. Pengurutan Data

```go
PROCEDURE urutkan(T,n)

    OUTPUT "1. Selection Sort"
    OUTPUT "2. Insertion Sort"

    INPUT pilihan

    IF pilihan = 1 THEN

        CALL selectionSort(T,n)

    ELSE IF pilihan = 2 THEN

        CALL insertionSort(T,n)

    ELSE

        OUTPUT "Pilihan tidak valid"

    ENDIF

ENDPROCEDURE
```

---

## Selection Sort

```go
PROCEDURE selectionSort(T,n)

    FOR i ← 0 TO n-2 DO

        idxMin ← i

        FOR j ← i+1 TO n-1 DO

            IF T.data[j].nama < T.data[idxMin].nama THEN

                idxMin ← j

            ENDIF

        ENDFOR

        SWAP T.data[i] WITH T.data[idxMin]

    ENDFOR

ENDPROCEDURE
```

---

## Insertion Sort

```go
PROCEDURE insertionSort(T,n)

    FOR i ← 1 TO n-1 DO

        temp ← T.data[i]
        j ← i - 1

        WHILE j >= 0 AND T.data[j].id > temp.id DO

            T.data[j+1] ← T.data[j]

            j ← j - 1

        ENDWHILE

        T.data[j+1] ← temp

    ENDFOR

ENDPROCEDURE
```

---

# 8. Statistik Akun

```go
PROCEDURE statistik(T,n)

    IF n = 0 THEN
        OUTPUT "Belum ada data"
        RETURN
    ENDIF

    lemah ← 0
    sedang ← 0
    kuat ← 0

    FOR i ← 0 TO n-1 DO

        hasil ← kekuatanSandi(T.data[i].sandi)

        IF hasil = "Lemah" THEN

            lemah ← lemah + 1

        ELSE IF hasil = "Sedang" THEN

            sedang ← sedang + 1

        ELSE

            kuat ← kuat + 1

        ENDIF

    ENDFOR

    Cari sandi terpanjang

    Cari sandi terpendek

    OUTPUT total akun
    OUTPUT jumlah sandi kuat
    OUTPUT jumlah sandi sedang
    OUTPUT jumlah sandi lemah
    OUTPUT sandi terpanjang
    OUTPUT sandi terpendek

ENDPROCEDURE
```

---

# Fungsi Pencarian ID

```go
FUNCTION cariID(T,n,id)

    FOR i ← 0 TO n-1 DO

        IF T.data[i].id = id THEN

            RETURN i

        ENDIF

    ENDFOR

    RETURN -1

ENDFUNCTION
```
