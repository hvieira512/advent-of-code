package main

import (
	"fmt"
	"strings"

	"github.com/hvieira512/advent-of-code-2020/utils"
)

func solve(groups []string, predicate func(int, int) bool) int {
	total := 0

	for _, group := range groups {
		people := strings.Fields(group)
		answers := make(map[rune]int)

		for _, person := range people {
			for _, answer := range person {
				answers[answer]++
			}
		}

		for _, answerCount := range answers {
			if predicate(answerCount, len(people)) {
				total++
			}
		}
	}

	return total
}
func part1(groups []string) int {
	return solve(groups, func(count, _ int) bool {
		return count > 0
	})
}

func part2(groups []string) int {
	return solve(groups, func(count, people int) bool {
		return count == people
	})
}

func main() {
	groups := utils.ReadLinesBlocks("input.txt")

	fmt.Println("Part 1:", part1(groups))
	fmt.Println("Part 2:", part2(groups))
}
