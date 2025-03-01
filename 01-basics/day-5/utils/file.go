package utils

import "os"

/*
* simpan data ke file
fileName string -> Nama file tempat penyimpanan data
data string -> Data yang ingin disimpan di file

🎯 Kesimpulan
✔ os.WriteFile() digunakan untuk menyimpan data ke file.
✔ Data harus dikonversi ke []byte sebelum disimpan.
✔ Mode 0644 memastikan file bisa dibaca oleh orang lain tetapi hanya bisa diedit oleh pemiliknya.

Fungsi ini sangat berguna untuk menyimpan hasil JSON, log, atau data lain ke file secara otomatis. 🚀
*/

/*
*

  - Penjelasan Fungsi ReadFromFile di Golang
    Fungsi ini digunakan untuk membaca data dari file dalam bentuk string.

    🎯 Kesimpulan

✔ os.ReadFile() digunakan untuk membaca data dari file.
✔ Data hasil pembacaan dikembalikan sebagai string.
✔ Error harus ditangani untuk menghindari crash jika file tidak ada.

Fungsi ini sangat berguna untuk membaca konfigurasi, log, atau data JSON dari file! 🚀
*/
func SaveToFile(fileName string, data string) error {
	return os.WriteFile(fileName, []byte(data), 0644)
	/**
	- os.WriteFile() -> fungsi bawaan go untuk menulis data ke file
	- fileName -> Nama file tujuan (misalnya "data:json")
	- []byte(data) -> Mengubah string menjadi []byte,
	  karena WriteFile() hanya menerima data dalam bentuk byte
	- 0644 -> Mode permission file:
		+ 6 (110) -> Pemilik file bisa membaca & menulis
		+ 4 (100) -> Grup bisa membaca
		+ 4 (100) -> Pengguna lain bisa mengakses
		! Jika file belum ada, maka WriteFile akan membuatnya.
		  Jika suda ada, isinya akan di timpa!
	*/
}

// Baca data dari file
func ReadFromFile(fileName string) (string, error) {
	data, err := os.ReadFile(fileName)
	return string(data), err
	// fileName -> Nama file yang ingin dibaca.
	// Return Values:
	//	- string -> Data yang dibaca dari file dalam bentuk string
	//	- error -> Error jika terjadi kesalahan saat membaca file
	// os.ReadFile(fileName) -> fungsi bawaan go untuk membaca file
	// Hasil pembacaan disimpan dalam variabel data sebagai []byte
	// Jika terjadi kesalahan (misalnya file tidak ditemukan), err
	// - akan berisi pesan error
	// string(data) -> mengubah byte menjadi string agar lebih mudah digunakan
	// jika pembacaan sukses, fungsi mengembalikan data dalam bentuk string
	// jika ada error, error akan dikembalikan
}
