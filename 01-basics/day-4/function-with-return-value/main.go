package main

import "fmt"

/**

Fungsi bisa mengembalikan nilai menggunakan return

*/

// Fungsi takeANumber() mengembalikan number yang ditemukan

func takeANumber(selected int) (int, bool) {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	for _, number := range numbers {
		if number == selected {
			return number, true // jika angka ditemukan return angka dan true
		}
	}
	return 0, false // jika angka tidak ditemukan, return 0 dan false
}


func main() {

	var input int

	fmt.Print("Masukkan angka yang ada pilih: ")
	fmt.Scanln(&input)

	// memanggil fungsi takeANumbernya
	result, found := takeANumber(input)

	if found {
		fmt.Println("Angka yang anda pilih adalah: ", result)
	} else {
		fmt.Println("Angka tidak detemukan dalam daftar.")
	}
}

/**
Kesimpulan
1. Menggunakan paramater selected int dalam takeANumber()
	-> Paramater ini menyimpan angka yang dipilih user.
2. Mengecek apakah angka ada dalam slice numbers
	-> Jika ditemukan, dikembalikan angka dan true
	->Jika tidak ditemukan, dikembalikan 0 dan false
3. Di main() kita menangani input denan fr.Scanln(&input)
4. Cek apakah angka ditemukan sebelum mencetak hasilnya
*/
