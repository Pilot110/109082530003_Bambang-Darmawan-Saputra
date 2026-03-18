package main

import "fmt"

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func permutasi(n, r int) int {
	return factorial(n) / factorial(n-r)
}

func kombinasi(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var a, b, c, d int
	fmt.Print("masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("masukkan nilai b: ")
	fmt.Scan(&b)
	fmt.Print("masukkan nilai c: ")
	fmt.Scan(&c)
	fmt.Print("masukkan nilai d: ")
	fmt.Scan(&d)

	if a >= c && b >= d {
		fmt.Printf("hasil permutasi %d dan %d adalah: %d\n", a, c, permutasi(a, c))
		fmt.Printf("hasil kombinasi %d dan %d adalah: %d\n", a, c, kombinasi(a, c))
		fmt.Printf("hasil permutasi %d dan %d adalah: %d\n", b, d, permutasi(b, d))
		fmt.Printf("hasil kombinasi %d dan %d adalah: %d\n", b, d, kombinasi(b, d))
	} else {
		fmt.Println("Syarat a >= c dan b >= d tidak terpenuhi.")
	}
}