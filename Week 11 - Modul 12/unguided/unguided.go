package main

import "fmt"

func main() {
	var suara, suaraMasuk, suaraSah int
	var hitungSuara [21]int // Indeks 1 s.d 20 untuk calon ketua RT

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
	for i := 1; i <= 20; i++ {
		if hitungSuara[i] > 0 {
			fmt.Printf("%d:%d\n", i, hitungSuara[i])
		}
	}
}