package main

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/hvieira512/advent-of-code-2020/utils"
)

func main() {
	data := parse(utils.ReadLines("input.txt"))

	fmt.Println("Part 1:", part1(data))
}

func parse(data []string) []int {
	values := []int{}
	for _, row := range data {
		value, _ := strconv.Atoi(row)
		values = append(values, value)
	}
	values = append(values, 0, slices.Max(values)+3)
	slices.Sort(values)

	return values
}

func part1(data []int) int {
	differences := map[int]int{}
	i := 0

	for i < len(data)-1 {
		curr := data[i+1] - data[i]
		differences[curr]++
		i++
	}

	return differences[1] * differences[3]
}
