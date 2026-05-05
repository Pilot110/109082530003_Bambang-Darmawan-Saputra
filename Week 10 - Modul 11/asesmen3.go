package main
import "fmt"

const nProv = 10
type NamaProv [nProv + 1]string
type PopProv [nProv + 1]int
type TumbuhProv [nProv + 1]float64

func InputData(prov *NamaProv, pop *PopProv, tumbuh *TumbuhProv) {
	fmt.Println("--- Masukkan Nama Provinsi, Populasi Provinsi, Angka Pertumbuhan Provinsi ---")
	for i := 1; i <= nProv; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i)
		fmt.Scan(&prov[i], &pop[i], &tumbuh[i])
	}
}

func ProvinsiTercepat(tumbuh TumbuhProv) int {
	idxMax := 1
	for i := 2; i <= nProv; i++ {
		if tumbuh[i] > tumbuh[idxMax] {
			idxMax = i
		}
	}
	return idxMax
}

func IndeksProvinsi(prov NamaProv, nama string) int {
	for i := 1; i <= nProv; i++ {
		if prov[i] == nama {
			return i
		}
	}
	return -1
}

func Prediksi(prov NamaProv, pop PopProv, tumbuh TumbuhProv) {
	fmt.Println("\n=== Prediksi Jumlah Penduduk Tahun Depan Pada Provinsi Dengan Pertumbuhan Diatas 2% ===")
	for i := 1; i <= nProv; i++ {
		if tumbuh[i] > 0.02 {
			hasil := float64(pop[i]) * (tumbuh[i] + 1)
			fmt.Printf("%s %.0f\n", prov[i], hasil)
		}
	}
}

func main() {
	var p NamaProv
	var o PopProv
	var t TumbuhProv
	var cari string

	InputData(&p, &o, &t)
	
	fmt.Scan(&cari)

	fmt.Printf("\nProvinsi dengan angka pertumbuhan tercepat : %s\n", p[ProvinsiTercepat(t)])
	fmt.Printf("\nData provinsi yang dicari : %s\n", cari)

	Prediksi(p, o, t)
}