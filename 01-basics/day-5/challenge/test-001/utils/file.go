package utils

import "os"

func SaveToFile(fileName string, data string) error {
	return os.WriteFile(fileName, []byte(data), 0644)
}

func ReadFromFile(fileName string) (string, error) {
	data, err := os.ReadFile(fileName)
	return string(data), err
}
