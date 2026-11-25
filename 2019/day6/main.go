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

func pathToCOM(orbits map[string]string, start string) []string {
	path := []string{start}
	for start != "COM" {
		start = orbits[start]
		path = append(path, start)
	}

	return path
}

func part1(orbits map[string]string) int {
	total := 0
	for object := range orbits {
		total += len(pathToCOM(orbits, object)) - 1
	}

	return total
}

func part2(orbits map[string]string) int {
	you := pathToCOM(orbits, orbits["YOU"])
	san := pathToCOM(orbits, orbits["SAN"])

	shared := 0
	for shared < len(you) && shared < len(san) && you[len(you)-1-shared] == san[len(san)-1-shared] {
		shared++
	}

	return len(you) + len(san) - 2*shared
}

func main() {
	data := utils.ReadLines("input.txt")
	orbits := parse(data)

	fmt.Println("Part 1:", part1(orbits))
	fmt.Println("Part 2:", part2(orbits))
}
