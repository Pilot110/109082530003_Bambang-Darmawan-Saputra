package main

import "fmt"

func hitungLuasPersegiPanjang(panjang int, lebar int) int {
	luas := panjang * lebar
	return luas
}

func main() {
	p := 10
	l := 5
	
	hasil := hitungLuasPersegiPanjang(p, l)
	
	fmt.Printf("Luas persegi panjang dengan panjang %d dan lebar %d adalah: %d\n", p, l, hasil)
}