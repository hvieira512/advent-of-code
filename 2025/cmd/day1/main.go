package main

import (
	"fmt"
	"strconv"
	"strings"

	u "github.com/hvieira512/aoc2025/cmd/utils"
)

func partTwo(data []string) int {
	const size = 100

	dial := 50
	zeros := 0

	for _, row := range data {
		if row = strings.TrimSpace(row); row == "" {
			continue
		}

		value, _ := strconv.Atoi(row[1:])
		if row[0] == 'L' {
			value = -value
		}

		fullTurns := abs(value) / size
		zeros += fullTurns

		rem := abs(value) % size
		move := 1
		if value < 0 {
			move = -1
		}

		for range rem {
			dial = (dial + move + size) % size
			if dial == 0 {
				zeros++
			}
		}
	}

	return zeros
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

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
	fmt.Printf("Part 2: %v\n", partTwo(data))
}
