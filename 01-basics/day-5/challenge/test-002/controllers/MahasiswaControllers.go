package controllers

import (
	"01-basics/01-basics/day-5/challenge/test-002/models"
	"01-basics/01-basics/day-5/challenge/test-002/utils"
	"fmt"
)

type MahasiswaController struct {
	FileName string
}

func (mahasiswaController MahasiswaController) Display() {
	data, err := utils.ReadFromFile()
	if err != nil {
		fmt.Println("Error membaca file: ", err)
		return
	}

	fmt.Println("======== Daftar Mahasiswa ========")
	for i, mahasiswa := range data.Mahasiswa {
		fmt.Printf("%d. %s, Umur: %d, Nilai: %.2f\n", i+1, mahasiswa.Name, mahasiswa.Age, mahasiswa.Score)
	}
}

func (mahasiswaController MahasiswaController) AddMahasiswa(name string, age int, score float64) {
	data, err := utils.ReadFromFile()
	if err != nil {
		fmt.Println("Error membaca file: ", err)
	}

	newMahasiswa := models.DataMahasiswa{
		Name:  name,
		Age:   age,
		Score: score,
	}
	data.Mahasiswa = append(data.Mahasiswa, newMahasiswa)

	utils.SaveToFile(data)
	fmt.Println("Data berhasil disimpan!")
}
