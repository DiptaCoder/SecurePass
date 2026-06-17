package main

import "fmt"

const NMAX = 100

var idAkun int = 1

type dataAkun struct {
	nama    string
	email   string
	sandi   string
	tanggal string
	id      int
}

type tabAkun struct {
	data [NMAX]dataAkun
}

func kekuatanSandi(s string) string {
	adaBesar := false
	adaKecil := false
	adaAngka := false
	adaSimbol := false

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			adaBesar = true
		} else if c >= 'a' && c <= 'z' {
			adaKecil = true
		} else if c >= '0' && c <= '9' {
			adaAngka = true
		} else {
			adaSimbol = true
		}
	}

	jumlah := 0
	if adaBesar {
		jumlah++
	}
	if adaKecil {
		jumlah++
	}
	if adaAngka {
		jumlah++
	}
	if adaSimbol {
		jumlah++
	}

	if jumlah <= 1 {
		return "Lemah"
	} else if jumlah == 2 {
		return "Sedang"
	}
	return "Kuat"
}

func samaStr(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		ca, cb := a[i], b[i]
		if ca >= 'A' && ca <= 'Z' {
			ca += 32
		}
		if cb >= 'A' && cb <= 'Z' {
			cb += 32
		}
		if ca != cb {
			return false
		}
	}
	return true
}

func lebihKecil(a, b string) bool {
	min := len(a)
	if len(b) < min {
		min = len(b)
	}
	for i := 0; i < min; i++ {
		ca, cb := a[i], b[i]
		if ca >= 'A' && ca <= 'Z' {
			ca += 32
		}
		if cb >= 'A' && cb <= 'Z' {
			cb += 32
		}
		if ca < cb {
			return true
		}
		if ca > cb {
			return false
		}
	}
	return len(a) < len(b)
}

func cariID(T tabAkun, n int, id int) int {
	for i := 0; i < n; i++ {
		if T.data[i].id == id {
			return i
		}
	}
	return -1
}

func cariSeq(T tabAkun, n int, cari string) {
	ketemu := false
	fmt.Println("\nHasil pencarian:")
	fmt.Println("ID  Layanan              Email                    Sandi            Kuat  Tgl Update")
	fmt.Println("----------------------------------------------------------------------------------------")
	for i := 0; i < n; i++ {
		if samaStr(T.data[i].nama, cari) {
			k := kekuatanSandi(T.data[i].sandi)
			kode := k[:1]
			fmt.Printf("%-3d %-20s %-24s %-16s %-5s %s\n",
				T.data[i].id, T.data[i].nama, T.data[i].email,
				T.data[i].sandi, kode, T.data[i].tanggal)
			ketemu = true
		}
	}
	if !ketemu {
		fmt.Println("Tidak ditemukan.")
	} else {
		fmt.Println("Keterangan kekuatan: K=Kuat  S=Sedang  L=Lemah")
	}
}

func cariBiner(T tabAkun, n int, cari string) {
	kiri, kanan := 0, n-1
	tengahIdx := -1

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if samaStr(T.data[tengah].nama, cari) {
			tengahIdx = tengah
			break
		} else if lebihKecil(cari, T.data[tengah].nama) {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	if tengahIdx == -1 {
		fmt.Println("Tidak ditemukan.")
		return
	}

	var hasil [NMAX]int
	jml := 0
	hasil[jml] = tengahIdx
	jml++

	i := tengahIdx - 1
	for i >= 0 && samaStr(T.data[i].nama, cari) {
		hasil[jml] = i
		jml++
		i--
	}
	i = tengahIdx + 1
	for i < n && samaStr(T.data[i].nama, cari) {
		hasil[jml] = i
		jml++
		i++
	}

	for x := 1; x < jml; x++ {
		key := hasil[x]
		y := x - 1
		for y >= 0 && hasil[y] > key {
			hasil[y+1] = hasil[y]
			y--
		}
		hasil[y+1] = key
	}

	fmt.Println("\nHasil pencarian:")
	fmt.Println("ID  Layanan              Email                    Sandi            Kuat  Tgl Update")
	fmt.Println("----------------------------------------------------------------------------------------")
	for x := 0; x < jml; x++ {
		idx := hasil[x]
		k := kekuatanSandi(T.data[idx].sandi)
		kode := k[:1]
		fmt.Printf("%-3d %-20s %-24s %-16s %-5s %s\n",
			T.data[idx].id, T.data[idx].nama, T.data[idx].email,
			T.data[idx].sandi, kode, T.data[idx].tanggal)
	}
	fmt.Println("Keterangan kekuatan: K=Kuat  S=Sedang  L=Lemah")
}

func selectionSort(T *tabAkun, n int) {
	for i := 0; i < n-1; i++ {
		idxMin := i
		for j := i + 1; j < n; j++ {
			if lebihKecil(T.data[j].nama, T.data[idxMin].nama) {
				idxMin = j
			}
		}
		T.data[i], T.data[idxMin] = T.data[idxMin], T.data[i]
	}
}

func insertionSort(T *tabAkun, n int) {
	for i := 1; i < n; i++ {
		temp := T.data[i]
		j := i - 1
		for j >= 0 && T.data[j].id > temp.id {
			T.data[j+1] = T.data[j]
			j--
		}
		T.data[j+1] = temp
	}
}

func tambah(T *tabAkun, n *int) {
	if *n >= NMAX {
		fmt.Println("Data penuh!")
		return
	}
	var d dataAkun
	fmt.Print("Nama layanan : ")
	fmt.Scan(&d.nama)
	fmt.Print("Email        : ")
	fmt.Scan(&d.email)
	fmt.Print("Kata sandi   : ")
	fmt.Scan(&d.sandi)
	fmt.Print("Tgl update   : ")
	fmt.Scan(&d.tanggal)

	d.id = idAkun
	idAkun++

	T.data[*n] = d
	*n++
	fmt.Println("Berhasil ditambahkan!")
}

func ubah(T *tabAkun, n int) {
	var id int
	fmt.Print("ID akun yang ingin diubah: ")
	fmt.Scan(&id)
	idx := cariID(*T, n, id)
	if idx == -1 {
		fmt.Println("Akun tidak ditemukan!")
		return
	}
	fmt.Printf("Mengubah akun: %s (%s)\n", T.data[idx].nama, T.data[idx].email)
	fmt.Print("Email baru   : ")
	fmt.Scan(&T.data[idx].email)
	fmt.Print("Sandi baru   : ")
	fmt.Scan(&T.data[idx].sandi)
	fmt.Print("Tgl update   : ")
	fmt.Scan(&T.data[idx].tanggal)
	fmt.Println("Data berhasil diubah!")
}

func hapus(T *tabAkun, n *int) {
	var id int
	fmt.Print("ID akun yang ingin dihapus: ")
	fmt.Scan(&id)
	idx := cariID(*T, *n, id)
	if idx == -1 {
		fmt.Println("Akun tidak ditemukan!")
		return
	}
	fmt.Printf("Menghapus akun: %s (%s)\n", T.data[idx].nama, T.data[idx].email)
	for i := idx; i < *n-1; i++ {
		T.data[i] = T.data[i+1]
	}
	T.data[*n-1] = dataAkun{}
	*n--
	fmt.Println("Data berhasil dihapus!")
}

func tampil(T tabAkun, n int) {
	if n == 0 {
		fmt.Println("Belum ada data.")
		return
	}
	fmt.Println("\nID  Layanan              Email                    Sandi            Kuat  Tgl Update")
	fmt.Println("----------------------------------------------------------------------------------------")
	for i := 0; i < n; i++ {
		k := kekuatanSandi(T.data[i].sandi)
		kode := k[:1]
		fmt.Printf("%-3d %-20s %-24s %-16s %-5s %s\n",
			T.data[i].id, T.data[i].nama, T.data[i].email,
			T.data[i].sandi, kode, T.data[i].tanggal)
	}
	fmt.Println("Keterangan kekuatan: K=Kuat  S=Sedang  L=Lemah")
}

func cari(T tabAkun, n int) {
	var pilih int
	var nama string
	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search (data harus terurut A-Z)")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilih)
	fmt.Print("Nama layanan: ")
	fmt.Scan(&nama)

	if pilih == 1 {
		cariSeq(T, n, nama)
	} else if pilih == 2 {
		cariBiner(T, n, nama)
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

func urutkan(T *tabAkun, n int) {
	var pilih int
	fmt.Println("1. Selection Sort - nama layanan (A-Z)")
	fmt.Println("2. Insertion Sort - ID (urutan input)")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilih)
	if pilih == 1 {
		selectionSort(T, n)
		fmt.Println("Diurutkan berdasarkan nama (A-Z).")
	} else if pilih == 2 {
		insertionSort(T, n)
		fmt.Println("Diurutkan berdasarkan ID (urutan input).")
	} else {
		fmt.Println("Pilihan tidak valid.")
		return
	}
	tampil(*T, n)
}

func statistik(T tabAkun, n int) {
	if n == 0 {
		fmt.Println("Belum ada data.")
		return
	}
	lemah, sedang, kuat := 0, 0, 0
	for i := 0; i < n; i++ {
		k := kekuatanSandi(T.data[i].sandi)
		if k == "Lemah" {
			lemah++
		} else if k == "Sedang" {
			sedang++
		} else {
			kuat++
		}
	}
	idxMax, idxMin := 0, 0
	for i := 1; i < n; i++ {
		if len(T.data[i].sandi) > len(T.data[idxMax].sandi) {
			idxMax = i
		}
		if len(T.data[i].sandi) < len(T.data[idxMin].sandi) {
			idxMin = i
		}
	}
	fmt.Println("Total akun      :", n)
	fmt.Println("Kuat            :", kuat)
	fmt.Println("Sedang          :", sedang)
	fmt.Println("Lemah           :", lemah)
	fmt.Printf("Sandi terpanjang: %s (%d karakter)\n", T.data[idxMax].nama, len(T.data[idxMax].sandi))
	fmt.Printf("Sandi terpendek : %s (%d karakter)\n", T.data[idxMin].nama, len(T.data[idxMin].sandi))
}

func main() {
	var daftar tabAkun
	n := 0
	var pilih int

	for {
		fmt.Println("\n=== SECUREPASS ===")
		fmt.Println("1. Tambah Akun")
		fmt.Println("2. Ubah Akun")
		fmt.Println("3. Hapus Akun")
		fmt.Println("4. Tampil Semua")
		fmt.Println("5. Cari Akun")
		fmt.Println("6. Urutkan Data")
		fmt.Println("7. Statistik")
		fmt.Println("8. Keluar")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			tambah(&daftar, &n)
		case 2:
			ubah(&daftar, n)
		case 3:
			hapus(&daftar, &n)
		case 4:
			tampil(daftar, n)
		case 5:
			cari(daftar, n)
		case 6:
			urutkan(&daftar, n)
		case 7:
			statistik(daftar, n)
		case 8:
			fmt.Println("Sampai jumpa!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
