package main

import (
	"fmt"
	"strconv"

	"github.com/hvieira512/advent-of-code-2020/utils"
)

func main() {
	data := parse(utils.ReadLines("input.txt"))

	fmt.Println("Part 1:", part1(data))
}

func parse(data []string) []int {
	values := []int{}

	for _, row := range data {
		value, _ := strconv.Atoi(row)
		values = append(values, value)
	}

	return values
}

func solve(data []int, target int) bool {
	for i := range data {
		for j := i + 1; j < len(data); j++ {
			a, b := data[i], data[j]

			if a+b == target {
				return true
			}
		}
	}

	return false
}
func part1(data []int) int {
	const preambleSize = 25

	for i := 0; i+preambleSize < len(data); i++ {
		preamble := data[i : i+preambleSize]
		target := data[i+preambleSize]

		if !solve(preamble, target) {
			return target
		}
	}

	return -1
}
