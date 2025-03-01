package main

import (
	"01-basics/01-basics/day-5/challenge/test-001/models"
	"01-basics/01-basics/day-5/challenge/test-001/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	var mahasiswa models.Students

	fmt.Println("========================")
	fmt.Print("Masukkan Nama: ")
	name, _ := reader.ReadString('\n')
	mahasiswa.Name = strings.TrimSpace(name)

	fmt.Println("========================")
	fmt.Print("Masukkan Umur: ")
	age, _ := reader.ReadString('\n')
	mahasiswa.Age, _ = strconv.Atoi(strings.TrimSpace(age))

	fmt.Println("========================")
	fmt.Print("Masukkan Nilai: ")
	score, _ := reader.ReadString('\n')
	mahasiswa.Score, _ = strconv.ParseFloat(strings.TrimSpace(score), 64)

	jsonData, err := utils.EncodeJSON(mahasiswa)
	if err != nil {
		fmt.Println("Gagal encode JSON: ", err)
		return
	}

	fileName := "mahasiswa.json"
	err = utils.SaveToFile(fileName, jsonData)
	if err != nil {
		fmt.Println("Gagal menyimpan ke file: ", err)
		return
	}

	fmt.Println("Data mahasiswa berhasil disimpan!")

	fmt.Println("\nMembaca ulang data dari file...")
	data, err := utils.ReadFromFile(fileName)
	if err != nil {
		fmt.Println("Gagal membaca file: ", err)
		return
	}

	fmt.Println("Data dari file:\n", data)
}
