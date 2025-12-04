package main

import (
	"fmt"

	u "github.com/hvieira512/aoc2025/cmd/utils"
)

func isPaperRollValid(data []string, x, y int) bool {
	adjacent := 0

	dirs := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1},           {0, 1},
		{1, -1},  {1, 0},  {1, 1},
	}

	for _, dir := range dirs {
		nextX, nextY := x+dir[0], y+dir[1]

		if nextX < 0 || nextX >= len(data) || nextY < 0 || nextY >= len(data[nextX]) {
			continue
		}

		if data[nextX][nextY] == '@' {
			adjacent++
		}
	}

	return adjacent < 4
}

func partOne(data []string) int {
	total := 0

	for i := range data {
		for j := range data[i] {
			if data[i][j] != '@' {
				continue
			}

			if isPaperRollValid(data, i, j) {
				total++
			}
		}
	}

	return total
}

func partTwo(data []string) int {
	total := 0

	for {
		removedThisRound := 0

		for i := range data {
			row := []byte(data[i])

			for j := range row {
				if row[j] != '@' {
					continue
				}

				if isPaperRollValid(data, i, j) {
					row[j] = '.'
					total++
					removedThisRound++
				}
			}

			data[i] = string(row)
		}

		if removedThisRound == 0 {
			break
		}
	}

	return total
}


func main () {
	data, _ := u.Strings("cmd/day4/input.txt")
	fmt.Printf("Part 1: %v\n", partOne(data))
	fmt.Printf("Part 2: %v\n", partTwo(data))
}
