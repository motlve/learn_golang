package utils

import (
	"01-basics/01-basics/day-5/challenge/test-001/models"
	"encoding/json"
)

func EncodeJSON(students models.Students) (string, error) {
	jsonData, err := json.Marshal(students)
	return string(jsonData), err
}

func DecodeJSON(data string) (models.Students, error) {
	var student models.Students
	err := json.Unmarshal([]byte(data), &student)
	return student, err
}
