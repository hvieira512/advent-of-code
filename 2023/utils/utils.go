package utils

import (
	"log"
	"os"
	"strings"
)

func GetLinesFile(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")
	return lines
}
