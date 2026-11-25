package main

import (
	"fmt"

	"github.com/hvieira512/advent-of-code-2020/utils"
)

type Partition struct {
	Low, High int
	Lower     byte
	Upper     byte
}

func binaryPartition(p Partition, instructions string) int {
	for i := 0; i < len(instructions); i++ {
		mid := (p.Low + p.High) / 2

		switch instructions[i] {
		case p.Lower:
			p.High = mid
		case p.Upper:
			p.Low = mid + 1
		}
	}

	return p.Low
}

func part1(data []string) int {
	max := 0

	for _, line := range data {
		row := binaryPartition(Partition{0, 127, 'F', 'B'}, line[:7])
		col := binaryPartition(Partition{0, 7, 'L', 'R'}, line[7:])

		seatID := row*8 + col

		if seatID > max {
			max = seatID
		}
	}

	return max
}

func main() {
	data := utils.ReadLines("input.txt")

	fmt.Println("Part 1:", part1(data))
}
