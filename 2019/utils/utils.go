package utils

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadFile(input string) (string, error) {
	file, err := os.ReadFile(input)
	if err != nil {
		return "", err
	}

	return string(file), nil
}

func ReadCSVFile(input string) ([]string, error) {
	file, err := os.Open(input)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println("Error reading records")
	}

	return records[0], err
}
