Pertanyaan untuk Hari 3 (Variabel & Tipe Data Lanjutan)
1️⃣ Apa perbedaan antara float64 dan int dalam Go?
2️⃣ Bagaimana cara mengonversi float64 ke int di Go?
3️⃣ Apa yang terjadi jika kita melakukan operasi antara int dan float64 tanpa konversi?
4️⃣ Bagaimana cara membaca input angka dari pengguna dan memastikan itu adalah angka yang valid?
5️⃣ Coba jelaskan bagaimana cara kerja fmt.Scanln() dalam membaca input dari pengguna.

Jawaban:
1. float64 itu adalah tipe data yang bertipe bilangan desimal dengan 64nya itu ada bit dari sebuah komputer, lalu int itu adalah tipe data yang bertipe bilangan bulat.
2. dengan cara type conversion -> misal ada 
    var angka float = 10.75
    var angkaInt int = int(angka) // ini konversi dari float ke int
    jadi ini akan membuat angka 10.75 menjadi 10
3. Ini akan terjadi error invalid konversi pada int ke float64 sebab karena mereka berbeda tipe data
4. menggunakan bufio.Scanner terus dan memastikan itu angka yang valid bisa menggubakan sebuah if else untuk menentukan angka yg valid dan bisa membuat regex untuk inputan tidak boleh inputan kecuali hanya angka
5. fmt.Scanln() -> ini seperti scan tetapi scanln ini tidak membaca dari newline atau garis baru dan membaca dengan akhir dari garis baru

Jawaban yang Diperbaiki & Dilengkapi:
1️⃣ Apa perbedaan antara float64 dan int dalam Go?
✅ float64 adalah tipe data untuk bilangan desimal (floating point) yang memiliki presisi 64-bit.
✅ int adalah tipe data untuk bilangan bulat (integer) tanpa nilai desimal.

Contoh:


var a int = 10
var b float64 = 10.75
a hanya bisa menyimpan angka bulat (tanpa koma).
b bisa menyimpan angka desimal.
2️⃣ Bagaimana cara mengonversi float64 ke int di Go?
✅ Dengan type conversion, yaitu menggunakan int() untuk mengubah float64 menjadi int.

Contoh:


var angka float64 = 10.75
var angkaInt int = int(angka) // Konversi dari float64 ke int
fmt.Println(angkaInt) // Output: 10
👉 Catatan: Konversi ini membuang bagian desimal, bukan membulatkan!

int(10.75) → 10
int(5.99) → 5
int(-2.9) → -2
Kalau ingin membulatkan ke angka terdekat, pakai math.Round():


import "math"

var angka float64 = 10.75
var angkaInt int = int(math.Round(angka))
fmt.Println(angkaInt) // Output: 11
3️⃣ Apa yang terjadi jika kita melakukan operasi antara int dan float64 tanpa konversi?
❌ Error karena Go tidak mengizinkan operasi antara tipe data berbeda tanpa konversi eksplisit.

Contoh yang salah (error):

var a int = 10
var b float64 = 5.5

fmt.Println(a + b) // ❌ ERROR: mismatched types int and float64
Solusi: Konversi salah satu variabel agar sesuai.


fmt.Println(float64(a) + b) // ✅ 15.5
4️⃣ Bagaimana cara membaca input angka dari pengguna dan memastikan itu adalah angka yang valid?
📌 Metode 1: Menggunakan fmt.Scanln()


var angka int
fmt.Print("Masukkan angka: ")
_, err := fmt.Scanln(&angka)

if err != nil {
    fmt.Println("Input tidak valid! Harap masukkan angka.")
} else {
    fmt.Println("Angka yang dimasukkan:", angka)
}
Jika pengguna memasukkan angka, program akan lanjut.
Jika pengguna memasukkan teks, program akan mendeteksi error.
📌 Metode 2: Menggunakan bufio.Scanner (lebih fleksibel)

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Masukkan angka: ")
    input, _ := reader.ReadString('\n')

    input = strings.TrimSpace(input) // Hapus spasi & newline
    angka, err := strconv.Atoi(input) // Konversi ke int

    if err != nil {
        fmt.Println("Input tidak valid! Harap masukkan angka.")
    } else {
        fmt.Println("Angka yang dimasukkan:", angka)
    }
}
Kelebihan metode ini: ✅ Bisa membaca input panjang
✅ Bisa menangani spasi & newline
✅ Bisa mengonversi string ke angka dengan validasi

5️⃣ Coba jelaskan bagaimana cara kerja fmt.Scanln() dalam membaca input dari pengguna.
🔹 fmt.Scanln() digunakan untuk membaca input dari pengguna hingga baris baru (newline).
🔹 Jika ada lebih dari satu nilai yang dimasukkan, Scanln() akan membaca hanya satu kata pertama dan mengabaikan sisanya.
🔹 Contoh penggunaan:


var nama string
fmt.Print("Masukkan nama: ")
fmt.Scanln(&nama)
fmt.Println("Halo", nama)
Jika pengguna mengetik "Aqil Iye", hanya "Aqil" yang terbaca.
Untuk membaca input dengan spasi, lebih baik gunakan bufio.Reader.