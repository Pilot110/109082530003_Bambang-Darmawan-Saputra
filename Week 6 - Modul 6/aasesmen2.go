package main

import "fmt"

func hitungBiaya(jumlah, lama int, tujuan string) float64 {
	var biayaPerHari, totalBiaya float64
	var batasHari int

	biayaPerHari = 35000*2 + 250000 + 300000

	if tujuan == "domestik" {
		batasHari = 3
	} else if tujuan == "mancanegara" {
		batasHari = 8
		biayaPerHari = biayaPerHari * 1.5
	}

	hariDitanggung := lama
	if lama > batasHari {
		hariDitanggung = batasHari
	}

	totalBiaya = float64(jumlah) * float64(hariDitanggung) * biayaPerHari
	return totalBiaya
}

func main() {
	var jumlah, lama int
	var tujuan string

	fmt.Print("masukkan jumlah mahasiswa : ")
	fmt.Scan(&jumlah)
	fmt.Print("Masukkan lama hari study tour : ")
	fmt.Scan(&lama)
	fmt.Print("Masukkan tujuan study tour (domestik/mancanegara) : ")
	fmt.Scan(&tujuan)

	biayaTotal := hitungBiaya(jumlah, lama, tujuan)

	fmt.Printf("\nBiaya perjalanan yang harus dikeluarkan Tel-U : Rp. %.0f\n", biayaTotal)
}