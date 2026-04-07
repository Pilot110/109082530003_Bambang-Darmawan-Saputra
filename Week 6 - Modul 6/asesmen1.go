package main

import (
	"fmt"
	"math"
)

const pi float64 = 3.14

func volume(r, t float64) float64 {
	
	return pi * r * r * t
}

func massa(r, t, p float64) float64 {

	return volume(r, t) * p
}

func display(m1, m2 float64) {
	if m1 == m2 {
		fmt.Println("BALANCE")
	} else {
		
		selisih := math.Abs(m1 - m2)
		fmt.Printf("Selisih massa zat cair kiri dan massa zat cair kanan : %g\n", selisih)
	}
}

func main() {
	var r, tKiri, tKanan, mjKiri, mjKanan, massaKiri, massaKanan float64

	fmt.Print("Masukkan jari-jari alas tabung : ")
	fmt.Scan(&r)

	fmt.Print("Masukkan tinggi zat cair tabung kiri : ")
	fmt.Scan(&tKiri)
	fmt.Print("Masukkan massa jenis zat cair tabung kiri : ")
	fmt.Scan(&mjKiri)

	fmt.Print("Masukkan tinggi zat cair tabung kanan : ")
	fmt.Scan(&tKanan)
	fmt.Print("Masukkan massa jenis zat cair tabung kanan : ")
	fmt.Scan(&mjKanan)

	massaKiri = massa(r, tKiri, mjKiri)
	massaKanan = massa(r, tKanan, mjKanan)

	display(massaKiri, massaKanan)
}