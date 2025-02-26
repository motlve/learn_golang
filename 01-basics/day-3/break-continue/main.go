package main

import "fmt"

/**
Break dan Continue
break -> menghenttikan perulangan
continue -> melewati iterasi saat ini dan melanjutkan ke berikutnya
*/

/**
i := 1

	for i <= 10 {
		if i == 5 {
			continue // lewati angka 5
		}

		fmt.Println("Angka: ", i)

		i++
	}

Kenapa hasilnya tidak sesuai harapan?
Saat i == 5, program masuk ke blok if dan menjalankan continue, yang langsung melompat ke iterasi berikutnya tanpa menjalankan i++. Akibatnya:

i tetap 5 selamanya.
Program terjebak di if i == 5 terus-menerus (infinite loop).
Output hanya sampai angka 4 karena saat i == 5, kode terus melewati bagian fmt.Println().
*/

func main() {
	i := 1

	for i <= 10 {
		if i == 5 {
			i++
			continue // lewati angka 5
		}

		fmt.Println("Angka: ", i)

		i++
	}

	/**
		Kesimpulan
	continue melewatkan iterasi saat ini, tapi tidak meningkatkan nilai i otomatis.
	Jika continue digunakan dalam loop dengan variabel penghitung, pastikan variabel tersebut tetap bertambah agar loop tidak macet (infinite loop).
	Solusinya tambahkan i++ sebelum continue agar i tetap bertambah! ✅
	*/
}
