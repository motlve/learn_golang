package main

import "fmt"

// for loop dengan kondisi
// kita bisa menggunakan for seperti while di bahasa lain

func main() {
	// perulangan dengan for tanpa ekspresi inisialisasi dan increment dalam headet for loop
	// ini disebut sebagai while-like for loop di dalam go

	// mendeklarasikan dan inisialisasi variabel untuk perhitungan dalam perulangan (nilai awal)
	i := 1

	// perulangan for: ini bentuk lain dari perulangan for yg mirip dengan while di bahasa lain
	for i <= 5 {
		// kondisi i <= 5 akan diperiksa lebih dulu:
		// dimana jika bernilai true. maka blok kode dalam for akan dieksekusi
		// jika bernilai false. maka perulangan berhenti
		fmt.Println("Looping ke- ", i)

		// increment i++ berrati menambahkan nilai i sebanyak 1 setiap iterasi
		// ini memastikan bahwa i akhirnya mencapai 6 yang akan menyebabkan kondisi i <= 5 menjadi false sehingga loop berhenti
		i++

	}

	/**

	Iterasi	Nilai i Sebelum Kondisi	Kondisi i <= 5	Cetak Output	i++ (Nilai i Setelah Iterasi)
	1	1	✅ true	Looping ke- 1	2
	2	2	✅ true	Looping ke- 2	3
	3	3	✅ true	Looping ke- 3	4
	4	4	✅ true	Looping ke- 4	5
	5	5	✅ true	Looping ke- 5	6
	6	6	❌ false	Loop berhenti	-
	*/

}
