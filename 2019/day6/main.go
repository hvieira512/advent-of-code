package main

import (
	"fmt"
	"strings"

	"github.com/hvieira512/advent-of-code-2019/utils"
)

func parse(lines []string) map[string]string {
	orbits := make(map[string]string)

	for _, line := range lines {
		left, right, _ := strings.Cut(line, ")")
		orbits[right] = left
	}

	return orbits
}

func part1(orbits map[string]string) int {
	total := 0

	for k := range orbits {
		current, count := k, 0

		for current != "COM" {
			count++
			current = orbits[current]

		}

		total += count
	}

	return total
}

func main() {
	data := utils.ReadLines("input.txt")
	orbits := parse(data)

	fmt.Println("Part 1:", part1(orbits))
}
