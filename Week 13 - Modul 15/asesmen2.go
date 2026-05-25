package main

import "fmt"

type Pemain struct {
	namaDepan    string
	namaBelakang string
	gol          int
	assist       int
}

func main() {
	var n int
	fmt.Scan(&n)

	var p [1001]Pemain
	for i := 0; i < n; i++ {
		// Asumsi input nama selalu terdiri dari dua kata (sesuai soal)
		fmt.Scan(&p[i].namaDepan, &p[i].namaBelakang, &p[i].gol, &p[i].assist)
	}

	// Menggunakan Insertion Sort untuk mengurutkan data secara descending
	for i := 1; i < n; i++ {
		temp := p[i]
		j := i - 1
		
		// Kriteria: Gol lebih besar, ATAU jika gol sama, assist lebih besar
		for j >= 0 && (temp.gol > p[j].gol || (temp.gol == p[j].gol && temp.assist > p[j].assist)) {
			p[j+1] = p[j]
			j--
		}
		p[j+1] = temp
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%s %s %d %d\n", p[i].namaDepan, p[i].namaBelakang, p[i].gol, p[i].assist)
	}
}