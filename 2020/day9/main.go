package main

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/hvieira512/advent-of-code-2020/utils"
)

const preambleSize = 25

func main() {
	data := parse(utils.ReadLines("input.txt"))

	fmt.Println("Part 1:", part1(data))
	fmt.Println("Part 2:", part2(data))
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
	for i := 0; i+preambleSize < len(data); i++ {
		preamble := data[i : i+preambleSize]
		target := data[i+preambleSize]

		if !solve(preamble, target) {
			return target
		}
	}

	return -1
}

func part2(data []int) int {
	target := part1(data)

	for i := range data {
		sum := data[i]
		for j := i + 1; j < len(data) && sum < target; j++ {
			sum += data[j]
			if sum == target {
				window := data[i : j+1]
				return slices.Min(window) + slices.Max(window)
			}
		}
	}

	return -1
}
