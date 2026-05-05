package main
import "fmt"

type mahasiswa struct {
    NIM, nama string
    nilai     int
}

func main() {
    var n int
    var t [51]mahasiswa
    
    fmt.Print("Masukkan jumlah data: ")
    fmt.Scan(&n)
    
    for i := 0; i < n; i++ {
        fmt.Printf("Masukkan data ke-%d : ", i+1)
        fmt.Scan(&t[i].NIM, &t[i].nama, &t[i].nilai)
    }
    
    var cariNIM string
    fmt.Print("Masukkan NIM mahasiswa yang ingin dicari: ")
    fmt.Scan(&cariNIM)
    
    pertama := -1
    terbesar := -1
    
    for i := 0; i < n; i++ {
        if t[i].NIM == cariNIM {
            if pertama == -1 {
                pertama = t[i].nilai
            }
            if t[i].nilai > terbesar {
                terbesar = t[i].nilai
            }
        }
    }
    
    fmt.Printf("Nilai pertama dari NIM %s adalah %d\n", cariNIM, pertama)
    fmt.Printf("Nilai terbesar dari NIM %s adalah %d\n", cariNIM, terbesar)
}