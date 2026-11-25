package utils

import (
	"os"
	"strings"
)

func ReadLines(path string) []string {
	raw, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.TrimSpace(string(raw)), "\n")
}

func ReadLinesBlocks(path string) []string {
	raw, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.TrimSpace(string(raw)), "\n\n")
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
