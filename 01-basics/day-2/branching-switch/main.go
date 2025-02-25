package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

/**
Pada Go, kita tidak bisa menggunakan ekspresi logika (seperti namaInput == nama) langsung dalam case pada switch.

Namun, kita bisa menggunakan switch dengan ekspresi tunggal

switch namaInput {
switch di sini langsung mengevaluasi nilai namaInput.

case nama:
Jika namaInput sama persis dengan nama, maka akan menjalankan perintah dalam case.

default:
Jika tidak ada yang cocok, maka akan menjalankan default.
*/

func validNama(nama string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z\s]+$`)
	return regex.MatchString(nama)
}

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

	switch {
	case namaInput == nama:
		fmt.Printf("%s terdaftar di data kami.\n", namaInput)
	case !validNama(namaInput):
		// Di Go, kita tidak bisa langsung menggunakan regexp (regular expressions) di switch, karena switch hanya bekerja dengan nilai sederhana (bukan ekspresi logika).
		// Namun, kita bisa menggunakan switch true lalu mengevaluasi kondisi dengan regexp di setiap case.
		fmt.Printf("%s input tidak valid! Nama hanya mengandung huruf dan spasi!\n", namaInput)
	default:
		fmt.Printf("%s tidak terdaftar di data kami.\n", namaInput)
	}
}
