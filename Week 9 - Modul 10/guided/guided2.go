package main

import "fmt"

type mahasiswa struct {
	nama string
	nim  string
	IPK  float64
}

type arrMahasiswa [3]mahasiswa

func IPKTerbesar(array []mahasiswa) (mahasiswa, int) {
	var mhsTerbesar mahasiswa = array[0]
	var idxDitemukan int = 0
	var i int

	for i = 1; i < len(array); i++ {
		if array[i].IPK > mhsTerbesar.IPK {
			mhsTerbesar = array[i]
			idxDitemukan = i
		}
	}

	return mhsTerbesar, idxDitemukan
}

func IPKTerkecil(array []mahasiswa) (mahasiswa, int) {
	var mhsTerkecil mahasiswa = array[0]
	var idxDitemukan int = 0
	var i int

	for i = 1; i < len(array); i++ {
		if array[i].IPK < mhsTerkecil.IPK {
			mhsTerkecil = array[i]
			idxDitemukan = i
		}
	}
	return mhsTerkecil, idxDitemukan
}

func main() {
	var arrMhs arrMahasiswa
	var i, j int
	for i = 0; i < len(arrMhs); i++ {
		fmt.Println("DATA MAHASISWA INDEKS KE-", i)
		fmt.Print("Masukkan nama : ")
		fmt.Scan(&arrMhs[i].nama)
		fmt.Print("Masukkan NIM :")
		fmt.Scan(&arrMhs[i].nim)
		fmt.Print("Masukkan IPK : ")
		fmt.Scan(&arrMhs[i].IPK)
		fmt.Println("========================")
	}

	fmt.Println()
	for j = 0; j < len(arrMhs); j++ {
		fmt.Println("DATA MAHASISWA INDEKS KE-", j)
		fmt.Println("Nama : ", arrMhs[j].nama)
		fmt.Println("NIM : ", arrMhs[j].nim)
		fmt.Println("IPK : ", arrMhs[j].IPK)
		fmt.Println("========================")
	}
	fmt.Println()

	var mhsMaks, mhsMin mahasiswa
	var idxMaks, idxMin int

	mhsMaks, idxMaks = IPKTerbesar(arrMhs[:])
	mhsMin, idxMin = IPKTerkecil(arrMhs[:])

	fmt.Println("=== DATA IPK TERKECIL & TERBESAR MAHASISWA ===")
	fmt.Printf("IPK Terbesar : %.2f, atas nama %s dengan NIM %s, ditemukan pada indeks ke-%d", mhsMaks.IPK, mhsMaks.nama, mhsMaks.nim, idxMaks)
	fmt.Printf("IPK Terkecil : %.2f, atas nama %s dengan NIM %s, ditemukan pada indeks ke-%d", mhsMin.IPK, mhsMin.nama, mhsMin.nim, idxMin)
}
