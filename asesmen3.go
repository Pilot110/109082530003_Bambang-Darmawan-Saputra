package main

import "fmt"

const NMAX = 1000000

type partai struct {
	nama  int
	suara int
}

type tabPartai [NMAX]partai

func posisi(t tabPartai, n int, nama int) int {
	// Sequential search mencari nama partai
	for i := 0; i < n; i++ {
		if t[i].nama == nama {
			return i
		}
	}
	return -1
}

func main() {
	var p tabPartai
	var n int = 0
	var x int

	fmt.Scan(&x)
	for x != -1 {
		idx := posisi(p, n, x)
		if idx != -1 {
			// Jika partai sudah ada di array, tambahkan suaranya
			p[idx].suara++
		} else {
			// Jika belum ada, masukkan sebagai partai baru
			p[n].nama = x
			p[n].suara = 1
			n++
		}
		fmt.Scan(&x)
	}

	// Proses pengurutan dengan Insertion Sort descending (berdasarkan jumlah suara)
	for i := 1; i < n; i++ {
		temp := p[i]
		j := i - 1
		for j >= 0 && temp.suara > p[j].suara {
			p[j+1] = p[j]
			j--
		}
		p[j+1] = temp
	}

	// Tampilkan array p
	for i := 0; i < n; i++ {
		fmt.Printf("%d(%d) ", p[i].nama, p[i].suara)
	}
	fmt.Println()
}