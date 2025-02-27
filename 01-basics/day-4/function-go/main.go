package main

import "fmt"

/**

Dasar fungsi di go
-> Fungsi adalah blok kode yang bisa dipanggil berkali-kali
   untuk menghindari duplikasi

*/


/**
-> Fungsi sayHello menerima parameter nama bertipe string.
*/
func sayHello(nama string) {
	fmt.Println("Halo,", nama)
}


func main() {
	sayHello("Rizky Aditiyo")
}
