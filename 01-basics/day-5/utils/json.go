package utils

/**
	json.Marshal(student) -> fungsi bawaaan dari package encoding/json
	untuk mengubah struct ke JSON

	🎯 Kesimpulan
✔ json.Marshal() digunakan untuk mengubah struct menjadi JSON
✔ Hasilnya dalam bentuk byte slice, jadi perlu dikonversi ke string
✔ Jika ada error saat encoding, fungsi akan mengembalikannya

Jadi, EncodeJSON ini berguna saat kita ingin mengirim data ke API, menyimpan ke file,
atau mengirim data ke database dalam format JSON.
*/

/**
	json.Unmarshal([]byte(data), &student)
- []byte(data) → Mengubah string JSON menjadi slice of bytes, karena json.Unmarshal() hanya menerima data dalam bentuk byte.
- &student → Menggunakan pointer agar JSON langsung di-decode ke dalam variabel student.
- err → Jika ada kesalahan saat decoding, nilainya akan berisi error.

Kesimpulan
✔ json.Unmarshal() digunakan untuk mengubah JSON menjadi struct
✔ JSON harus diubah ke byte slice sebelum di-decode
✔ Jika ada error saat decoding, fungsi akan mengembalikannya

Fungsi DecodeJSON ini berguna saat kita ingin mengambil data dari API, membaca file JSON,
atau menerima data dari frontend dalam format JSON. 🚀
*/

import (
	"01-basics/01-basics/day-5/models"
	"encoding/json"
)

// Encode Student ke JSON -> mengubah (encode) data struct models.Students menjadi format JSON.
func EncodeJSON(student models.Students) (string, error) {
	jsonData, err := json.Marshal(student) // mengubah struct menjadi JSON encoding
	return string(jsonData), err           // mengembalikan hasil JSON dalam bentuk string
}

// Decode JSON ke Student -> mengubah (decode) data JSON menjadi struct models.Students
func DecodeJSON(data string) (models.Students, error) {
	// kita harus menyiapkan struct untuk menampung hasil decoding
	var student models.Students

	err := json.Unmarshal([]byte(data), &student) // lalu diubah JSON menjadi struct
	return student, err                           // mengembalikan struct hasil decoding
}
