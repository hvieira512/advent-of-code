package main

import (
	"fmt"
	"strings"

	"github.com/hvieira512/advent-of-code-2020/utils"
)

func part1(groups []string) int {
	total := 0

	for _, group := range groups {
		answers := map[rune]bool{}
		for person := range strings.FieldsSeq(group) {
			for _, answer := range person {
				answers[answer] = true
			}
		}
		total += len(answers)
	}

	return total
}

func part2(groups []string) int {
	total := 0

	for _, group := range groups {
		answers := map[rune]int{}
		people := strings.Fields(group)

		for _, person := range people {
			for _, answer := range person {
				answers[answer]++
			}
		}

		for _, count := range answers {
			if count == len(people) {
				total++
			}
		}
	}

	return total
}

func main() {
	groups := utils.ReadLinesBlocks("input.txt")

	fmt.Println("Part 1:", part1(groups))
	fmt.Println("Part 2:", part2(groups))
}
