package main

import (
	"fmt"
	"strconv"

	"github.com/hvieira512/advent-of-code-2020/utils"
)

func parse(lines []string) []int {
	values := []int{}
	for _, line := range lines {
		value, _ := strconv.Atoi(line)
		values = append(values, value)
	}
	return values
}

func part1(data []int) int {

	for i := range data {
		for j := range i {
			a, b := data[i], data[j]
			if a+b == 2020 {
				return a * b
			}
		}
	}

	return 0
}

func part2(data []int) int {
	for i := range data {
		for j := range i {
			for k := range j {
				a, b, c := data[i], data[j], data[k]
				if a+b+c == 2020 {
					return a * b * c
				}
			}
		}
	}

	return 0
}

func main() {
	lines := utils.ReadLines("input.txt")
	data := parse(lines)

	fmt.Println("Part 1: data", part1(data))
	fmt.Println("Part 2: data", part2(data))
}
