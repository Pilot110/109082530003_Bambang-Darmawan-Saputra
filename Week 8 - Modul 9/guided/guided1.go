package main

import "fmt"

func main() {

    var rakSepatu = [3]string{"Nike", "Adidas", "Jordan"}
    fmt.Println("Array:", rakSepatu)

    var keranjangBelanja []string 

    keranjangBelanja = append(keranjangBelanja, "Apel")
    keranjangBelanja = append(keranjangBelanja, "Mangga", "Jeruk")

    fmt.Println("Slice awal:", keranjangBelanja)
    fmt.Println("Jumlah data (len):", len(keranjangBelanja)) 


    potongan := keranjangBelanja[0:2] 
    fmt.Println("Hasil Slicing [0:2]:", potongan)
}