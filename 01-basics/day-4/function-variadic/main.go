package main

import (
	"fmt"
	"math"
)

/**
Fungsi Variadic (Paramater Dinamis)
-> Fungsi variadic bisa menerima jumlah paramater yang tidak tetap
Fungsi variadic adalah fungsi yang bisa menerima jumlah argumen yang tidak tetap.
Dalam Go, kita bisa menggunakan ... sebelum tipe data parameter untuk mendeklarasikan fungsi variadic.

Bagaimana Fungsi Variadic Bekerja?
Ketika kita mendeklarasikan parameter dengan ...,
maka parameter ini dianggap sebagai slice dari tipe data tertentu ([]float64 dalam kasus ini).
Ini memungkinkan kita untuk memasukkan sejumlah argumen tanpa batas.
*/

// Gunakan ... sebelum tipe data untuk menerima banyak input

func hitung(operator string, numbers ...float64) float64 {
	if len(numbers) == 0 {
		fmt.Println("Error: Tidak ada angka yang diberikan.")
		return 0.0
	}

	result := numbers[0] // sebagai angka pertama atau nilai awal

	for _, num := range numbers[1:] { // loop mulai dari elemen kedua
		switch operator {
		case "+":
			result += num
		case "-":
			result -= num
		case "*":
			result *= num
		case "/":
			if num != 0.0 {
				result /= num
			} else {
				fmt.Println("Error: Pembagian dengan nol dilarang!.")
				return 0.0
			}
		case "%":
			if num != 0.0 {
				result = math.Mod(result, num)
			} else {
				fmt.Println("Error: Modulus dengan nol dilarang!.")
				return 0.0
			}
		default:
			fmt.Printf("Error: %v operator ini tidak ditemukan.\n", operator)
			return 0.0
		}
	}
	return result
}

func main() {
	var operator string
	var num1 float64
	var num2 float64

	fmt.Print("Masukkan angka pertama: ")
	fmt.Scanln(&num1)

	fmt.Println("------------------------")

	fmt.Println("Masukan operator (+,-,*,/,%)")
	fmt.Scanln(&operator)

	fmt.Println("------------------------")

	fmt.Print("Masukkan angka kedua: ")
	fmt.Scanln(&num2)

	fmt.Println("------------------------")

	fmt.Println("========================")

	result := hitung(operator, num1, num2)

	fmt.Printf("%2.f %s %2.f = %2.f\n", num1, operator, num2, result)
	fmt.Println("========================")

	/**
		Gunakan float64 untuk angka karena math.Mod hanya bekerja dengan float64.
	Pastikan ada return jika input kosong agar program tidak error.
	Gunakan result := hitung(operator, num1, num2) tanpa numbers..., karena nilai dikirim langsung
	*/
}
