package test

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
8️⃣ Tantangan Hari 5
1️⃣ Buat fungsi untuk menghitung luas persegi panjang.

Input: panjang, lebar
Output: luas
2️⃣ Buat fungsi yang mengecek apakah sebuah angka adalah bilangan prima.

Input: angka
Output: true jika prima, false jika tidak
3️⃣ Pisahkan fungsi-fungsi ke dalam package terpisah dan gunakan di main.go.

Setelah selesai, push ke GitHub seperti kemarin
*/

func MenuUtama() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("------- TEST CODING -------")
	fmt.Println()

	fmt.Println("================================")
	fmt.Println("1. Test Coding 1 ")
	fmt.Println("2. Test Coding 2 ")
	fmt.Println("x. Keluar ")
	fmt.Print("Pilih Menu Test: ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	switch input {
	case "1":
		fmt.Println("Menampilkan data........")
		CalculateRectangle()
	case "2":
		fmt.Println("Menampilkan data........")
		CheckPrimeNumber()
	case "x":
		break
	default:
		fmt.Println("Pilihan tidak ada di menu.")
	}

}

func CalculateRectangle() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("------- Menghitung Luas Persegi Panjang -------")
	fmt.Println()

	fmt.Print("Masukkan nilai panjang (atau 'Keluar' untuk kembali ke menu utama): ")
	panjangStr, _ := reader.ReadString('\n')
	panjangStr = strings.TrimSpace(panjangStr)

	if panjangStr == "Keluar" {
		fmt.Println("Kembali ke menu utama.")
		MenuUtama()
		return
	}

	panjang, err := strconv.ParseFloat(panjangStr, 64)
	if err != nil {
		fmt.Println("Input panjang tidak valid, masukkan angka!")
		return
	}
	fmt.Println("===============================================")

	fmt.Print("Masukkan nilai lebar (atau 'Keluar' untuk kembali ke menu utama): ")
	lebarStr, _ := reader.ReadString('\n')
	lebarStr = strings.TrimSpace(lebarStr)

	if lebarStr == "Keluar" {
		fmt.Println("Kembali ke menu utama.")
		return
	}

	lebar, err := strconv.ParseFloat(lebarStr, 64)
	if err != nil {
		fmt.Println("Input lebar tidak valid, masukkan angka!")
		return
	}

	fmt.Println("==============================================")
	fmt.Println("Hasil Luas Persegi Panjang")
	fmt.Println("==============================================")
	result := panjang * lebar

	fmt.Println("Hasil: ", result)

}

func CheckPrimeNumber() bool {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("------- Mengecek Nilai Prima -------")
	fmt.Println()

	fmt.Print("Masukkan angka (atau 'Keluar' untuk kembali ke menu utama): ")
	numStr, _ := reader.ReadString('\n')
	numStr = strings.TrimSpace(numStr)

	if numStr == "Keluar" {
		fmt.Println("Kembali ke menu utama.")
		return true
	}

	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Input panjang tidak valid, masukkan angka!")
		return false
	}

	if num < 2 {
		fmt.Println(num, "bukan bilangan prima.")
		return false
	}
	for i := 2; i <= num; i++ {
		if num%i == 0 {
			fmt.Println(num, "bukan bilangan prima.")
			return false
		}
	}
	fmt.Println(num, "adalah bilangan prima.")
	return true
}
