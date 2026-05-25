package main

import "fmt"

func selecetionSortArray(angka *[5]int){
	var idx_min, i, j int
	for i = 0; i < len(angka) - 1; i++ { //loop luar (iterasi sortingnya)
		idx_min = i
		for j = i + 1; j < len(angka); j++ { //loop dalam (mencari nilai ekstrim)
			if angka[j] <= angka[idx_min] { //sorting terkecil ke terbesar
				idx_min = j
			}
		}
		if idx_min != i {
			angka[i], angka[idx_min] = angka[idx_min], angka[i]
		}
	}
}

func main(){
	var arrAngka [5]int

	for i := 0; i < len(arrAngka); i++ {
		fmt.Printf("Masukkan data angka ke-%d : ", i)
		fmt.Scan(&arrAngka[i])
	}
	fmt.Println()

	fmt.Println("=== SEBELUM DISORTING ===")
	for i := 0; i < len(arrAngka); i++ {
		fmt.Print(arrAngka[i], " - ")
	}

	fmt.Println("\n=== SETELAH DISORITNG ===")
	selecetionSortArray(&arrAngka)
	for i := 0; i < len(arrAngka); i++ {
		fmt.Print(arrAngka[i], " - ")
	}
}