package main

import "fmt"

/**
Operator	Fungsi	Contoh
+		Penjumlahan	a + b
-		Pengurangan	a - b
*		Perkalian	a * b
/		Pembagian	a / b
%		Modulus		a % b
*/

func main() {
	var (
		value1, value2 int = 20, 5
	)

	fmt.Println("Hasil Penjumlahan : ", value1+value2)
	fmt.Println("Hasil Pengurangan : ", value1-value2)
	fmt.Println("Hasil Perkalian   : ", value1*value2)
	fmt.Println("Hasil Pembagian   : ", value1/value2)
	fmt.Println("Hasil Modulus     : ", value1%value2)
}

// Perubahan kecil untuk test commit ulang
