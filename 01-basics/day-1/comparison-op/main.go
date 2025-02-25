package main

import "fmt"

/**
Operator	Fungsi				Contoh
==		Sama dengan				a == b
!=		Tidak sama dengan		a != b
>		Lebih besar				a > b
<		Lebih kecil				a < b
>=		Lebih besar atau sama	a >= b
<=		Lebih kecil atau sama	a <= b
*/

func main() {
	var (
		valueA, valueB int = 20, 5
	)

	fmt.Println("Apakah value A sama dengan dari value B? ", valueA == valueB)
	fmt.Println("Apakah value A tidak sama besar dari value B? ", valueA != valueB)
	fmt.Println("Apakah value A lebih besar dari value B?", valueA > valueB)
	fmt.Println("Apakah value A lebih kecil dari value B?", valueA < valueB)
	fmt.Println("Apakah value A lebih besar atau sama dari value B?", valueA >= valueB)
	fmt.Println("Apakah value A lebih kecil atau sama dari value B?", valueA <= valueB)
}
