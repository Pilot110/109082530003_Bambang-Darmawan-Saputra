package main

import "fmt"

func selectionSortAsc(arr []int, n int) {
	for i := 0; i < n-1; i++ {
		idxMin := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[idxMin] {
				idxMin = j
			}
		}
		arr[i], arr[idxMin] = arr[idxMin], arr[i]
	}
}

func selectionSortDesc(arr []int, n int) {
	for i := 0; i < n-1; i++ {
		idxMax := i
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[idxMax] {
				idxMax = j
			}
		}
		arr[i], arr[idxMax] = arr[idxMax], arr[i]
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)
		var ganjil, genap []int
		
		for j := 0; j < m; j++ {
			var val int
			fmt.Scan(&val)
			if val%2 != 0 {
				ganjil = append(ganjil, val)
			} else {
				genap = append(genap, val)
			}
		}
		
		selectionSortAsc(ganjil, len(ganjil))
		selectionSortDesc(genap, len(genap))

		// Cetak ganjil terlebih dahulu, kemudian genap
		for _, val := range ganjil {
			fmt.Print(val, " ")
		}
		for _, val := range genap {
			fmt.Print(val, " ")
		}
		fmt.Println()
	}
}