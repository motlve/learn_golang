package main

import "fmt"

/**
for loop dasar -> digunakan untuk mengulang kode beberapa kali
*/

func main() {
	for i := 1; i >= -10; i-- {
		fmt.Println("Iterasi ke- ", i)
	}

	// Penjelasan:
	// i := 1 -> inisialisasi variabel
	// i <= 5 -> kondisi perulangan
	// i++ -> menambah nilai i setiap iterasi
}
