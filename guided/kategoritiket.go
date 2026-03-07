package main
import "fmt"

func main() {
    var usia, harga int

    fmt.Print("Masukkan usia penonton: ")
    fmt.Scanln(&usia)

    if usia < 12 {
        harga = 30000
    } else if usia >= 12 && usia <= 17 {
        harga = 40000
    } else {
        harga = 50000
    }

    fmt.Printf("Usia penonton: %d tahun\n", usia)
    fmt.Printf("Harga tiket yang harus dibayar: Rp%d\n", harga)
}