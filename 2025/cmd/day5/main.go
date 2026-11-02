package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	u "github.com/hvieira512/aoc2025/cmd/utils"
)

func getData(data []string) ([][]int, []int) {
	freshRanges := [][]int{}
	ids := []int{}

	isRangeSection := true

	for _, row := range data {
		row = strings.TrimSpace(row)

		if row == "" {
			isRangeSection = false
			continue
		}

		if isRangeSection {
			parts := strings.Split(row, "-")
			if len(parts) != 2 {
				continue
			}

			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])

			freshRanges = append(freshRanges, []int{start, end})
		} else {
			id, _ := strconv.Atoi(row)
			ids = append(ids, id)
		}
	}

	return freshRanges, ids
}

func partOne(data []string) int {
	total := 0

	freshRanges, ids := getData(data)

	for _, id := range ids {
		for _, freshRange := range freshRanges {
			if id >= freshRange[0] && id <= freshRange[1] {
				total++
				break
			}
		}
	}

	return total
}

func partTwo(data []string) int {
	freshRanges, _ := getData(data)

	sort.Slice(freshRanges, func(i, j int) bool {
		return freshRanges[i][0] < freshRanges[j][0]
	})

	total := 0

	currentStart := freshRanges[0][0]
	currentEnd := freshRanges[0][1]

	for i := 1; i < len(freshRanges); i++ {
		start := freshRanges[i][0]
		end := freshRanges[i][1]

		if start <= currentEnd+1 {
			if end > currentEnd {
				currentEnd = end
			}
		} else {
			total += currentEnd - currentStart + 1
			currentStart = start
			currentEnd = end
		}
	}

	total += currentEnd - currentStart + 1

	return total
}

func main() {
	data, _ := u.Strings("cmd/day5/input.txt")
	fmt.Printf("Part 1: %v\n", partOne(data))
	fmt.Printf("Part 2: %v\n", partTwo(data))
}
