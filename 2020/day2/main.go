package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/hvieira512/advent-of-code-2020/utils"
)

func parse(row string) (int, int, string, string) {
	var re = regexp.MustCompile(`^(\d+)-(\d+)\s+(\w):\s+(\w+)$`)
	matches := re.FindStringSubmatch(row)

	if matches == nil {
		return 0, 0, "", ""
	}

	min, _ := strconv.Atoi(matches[1])
	max, _ := strconv.Atoi(matches[2])
	letter, password := matches[3], matches[4]

	return min, max, letter, password
}

func part1(data []string) int {
	valid := 0

	for _, row := range data {
		min, max, letter, password := parse(row)

		count := strings.Count(password, letter)

		if count >= min && count <= max {
			valid++
		}
	}

	return valid
}

func part2(data []string) int {
	valid := 0

	for _, row := range data {
		min, max, letter, password := parse(row)

		if min > len(password) || max > len(password) {
			continue
		}
		if (string(password[min-1]) == letter) != (string(password[max-1]) == letter) {
			valid++
		}
	}

	return valid
}

func main() {
	data := utils.ReadLines("input.txt")

	fmt.Println("Part 1:", part1(data))
	fmt.Println("Part 2:", part2(data))
}
