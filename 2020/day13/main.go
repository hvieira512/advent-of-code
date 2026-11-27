package main

import (
	"fmt"
	"maps"
	"slices"
	"strconv"
	"strings"

	"github.com/hvieira512/advent-of-code-2020/utils"
)

func main() {
	timestamp, buses := parse(utils.ReadLines("input.txt"))

	fmt.Println("Part 1:", part1(timestamp, buses))
	fmt.Println("Part 2:", part2(buses))
}

func parse(data []string) (int, []int) {
	timestamp, _ := strconv.Atoi(data[0])

	busesStr := strings.Split(data[1], ",")

	buses := make([]int, len(busesStr))

	for i, busStr := range busesStr {
		if busStr == "x" {
			buses[i] = -1
			continue
		}

		buses[i], _ = strconv.Atoi(busStr)
	}

	return timestamp, buses
}

func part1(timestamp int, buses []int) int {
	waits := make(map[int]int)

	for _, bus := range buses {
		if bus == -1 {
			continue
		}

		wait := (-timestamp%bus + bus) % bus
		waits[wait] = bus
	}

	minWait := slices.Min(slices.Collect(maps.Keys(waits)))

	return minWait * waits[minWait]
}

func part2(buses []int) int {
	t, step := 0, 1

	for position, bus := range buses {
		if bus == -1 {
			continue
		}

		for (t+position)%bus != 0 {
			t += step
		}

		step *= bus
	}

	return t
}
