package main

import (
	"fmt"
	"strconv"
	"strings"

	u "github.com/hvieira512/aoc2025/cmd/utils"
)

const OPERATORS = 3

func partOne(data []string) int {
	total := 0

	parsed := make([][]string, len(data))
	for i, row := range data {
		parsed[i] = strings.Fields(row)
	}

	numCols := len(parsed[0])

	for col := range numCols {
		operator := parsed[OPERATORS][col]

		var result int
		for row := range OPERATORS {
			value, _ := strconv.Atoi(parsed[row][col])

			if row == 0 {
				result = value
				continue
			}

			switch operator {
			case "*":
				result *= value
			case "+":
				result += value
			}
		}

		total += result
	}

	return total
}

func main() {
	data, _ := u.Strings("cmd/day6/example.txt")
	fmt.Printf("Part 1: %v\n", partOne(data))
}
