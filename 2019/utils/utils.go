package utils

import "os"

func ReadFile(input string) (string, error) {
	file, err := os.ReadFile(input)
	if err != nil {
		return "", err
	}

	return string(file), nil
}
