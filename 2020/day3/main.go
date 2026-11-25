package main

import (
	"fmt"

	"github.com/hvieira512/advent-of-code-2020/utils"
)

func countTrees(data []string, right int, down int) int {
	trees := 0
	x := 0

	for y := 0; y < len(data); y += down {
		row := data[y]

		if row[x%len(row)] == '#' {
			trees++
		}

		x += right
	}

	return trees
}

func part2(data []string) int {
	slopes := [][2]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	result := 1

	for _, slope := range slopes {
		trees := countTrees(data, slope[0], slope[1])
		result *= trees
	}

	return result
}

func part1(data []string) int {
	return countTrees(data, 3, 1)
}

func main() {
	data := utils.ReadLines("example.txt")

	fmt.Println("Part 1:", part1(data))
	fmt.Println("Part 2:", part2(data))
}
