package main

import (
	"01-basics/01-basics/day-5/challenge/test-002/controllers"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	MenuUtama()
}

func MenuUtama() {
	controller := controllers.MahasiswaController{FileName: "mahasiswa.json"}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("======== Menu Utama ========")
		fmt.Println("1. Tambah Mahasiswa")
		fmt.Println("2. Lihat Semua Mahasiswa")
		fmt.Println("x. Keluar")
		fmt.Print("Pilih Menu: ")
		scanner.Scan()
		menu := scanner.Text()

		switch menu {
		case "1":
			InputAddedMahasiswa()
		case "2":
			controller.Display()
		case "x":
			fmt.Println("Keluar dari program!")
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi!")
		}
	}
}

func InputAddedMahasiswa() {
	controllerAdded := controllers.MahasiswaController{FileName: "mahasiswa.json"}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("======== Menu Tambah Mahasiswa ========")
	fmt.Print("Masukkan Nama: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Println("======================================")
	fmt.Print("Masukkan Umur: ")
	scanner.Scan()
	age, _ := strconv.Atoi(scanner.Text())

	fmt.Println("======================================")
	fmt.Print("Masukkan Nilai: ")
	scanner.Scan()
	score, _ := strconv.ParseFloat(scanner.Text(), 64)

	controllerAdded.AddMahasiswa(name, age, score)

}
