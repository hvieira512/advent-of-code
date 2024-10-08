package main

import (
	"fmt"

	"github.com/hvieira512/aoc2023/utils"
)

type Race struct {
	timestamp int
	distance  int
}

func getRaces(lines []string) []Race {
	races := []Race{}
	timestamps, _ := utils.GetValuesLine(lines, 0, "Time:")
	distances, _ := utils.GetValuesLine(lines, 1, "Distance:")

	for i := 0; i < len(timestamps); i++ {
		races = append(races, Race{timestamps[i], distances[i]})
	}

	return races
}

func getPartOne(races []Race) int {
	records := []int{}
	for _, race := range races {
		recordBeaten := 0
		for btnPress := 1; btnPress < race.timestamp; btnPress++ {
			speed := btnPress
			timeLeft := race.timestamp - btnPress
			distanceTravelled := timeLeft * speed
			if distanceTravelled > race.distance {
				recordBeaten++
			}
		}
		records = append(records, recordBeaten)
	}

	result := records[0] * records[1] * records[2] * records[3]

	return result
}

func main() {
	lines := utils.GetLinesFile("input.txt")
	races := getRaces(lines)

	fmt.Printf("Part 1: %d\n", getPartOne(races))
	// fmt.Printf("Part 2: %d\n", getPartTwo(races))
}
