package main

import "fmt"

/**
Operator	Fungsi	Contoh
&&			AND		(x > 10) && (y < 5)
||			OR		(x < 10) && (y > 5)
!			NOT		!(x > 10)
*/

func main() {
	var (
		x, y int = 10, 25
	)

	fmt.Println("AND: ", (x > 10) && (y < 5))
	fmt.Println("OR : ", (x < 10) || (y > 5))
	fmt.Println("NOT: ", !(x > 10))
}
