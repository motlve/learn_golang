package main

import "fmt"

// Penjelasan for loop dengan range di go
// for .. range digunakan untuk mengiterasi elemen dalam array, slice, atau map
// secara langsung tanpa perlu menggunakan indeks secara manual

// Bagaimana cara kerjanya for .. range ?
// names := []string{"Rehan", "Adam", "Zull"}
// ini adalah slice yang berisi 3 elemen string
// Loop for index, name := range names
// range names -> akan mengiterasi setiap elemen di dalam names
// index -> menyimpan indeks elemen saat ini (0,1,2)
// name -> menyimpan nilai elemen saat ini {"Rehan", "Adam", "Zull"}

func main() {
	names := []string{"Rehan", "Adam", "Zull"} // slice string

	for index, name := range names {
		fmt.Printf("Index: %d, Nama: %s \n", index, name)
	}

	// jika tidak ada index bisa pakae , _
}
