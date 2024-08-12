package utils

import "os"

func GetFileContent(input string) (string, error) {
	file, err := os.ReadFile(input)
	if err != nil {
		return "", err
	}

	return string(file), nil
}
