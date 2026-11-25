package main

import (
	"fmt"
	"strings"

	"github.com/hvieira512/advent-of-code-2020/utils"
)

func part1(data []string) int {
	valid := 0

	required := []string{
		"byr:", "iyr:", "eyr:", "hgt:",
		"hcl:", "ecl:", "pid:",
	}

	for _, passport := range data {
		isValid := true

		for _, field := range required {
			if !strings.Contains(passport, field) {
				isValid = false
				break
			}
		}

		if isValid {
			valid++
		}
	}

	return valid
}
func main() {
	data := utils.ReadLinesBlocks("input.txt")

	fmt.Println("Part 1:", part1(data))
}
