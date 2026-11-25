package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/hvieira512/advent-of-code-2020/utils"
)

type Ruleset map[string]map[string]int

var bagPattern = regexp.MustCompile(`(\d+) (\w+ \w+) bags?`)

func main() {
	rules := parse(utils.ReadLines("input.txt"))

	fmt.Println("Part 1:", rules.countContainers("shiny gold"))
	fmt.Println("Part 2:", rules.countBags("shiny gold"))
}

func parse(input []string) Ruleset {
	rules := make(Ruleset)

	for _, line := range input {
		parts := strings.SplitN(line, " bags contain ", 2)
		color := parts[0]
		contents := make(map[string]int)

		for _, match := range bagPattern.FindAllStringSubmatch(parts[1], -1) {
			count, _ := strconv.Atoi(match[1])
			contents[match[2]] = count
		}

		rules[color] = contents
	}

	return rules
}

func (r Ruleset) countContainers(color string) int {
	found := make(map[string]bool)

	var find func(string)
	find = func(target string) {
		for container, contents := range r {
			if contents[target] > 0 && !found[container] {
				found[container] = true
				find(container)
			}
		}
	}

	find(color)
	return len(found)
}

func (r Ruleset) countBags(color string) int {
	cache := make(map[string]int)

	var count func(string) int
	count = func(color string) int {
		if total, found := cache[color]; found {
			return total
		}

		total := 0
		for containedColor, quantity := range r[color] {
			total += quantity * (1 + count(containedColor))
		}

		cache[color] = total
		return total
	}

	return count(color)
}
