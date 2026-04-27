package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune


func isiArray(t *tabel, n *int) {
	var c rune
	*n = 0
	for {
		fmt.Scanf("%c", &c)
		
		if c == '.' || *n >= NMAX {
			break
		}
		
		if c != '\n' && c != '\r' {
			t[*n] = c
			*n++
		}
	}
}


func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}


func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		temp := t[i]
		t[i] = t[n-1-i]
		t[n-1-i] = temp
	}
}


func palindrom(t tabel, n int) bool {
	// Buat salinan dari array t
	var reverseTabel tabel = t
	
	
	balikanArray(&reverseTabel, n)

	
	for i := 0; i < n; i++ {
		if t[i] != reverseTabel[i] {
			return false 
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	fmt.Print("Teks: ")
	isiArray(&tab, &m) 

	fmt.Print("Reverse teks: ")
	var tabBalik = tab
	balikanArray(&tabBalik, m) 
	cetakArray(tabBalik, m)    
	fmt.Printf("Palindrom: %v\n", palindrom(tab, m)) 
}