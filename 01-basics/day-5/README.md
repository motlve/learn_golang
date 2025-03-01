Hari 6: 
- Struct
- Method
- JSON

materi:
1. Struct -> Cara membuat tipe data sendiri
2. Method -> Fungsi yang terhubung dengan struct
3. Encoding & Decoding JSON -> konversi data struct ke JSON
4. File Handling -> Menyimpan dan membaca data dari file


untuk root foolder 
hari-6/
│── main.go          # Entry point
│── models/
│   ├── student.go   # Struct Student
│── utils/
│   ├── file.go      # File handling functions
│   ├── json.go      # JSON encode/decode functions
│── controllers/
│   ├── student.go   # Method untuk Student


dibagian ini misal kita punya models struct json studentnya.
ini di bagian controllers method untuk student itu akan error.
karena di go:
- Method harus didefinisikan di packagae yang sama dengan structnya
- Tidak bisa menambahkan method di package berbeda langsung
- Jika ingin metode teteap di package controllers/ **gunakan fungsi biasa**

Apa ada cara lain> 
Ada, ini kita tetap mendefinisikan methid di package controllers tanpa mengubah models.

*Alternatif Solusi*
1. Gunakan fungsi biasa di controllers -> alih alih method, buat fungsi yang menerima models.students sebagai paramater.
   package controllers

import (
    "fmt"
    "models"
)

// ✅ Gunakan fungsi biasa, bukan method
func DisplayStudent(s models.Student) {
    fmt.Printf("Nama: %s, Umur: %d, Nilai: %.2f\n", s.Name, s.Age, s.Score)
}

2. Gunakan Embbeded Struct (*Wrapper*) -> Kita bisa membuat struct baru di controllers yang membungkus models.Students, lalu tambahkan method di struct ini.
    package controllers

import (
    "fmt"
    "models"
)

// ✅ Buat struct baru yang "membungkus" models.Student
type StudentController struct {
    Student models.Student
}

// ✅ Tambahkan method ke struct baru ini
func (sc StudentController) Display() {
    fmt.Printf("Nama: %s, Umur: %d, Nilai: %.2f\n", sc.Student.Name, sc.Student.Age, sc.Student.Score)
}
3. Gunakan *type alias* (Kurang direkomen) -> kita bisa membuat alias models.Students di controllers, tapi ini tidak direkomendasikan karena bisa membuat kode lebih sulit dibaca.
    package controllers

import (
    "fmt"
    "models"
)

// ❌ Buat alias (Kurang direkomendasikan)
type StudentAlias models.Student

// ✅ Tambahkan method ke alias
func (s StudentAlias) Display() {
    fmt.Printf("Nama: %s, Umur: %d, Nilai: %.2f\n", s.Name, s.Age, s.Score)
}
