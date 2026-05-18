5package main

import "fmt"

// fungsi binary search
func binarySearch(data []int, x int) int {

	// batas kiri array
	kiri := 0

	// batas kanan array
	kanan := len(data) - 1

	// perulangan selama kiri <= kanan
	for kiri <= kanan {

		// mencari index tengah
		tengah := (kiri + kanan) / 2

		// jika data ditemukan
		if data[tengah] == x {
			return tengah

		// jika x lebih besar, cari ke kanan
		} else if data[tengah] < x {
			kiri = tengah + 1

		// jika x lebih kecil, cari ke kiri
		} else {
			kanan = tengah - 1
		}
	}

	// jika data tidak ditemukan
	return -1
}

func main() {

	var n int
	var cari int

	// input jumlah data
	fmt.Print("Jumlah data: ")
	fmt.Scan(&n)

	// membuat array sesuai jumlah data
	data := make([]int, n)

	// input isi array
	fmt.Println("Masukkan data terurut:")
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}

	// input angka yang dicari
	fmt.Print("Angka yang dicari: ")
	fmt.Scan(&cari)

	// memanggil fungsi binary search
	hasil := binarySearch(data, cari)

	// output hasil pencarian
	if hasil != -1 {
		fmt.Println("Data ditemukan di index:", hasil)
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}