package main

import "fmt"

/**
Buat program sederhana yang meminta pengguna memasukkan angka, lalu menampilkan apakah angka tersebut ganjil atau genap.

Petunjuk:

Gunakan modulus % untuk menentukan ganjil/genap
Gunakan if-else untuk percabangan
*/

func main() {
	// membuat variabel angka
	var angka int

	// meminta input dari pengguna
	fmt.Print("Masukkan angka: ")
	fmt.Scanln(&angka)

	// logika if untuk statement ganjil dan genap
	if angka%2 == 0 {
		fmt.Printf("Angka %v yang anda masukan adalah genap.\n", angka)
	} else {
		fmt.Printf("Angka %v yang anda masukan adalah ganjil.\n", angka)
	}
}
