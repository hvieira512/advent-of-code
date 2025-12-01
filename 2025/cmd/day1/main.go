package main

import (
	"fmt"
	"strconv"
	"strings"

	u "github.com/hvieira512/aoc2025/cmd/utils"
)

func partOne(data []string) int {
	dial := 50
	zeros := 0

	for _, row := range data {
		row = strings.TrimSpace(row)
		if row == "" {
			continue
		}

		value, _ := strconv.Atoi(row[1:])

		switch row[0] {
		case 'L':
			dial -= value
		case 'R':
			dial += value
		}

		dial = ((dial % 100) + 100) % 100
		if dial == 0 {
			zeros++
		}
	}

	return zeros
}

func main() {
	data, _ := u.Strings("cmd/day1/input.txt")
	fmt.Printf("Part 1: %v\n", partOne(data))
}
