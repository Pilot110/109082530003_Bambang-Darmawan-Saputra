package main

import (
	"fmt"
	"math"
)

const NMAX = 100

func main() {
	var arr [NMAX]int
	var n, x, delIdx, cari int

	fmt.Print("Masukkan jumlah elemen (N): ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	
	fmt.Print("a. Isi array: ")
	for i := 0; i < n; i++ { fmt.Print(arr[i], " ") }
	fmt.Println()

	
	fmt.Print("b. Indeks ganjil: ")
	for i := 1; i < n; i += 2 { fmt.Print(arr[i], " ") }
	fmt.Println()

	
	fmt.Print("c. Indeks genap: ")
	for i := 0; i < n; i += 2 { fmt.Print(arr[i], " ") }
	fmt.Println()

	
	fmt.Print("d. Masukkan x untuk indeks kelipatan: ")
	fmt.Scan(&x)
	if x > 0 {
		fmt.Printf("   Indeks kelipatan %d: ", x)
		for i := x; i < n; i += x { fmt.Print(arr[i], " ") }
		fmt.Println()
	}

	
	fmt.Print("e. Indeks yang dihapus: ")
	fmt.Scan(&delIdx)
	if delIdx >= 0 && delIdx < n {
		for i := delIdx; i < n-1; i++ {
			arr[i] = arr[i+1]
		}
		n-- 
	}
	fmt.Print("   Array setelah dihapus: ")
	for i := 0; i < n; i++ { fmt.Print(arr[i], " ") }
	fmt.Println()

	sum := 0.0
	for i := 0; i < n; i++ { sum += float64(arr[i]) }
	avg := sum / float64(n)
	fmt.Printf("f. Rata-rata: %.2f\n", avg)


	var varSum float64
	for i := 0; i < n; i++ { varSum += math.Pow(float64(arr[i])-avg, 2) }
	fmt.Printf("g. Standar Deviasi: %.2f\n", math.Sqrt(varSum/float64(n)))

	fmt.Print("h. Bilangan yang dicari frekuensinya: ")
	fmt.Scan(&cari)
	freq := 0
	for i := 0; i < n; i++ {
		if arr[i] == cari { freq++ }
	}
	fmt.Printf("   Frekuensi %d: %d\n", cari, freq)
}