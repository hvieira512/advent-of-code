package main

import (
	"fmt"
	"strings"

	"github.com/hvieira512/aoc2024/cmd/utils"
)

const (
	Filename = "input"
)

var (
	cache    = make(map[string]bool)
	patterns = make(map[string]struct{})
	cache2   = make(map[string]int)
	maxlen   int
)

// part two
func possibilities(design string) int {
	if count, found := cache2[design]; found {
		return count
	}

	if design == "" {
		cache2[design] = 1
		return 1
	}

	totalCount := 0
	for i := 1; i <= min(len(design), maxlen); i++ {
		if _, exists := patterns[design[:i]]; exists {
			totalCount += possibilities(design[i:])
		}
	}

	cache2[design] = totalCount
	return totalCount
}

func partTwo(designs []string) int {
	count := 0
	for _, design := range designs {
		count += possibilities(design)
	}
	return count
}

// part one
func canObtain(design string) bool {
	if result, found := cache[design]; found {
		return result
	}

	if design == "" {
		cache[design] = true
		return true
	}

	for i := 1; i <= min(len(design), maxlen); i++ {
		_, exists := patterns[design[:i]]
		if exists && canObtain(design[i:]) {
			cache[design] = true
			return true
		}
	}

	cache[design] = false
	return false
}

func partOne(designs []string) int {
	count := 0
	for _, design := range designs {
		if canObtain(design) {
			count++
		}
	}
	return count
}

func parseInput(lines []string) []string {
	patternsLine := lines[0]
	designs := lines[2:]

	for _, pattern := range strings.Split(patternsLine, ", ") {
		patterns[pattern] = struct{}{}
		if len(pattern) > maxlen {
			maxlen = len(pattern)
		}
	}

	return designs
}

func main() {
	utils.RenderDayHeader(19)
	lines, _ := utils.Strings("cmd/day19/" + Filename + ".txt")
	designs := parseInput(lines)

	fmt.Printf("Part One: %v\n", partOne(designs))
	fmt.Printf("Part Two: %v\n", partTwo(designs))
}

