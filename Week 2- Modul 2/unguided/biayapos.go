package main

import "fmt"

func main() {
	var totalGram, kg, sisaGram int
	fmt.Print("Masukkan total berat (gram): ")
	fmt.Scan(&totalGram)

	kg = totalGram / 1000
	sisaGram = totalGram % 1000

	biayaKg := kg * 10000
	biayaSisa := 0

	
	if kg > 10 {
	
		biayaSisa = 0
	} else {
	
		if sisaGram >= 500 {
			biayaSisa = sisaGram * 5 
		} else {
			biayaSisa = sisaGram * 15 
		}
	}

	totalBiaya := biayaKg + biayaSisa

	fmt.Println("===== Detail Perhitungan =====")
	fmt.Printf("Detail berat : %d kg + %d gram\n", kg, sisaGram)
	fmt.Printf("Detail biaya : Rp. %d + Rp. %d\n", biayaKg, biayaSisa)
	fmt.Printf("Total biaya: Rp %d\n", totalBiaya)
}