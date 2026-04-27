package main
import (
	"fmt"
	"math"
)

type Titik struct { x, y int }
type Lingkaran struct {
	pusat Titik
	r     int
}

func jarak(p, q Titik) float64 {
	return math.Sqrt(float64((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)))
}

func diDalam(c Lingkaran, p Titik) bool {
	return jarak(c.pusat, p) <= float64(c.r)
}

func main() {
	var l1, l2 Lingkaran
	var p Titik
	fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.r)
	fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.r)
	fmt.Scan(&p.x, &p.y)

	res1 := diDalam(l1, p)
	res2 := diDalam(l2, p)

	if res1 && res2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if res1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if res2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}