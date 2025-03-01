package utils

import (
	"01-basics/01-basics/day-5/challenge/test-002/models"
	"encoding/json"
)

func EncodeJSON(mahasiswa models.DataMahasiswa) (string, error) {
	jsonData, err := json.MarshalIndent(mahasiswa, "", " ")
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func DecodeJSON(dataMahasiswa string) (models.ListMahasiswa, error) {
	var mahasiswaList models.ListMahasiswa
	err := json.Unmarshal([]byte(dataMahasiswa), &mahasiswaList)
	if err != nil {
		return mahasiswaList, err
	}
	return mahasiswaList, nil
}
