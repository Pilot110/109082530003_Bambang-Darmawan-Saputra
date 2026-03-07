package main
import "fmt"

func main() {
    var target, tabungan, total, hari int

    fmt.Print("Masukkan target uang yang ingin dicapai: ")
    fmt.Scanln(&target)

    total = 0
    hari = 0

    for total < target {
        hari++
        fmt.Printf("Masukkan nominal tabungan hari ke-%d: ", hari)
        fmt.Scanln(&tabungan)
        
        total = total + tabungan
    }

    fmt.Printf("Selamat! Target tercapai dalam %d hari.\n", hari)
    fmt.Printf("Total tabungan Anda terkumpul: Rp%d\n", total)
}