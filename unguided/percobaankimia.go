package main
import "fmt"

func main() {
    berhasil := true
    var g1, g2, g3, g4 string

    for i := 1; i <= 5; i++ {
        fmt.Printf("Percobaan %d : ", i)
        fmt.Scan(&g1, &g2, &g3, &g4)

        if g1 != "merah" || g2 != "kuning" || g3 != "hijau" || g4 != "ungu" {
            berhasil = false
        }
    }

    fmt.Printf("BERHASIL : %t\n", berhasil)
}