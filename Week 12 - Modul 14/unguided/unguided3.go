package main

import "fmt"

// Fungsi algoritma insertion sort membesar
func insertionSort(arr []int, n int) {
	for i := 1; i < n; i++ {
		temp := arr[i]
		j := i
		for j > 0 && temp < arr[j-1] {
			arr[j] = arr[j-1]
			j--
		}
		arr[j] = temp
	}
}

func main() {
	var arr []int
	var val int
	
	// Membaca masukan hingga ditemukan bilangan negatif
	for {
		fmt.Scan(&val)
		if val < 0 {
			break
		}
		arr = append(arr, val)
	}

	n := len(arr)
	insertionSort(arr, n)

	// Mencetak array yang telah terurut
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	// Memeriksa status jarak antar data
	if n < 2 {
		fmt.Println("Data berjarak tidak tetap")
	} else {
		diff := arr[1] - arr[0]
		isConstant := true
		for i := 2; i < n; i++ {
			if arr[i]-arr[i-1] != diff {
				isConstant = false
				break
			}
		}
		if isConstant {
			fmt.Printf("Data berjarak %d\n", diff)
		} else {
			fmt.Println("Data berjarak tidak tetap")
		}
	}
}