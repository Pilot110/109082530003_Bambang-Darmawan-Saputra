package main

import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB, match int
	var pemenang [100]string


	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)


	for {
		fmt.Printf("Pertandingan %d: ", match+1)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}


		if skorA > skorB {
			pemenang[match] = klubA
		} else if skorB > skorA {
			pemenang[match] = klubB
		} else {
			pemenang[match] = "Draw"
		}
		match++
	}

	for i := 0; i < match; i++ {
		fmt.Printf("Hasil %d: %s\n", i+1, pemenang[i])
	}
	fmt.Println("Pertandingan selesai")
}