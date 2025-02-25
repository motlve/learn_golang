package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	nama := "Rizky Aditiyo"
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan nama: ")

	/**
	⚠️ Catatan: fmt.Scan() hanya membaca satu kata. Jika pengguna memasukkan lebih dari satu kata (misalnya: "Rizky Aditiyo"),
				hanya "Rizky" yang akan tersimpan. Untuk membaca satu baris penuh, lebih baik gunakan fmt.Scanln(&namaInput) atau bufio.Scanner.
	*/
	namaInput, _ := reader.ReadString('\n')  // membaca satu baris input
	namaInput = namaInput[:len(namaInput)-1] // menghapus karakter new line

	if namaInput == nama {
		fmt.Printf("%s terdaftar di data kami.\n", namaInput)
	} else if namaInput != nama {
		fmt.Printf("%s tidak terdaftar di data kami.\n", namaInput)
	}
}
