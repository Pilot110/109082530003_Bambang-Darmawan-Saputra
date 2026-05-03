package main

import "fmt"

func nilaiTerbesar(array []int) (int, int) {
	var idx int
	var nilaiTerbesar int = array[0]
	var idxDitemukan int = 0
	for idx = 1; idx < len(array); idx++ {
		if array[idx] > nilaiTerbesar {
			nilaiTerbesar = array[idx]
			idxDitemukan = idx
		}
	}
	return nilaiTerbesar, idxDitemukan
}

func nilaiTerkecil(array []int) (int, int) {
	var idx int
	var nilaiTerkecil int = array[0]
	var idxDitemukan int = 0
	for idx = 1; idx < len(array); idx++ {
		if array[idx] < nilaiTerkecil {
			nilaiTerkecil = array[idx]
			idxDitemukan = idx
		}
	}
	return nilaiTerkecil, idxDitemukan
}

func main() {
	var arrAngka [5]int = [5]int{45, 23, 98, 66, 12}

	var nilaiMaks, idxNilaiMaks int = nilaiTerbesar(arrAngka[:])
	var nilaiMin, idxNilaiMin int = nilaiTerkecil(arrAngka[:])

	fmt.Printf("Nilai terkecil : %d ditemukan pada index ke-%d", nilaiMin, idxNilaiMin)
	fmt.Println()
	fmt.Printf("Nilai terbesar : %d ditemukan pada index ke-%d", nilaiMaks, idxNilaiMaks)

}
