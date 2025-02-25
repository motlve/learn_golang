package main

import (
	"fmt"
	"math"
)

/**
Latihan: Kalkulator Sederhana
Sekarang, buat program kalkulator sederhana yang bisa melakukan penjumlahan, pengurangan, perkalian, dan pembagian berdasarkan input dari pengguna.
*/

func main() {
	var (
		a, b float64
	)

	var operator string

	fmt.Print("Masukkan angka pertama: ")
	fmt.Scan(&a)

	fmt.Println("Masukkan operator (+,-,*,/,%)")
	fmt.Scan(&operator)

	fmt.Print("Masukkan angka kedua: ")
	fmt.Scan(&b)

	switch operator {
	case "+":
		fmt.Printf("%.2f + %.2f adalah %.2f\n", a, b, (a + b))
	case "-":
		fmt.Printf("%.2f - %.2f adalah %.2f\n", a, b, (a - b))
	case "*":
		fmt.Printf("%.2f * %.2f adalah %.2f\n", a, b, (a * b))
	case "/":
		fmt.Printf("%.2f / %.2f adalah %.2f\n", a, b, (a / b))
	case "%":
		fmt.Printf("%.2f %% %.2f adalah %.2f\n", a, b, math.Mod(a, b))

		//math.Mod(a, b) = Mengembalikan sisa pembagian a / b untuk float64.
		//fmt.Printf, tidak bisa langsung ditulis begitu saja. % adalah bagian dari format string di Printf,
		//jadi Anda harus menggunakan %% untuk mencetak tanda persen.
	default:
		fmt.Printf("%s operator tidak diketahui\n", operator)
	}
}
