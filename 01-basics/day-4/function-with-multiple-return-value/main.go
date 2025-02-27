package main

import "fmt"

// Go mendukung mengembalikan lebih dari satu nilai dalam satu fungsi

func takeName(selected string) (string, bool) {
	names := []string{"Rizky", "Lia", "Chiara", "Jefri", "Lea", "Kenan"}

	fmt.Println(names)

	for _, name := range names {
		if name == selected {
			return name, true
		}
	}
	return selected, false
}

func main() {
	var input string

	fmt.Print("Masukkan nama: ")
	fmt.Scanln(&input)

	result, found := takeName(input)
	if found {
		fmt.Printf("%s ada di dalam daftar nama.\n", result)
	} else {
		fmt.Printf("%s tidak ada di dalam daftar nama.\n", result)
	}
}
