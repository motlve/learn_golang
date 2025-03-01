package utils

import (
	"01-basics/01-basics/day-5/challenge/test-002/models"
	"encoding/json"
	"fmt"
	"os"
)

const fileName = "mahasiswa.json"

func SaveToFile(mahasiswaList models.ListMahasiswa) {
	data, err := json.MarshalIndent(mahasiswaList, "", " ")
	if err != nil {
		fmt.Println("Gagal menyimpan data: ", err)
		return
	}

	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		fmt.Println("Gagal menyimpan ke file: ", err)
	}
}

func ReadFromFile() (models.ListMahasiswa, error) {
	var mahasiswaList models.ListMahasiswa

	data, err := os.ReadFile(fileName)
	if err != nil {
		return mahasiswaList, err
	}

	err = json.Unmarshal(data, &mahasiswaList)
	if err != nil {
		fmt.Println("Error membaca JSON: ", err)
	}
	return mahasiswaList, nil
}
