package main

import "fmt"

func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

func main() {
	var x, y, z int
	fmt.Print("Masukkan nilai x : ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai y : ")
	fmt.Scan(&y)
	fmt.Print("Masukkan nilai z : ")
	fmt.Scan(&z)

	fmt.Printf("F(G(H( %d ))) : %d\n", x, f(g(h(x))))
	fmt.Printf("G(H(F( %d ))) : %d\n", y, g(h(f(y))))
	fmt.Printf("H(F(G( %d ))) : %d\n", z, h(f(g(z))))
}