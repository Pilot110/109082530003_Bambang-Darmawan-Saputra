package main

import "fmt"

const nMax int = 4321

type arrInt [nMax]int

func insertionSort(T *arrInt, n int) {
	/* I.S. terdefinisi array T yang berisi n bilangan bulat
	   F.S. array T terurut secara mengecil atau descending dengan INSERTION SORT */

	var temp, i, j int

	i = 1
	for i <= n-1 {
		j = i
		temp = T[j]

		for j > 0 && temp > T[j-1] {
			T[j] = T[j-1]
			j = j - 1
		}

		T[j] = temp
		i = i + 1
	}
}

func main() {
	var data arrInt
	var n int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}

	insertionSort(&data, n)

	for i := 0; i < n; i++ {
		fmt.Print(data[i], " ")
	}
}
