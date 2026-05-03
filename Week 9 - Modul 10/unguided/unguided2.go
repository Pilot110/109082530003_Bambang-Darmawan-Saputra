package main
import "fmt"

func main() {
    var x, y int
    var beratIkan [1000]float64
    
    fmt.Scan(&x, &y)
    for i := 0; i < x; i++ {
        fmt.Scan(&beratIkan[i])
    }

    var totalSeluruhBerat float64 = 0
    var jumlahWadah int = 0
    var beratWadah float64 = 0

    
    for i := 0; i < x; i++ {
        beratWadah += beratIkan[i]
        totalSeluruhBerat += beratIkan[i]
        
        
        if (i+1)%y == 0 || i == x-1 {
            fmt.Printf("%.2f ", beratWadah)
            beratWadah = 0
            jumlahWadah++
        }
    }
    fmt.Println()

    
    fmt.Printf("%.2f\n", totalSeluruhBerat/float64(jumlahWadah))
}