package main

import (
	"fmt"
	"strconv"

	"github.com/hvieira512/aoc2023/utils"
)

type Race struct {
	timestamp int
	distance  int
}

func getRaces(lines []string, part int) []Race {
	races := []Race{}

	// group all the numbers into one big one
	timestamps, _ := utils.GetValuesLine(lines, 0, "Time:")
	distances, _ := utils.GetValuesLine(lines, 1, "Distance:")

	switch part {
	case 1:
		for i := 0; i < len(timestamps); i++ {
			races = append(races, Race{timestamps[i], distances[i]})
		}
		break
	case 2:
		newTimestamp := groupNumbers(timestamps)
		newDistance := groupNumbers(distances)
		races = append(races, Race{newTimestamp, newDistance})
		break
	}

	return races
}

func groupNumbers(numbers []int) int {
	newNumStr := ""

	for _, number := range numbers {
		aux := strconv.Itoa(number)
		newNumStr += aux
	}

	newNum, _ := strconv.Atoi(newNumStr)

	return newNum
}

func getResult(races []Race) int {
	records := []int{}
	result := 0

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

	if len(records) > 2 {
		// part one
		result = records[0] * records[1] * records[2] * records[3]
	} else {
		// part two
		result = records[0]
	}

	return result
}

func main() {
	lines := utils.GetLinesFile("input.txt")

	races := getRaces(lines, 1)
	fmt.Printf("Part 1: %d\n", getResult(races))

	newRaces := getRaces(lines, 2)
	fmt.Printf("Part 2: %d\n", getResult(newRaces))
}
