package main

import (
	"fmt"
	"strconv"

	"github.com/hvieira512/advent-of-code-2020/utils"
)

type Instruction struct {
	action byte
	value  int
}

type Coordinate struct {
	x int
	y int
}

var compass = map[byte]Coordinate{
	'N': {0, 1},
	'S': {0, -1},
	'E': {1, 0},
	'W': {-1, 0},
}

func main() {
	data := parse(utils.ReadLines("input.txt"))

	fmt.Println("Part 1:", navigate(data, Coordinate{1, 0}, false))
	fmt.Println("Part 2:", navigate(data, Coordinate{10, 1}, true))
}

func parse(data []string) []Instruction {
	instructions := []Instruction{}

	for _, row := range data {
		value, _ := strconv.Atoi(row[1:])
		instructions = append(instructions, Instruction{row[0], value})
	}

	return instructions
}

func navigate(instructions []Instruction, vector Coordinate, useWaypoint bool) int {
	ship := Coordinate{0, 0}

	target := &ship
	if useWaypoint {
		target = &vector
	}

	for _, ins := range instructions {
		switch ins.action {
		case 'N', 'S', 'E', 'W':
			target.x += compass[ins.action].x * ins.value
			target.y += compass[ins.action].y * ins.value
		case 'L', 'R':
			turns := ins.value / 90
			if ins.action == 'L' {
				turns = 4 - turns
			}
			for range turns {
				vector.x, vector.y = vector.y, -vector.x
			}
		case 'F':
			ship.x += vector.x * ins.value
			ship.y += vector.y * ins.value
		}
	}

	return utils.Abs(ship.x) + utils.Abs(ship.y)
}
