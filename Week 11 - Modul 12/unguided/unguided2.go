package main

import "fmt"

func main() {
	var suara, suaraMasuk, suaraSah int
	var hitungSuara [21]int

	for {
		fmt.Scan(&suara)
		if suara == 0 {
			break
		}
		suaraMasuk++
		if suara >= 1 && suara <= 20 {
			suaraSah++
			hitungSuara[suara]++
		}
	}

	fmt.Printf("Suara masuk: %d\n", suaraMasuk)
	fmt.Printf("Suara sah: %d\n", suaraSah)

	ketua := -1
	wakil := -1

	// Mencari Ketua RT (Suara terbanyak, nomor terkecil jika seri)
	for i := 1; i <= 20; i++ {
		if hitungSuara[i] > 0 {
			if ketua == -1 || hitungSuara[i] > hitungSuara[ketua] {
				ketua = i
			}
		}
	}

	// Mencari Wakil RT (Suara terbanyak kedua, nomor terkecil jika seri)
	for i := 1; i <= 20; i++ {
		if hitungSuara[i] > 0 && i != ketua {
			if wakil == -1 || hitungSuara[i] > hitungSuara[wakil] {
				wakil = i
			}
		}
	}

	fmt.Printf("Ketua RT: %d\n", ketua)
	fmt.Printf("Wakil ketua: %d\n", wakil)
}