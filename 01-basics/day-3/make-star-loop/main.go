package main

import "fmt"

/**
Studi Kasus: Membuat Pola Bintang *

⭐ Segitiga Siku-siku

Bagaimana kode ini bekerja?
Loop pertama (for i := 1; i <= 5; i++)

i mengontrol jumlah baris dari pola bintang.
Nilai i mulai dari 1 hingga 5, sehingga akan ada 5 baris yang dicetak.
Loop kedua (for j := 1; j <= i; j++)

j mengontrol jumlah bintang per baris.
Pada setiap iterasi loop luar (i), loop dalam (j) akan berjalan sebanyak nilai i.
Misalnya:
Jika i = 1, maka j hanya berjalan sekali (menampilkan 1 bintang).
Jika i = 2, maka j berjalan dua kali (menampilkan 2 bintang).
Jika i = 3, maka j berjalan tiga kali, dan seterusnya.
fmt.Print("*")

Mencetak bintang tanpa pindah ke baris baru (karena pakai Print, bukan Println).
fmt.Println()

Setelah j selesai, perintah ini memindahkan kursor ke baris baru sebelum kembali ke loop luar (i).

*/

func main() {
	for i := 1; i <= 5; i++ { // loop pertama (mengatur jumlah baris)
		for j := 1; j <= i; j++ { // loopkedua (mengatur jumlah bintar  per baris)
			fmt.Print("*") // mencetak bintang tanpa pindah baris
		}
		fmt.Println() // pindahke baris baru setelah inner loop selesai
	}
	/**
		Penjelasan langkah per langkah
	Iterasi i	Iterasi j (jumlah bintang)	Output per baris
	i = 1		j = 1								*
	i = 2		j = 1, 2							**
	i = 3		j = 1, 2, 3							***
	i = 4		j = 1, 2, 3, 4						****
	i = 5		j = 1, 2, 3, 4, 5					*****
	*/
}
