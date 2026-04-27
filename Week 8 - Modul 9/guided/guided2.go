package main

import "fmt"

func main(){
	//deklarasi map bernama nilai dengan key string dan value integer
	var nilai map[string]int = make(map[string]int)

	//menambahkan 2 data kedalam map data 1 (key = Dhimas, value = 90), data 2 (key = Ichya, value = 90)
	nilai["Dhimas"] = 90
	nilai["Ichya"] = 90

	//menampilkan seluruh isi map menggunakan perulangan for
	fmt.Println("Data nilai : ")
	var nama string //variabel bantuan yang merujuk ke key 
	var grade int //variabel bantuan yang merujuk ke value
	//perulangan untuk pencarian (narasikan = untuk setiap nama (key) & grade (value) didalam range map nilai, maka tampilkan nama = nilai)
	for nama, grade = range nilai {
		fmt.Println(nama, " = ", grade)
	}

	//operasi update (mengupdate value yang memiliki key Ichya)
	nilai["Ichya"] = 80

	//operasi searching
	var cariNama string = "hafizh" //variabek bantuan yang menyimpan data key (nama) yang akan dicari
	var isiData int //vaiabel bantuan yang merujuk ke value
	var ok bool //vaiabel bantuan untuk pengecekan true/false (boolean)
	//mencari key (cariNama) didalam map nilai dan mengambil valuenya (disimpan ke isiData) dan menandakan data sudah ditemukan (TRUE ke ok)
	isiData, ok = nilai[cariNama] 
	//pengecekan apakah data sudah ditemukan
	if ok { //jika ok bernilai true, maka tampilkan hasil searching nya (nilai cariNama = isiData)
		fmt.Println("Nilai", cariNama, " = ", isiData)
	} else { //jika ok bernilai false, maka tampilkan teks data tidak ditemukan
		fmt.Println("Data tidak ditemukan")
	}

	//operasi penghapusan data/value dengan ke Dhimas didalam map nilai
	delete(nilai, "Dhimas")
}