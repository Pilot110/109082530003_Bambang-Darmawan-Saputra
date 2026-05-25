package main

import "fmt"

const NMAX = 1000000

type arrInt [NMAX]int

func SelectionSort(T *arrInt, n int) {
	for i := 0; i < n-1; i++ {
		idxMin := i
		for j := i + 1; j < n; j++ {
			if T[j] < T[idxMin] {
				idxMin = j
			}
		}
		T[i], T[idxMin] = T[idxMin], T[i]
	}
}

func median(T arrInt, n int) float64 {
	if n == 0 {
		return 0
	}
	if n%2 != 0 {
		return float64(T[n/2])
	}
	// Menggunakan float64 agar menghasilkan desimal seperti 21.5 pada contoh soal
	return float64(T[(n-1)/2]+T[n/2]) / 2.0
}

func main() {
	var A arrInt
	var x int
	var n int = 0
	
	fmt.Scan(&x)
	for x != -5313541 && n < NMAX {
		if x == 0 {
			SelectionSort(&A, n)
			fmt.Println(median(A, n))
		} else {
			A[n] = x
			n++
		}
		fmt.Scan(&x)
	}
}