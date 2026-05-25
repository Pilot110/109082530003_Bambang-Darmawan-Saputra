package main

import "fmt"

const nMax int = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	fmt.Scan(n)
	for i := 0; i < *n; i++ {
		fmt.Scan(&pustaka[i].id, &pustaka[i].judul, &pustaka[i].penulis, &pustaka[i].penerbit, 
		         &pustaka[i].eksemplar, &pustaka[i].tahun, &pustaka[i].rating)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n == 0 {
		return
	}
	maxIdx := 0
	// Pencarian sekuensial untuk rating tertinggi pada data yang belum terurut
	for i := 1; i < n; i++ {
		if pustaka[i].rating > pustaka[maxIdx].rating {
			maxIdx = i
		}
	}
	fav := pustaka[maxIdx]
	fmt.Printf("%s, %s, %s, %d\n", fav.judul, fav.penulis, fav.penerbit, fav.tahun)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	// Insertion sort mengecil (descending) berdasarkan rating
	for i := 1; i < n; i++ {
		temp := pustaka[i]
		j := i
		for j > 0 && temp.rating > pustaka[j-1].rating {
			pustaka[j] = pustaka[j-1]
			j--
		}
		pustaka[j] = temp
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	limit := 5
	if n < 5 {
		limit = n
	}
	for i := 0; i < limit; i++ {
		fmt.Println(pustaka[i].judul)
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	left := 0
	right := n - 1
	found := false
	var mid int

	// Pencarian Biner (Binary Search)
	for left <= right && !found {
		mid = (left + right) / 2
		if pustaka[mid].rating == r {
			found = true
		} else if pustaka[mid].rating < r {
			// Karena array terurut mengecil, jika nilai tengah lebih kecil dari yang dicari,
			// maka cari di sisi kiri
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if found {
		b := pustaka[mid]
		fmt.Printf("%s, %s, %s, %d, %d, %d\n", b.judul, b.penulis, b.penerbit, b.tahun, b.eksemplar, b.rating)
	} else {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var pustaka DaftarBuku
	var n, r int

	DaftarkanBuku(&pustaka, &n)
	CetakTerfavorit(pustaka, n)
	UrutBuku(&pustaka, n)
	Cetak5Terbaru(pustaka, n)

	// Membaca rating yang ingin dicari di akhir
	fmt.Scan(&r)
	CariBuku(pustaka, n, r)
}