package main

import "fmt"

type arrData [5]string //tipe data alias

func seqSearch(arr arrData, binatangCari string) int {
	var found int = -1
	for i := 0; i < len(arr); i++ {
		if arr[i] == binatangCari {
			found = i
			break
		}
	}
	return found
}

func main() {
	var arrBinatang arrData //array tipe string

	for i := 0; i < len(arrBinatang); i++ {
		fmt.Printf("Masukkan data binatang indeks ke-%d : ", i)
		fmt.Scan(&arrBinatang[i])
	}
	fmt.Println()

	var binatangCari string
	fmt.Print("Masukkan nama binatang yang mau dicari : ")
	fmt.Scan(&binatangCari)

	var idxCari int
	idxCari = seqSearch(arrBinatang, binatangCari)

	if idxCari > -1 {
		fmt.Printf("Data %s ditemukan pada indeks ke-%d!", binatangCari, idxCari)
	} else if idxCari == -1 {
		fmt.Printf("Data %s tidak ditemukan!", binatangCari)
	}

}