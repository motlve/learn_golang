package main

import "fmt"

/**
	Tugas Hari 4
1️⃣ Buat program yang mencetak bilangan ganjil dari 1 sampai 20.
2️⃣ Buat pola segitiga terbalik dengan * seperti ini:


*****
****
***
**
*
3️⃣ Buat program yang membaca angka dari user, lalu cetak angka tersebut sebanyak jumlah angka yang dimasukkan.

Contoh Input:

Masukkan angka: 4
Output:

Salin
Edit
4
4
4
4
*/

func soal1() {

	for i := 1; i <= 20; i++ {
		if i%2 != 0 {
			fmt.Println("Bilangan ganjil: ", i)
		}
	}
}

func soal2() {
	for i := 5; i > 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}

func soal3() {
	var angka int

	fmt.Print("Masukkan angka: ")
	fmt.Scanln(&angka)

	for i := 1; i <= angka; i++ {
		fmt.Println(angka)
	}
}

func main() {
	soal1()
	soal2()
	soal3()
}
