package main

import (
	"fmt"
	"strconv"

	"github.com/hvieira512/advent-of-code-2019/utils"
)

func fuel(mass int) int {
	return mass/3 - 2
}

func part1(lines []string) int {
	total := 0

	for _, line := range lines {
		mass, _ := strconv.Atoi(line)
		total += fuel(mass)
	}

	return total
}

func part2(lines []string) int {
	total := 0

	for _, line := range lines {
		mass, _ := strconv.Atoi(line)
		for fuel(mass) > 0 {
			mass = fuel(mass)
			total += mass
		}
	}

	return total
}

func main() {
	lines := utils.ReadLines("input.txt")
	fmt.Println("Part 1: ", part1(lines))
	fmt.Println("Part 2: ", part2(lines))
}
