package main

import "fmt"

// Fungsi algoritma selection sort membesar (ascending)
func selectionSortAsc(arr []int, n int) {
	for i := 0; i < n-1; i++ {
		idxMin := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[idxMin] {
				idxMin = j
			}
		}
		// Proses pertukaran nilai (swap)
		arr[i], arr[idxMin] = arr[idxMin], arr[i]
	}
}

func main() {
	var n int
	fmt.Scan(&n) // Membaca jumlah daerah	

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m) // Membaca jumlah rumah di daerah tersebut
		arr := make([]int, m)
		
		for j := 0; j < m; j++ {
			fmt.Scan(&arr[j])
		}
		
		selectionSortAsc(arr, m)
		
		for j := 0; j < m; j++ {
			fmt.Print(arr[j], " ")
		}
		fmt.Println()
	}
}