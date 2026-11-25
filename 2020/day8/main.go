package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hvieira512/advent-of-code-2020/utils"
)

func main() {
	instructions := parse(utils.ReadLines("input.txt"))

	fmt.Println("Part 1:", part1(instructions))
	fmt.Println("Part 2:", part2(instructions))
}

type Instruction struct {
	Operation string
	Argument  int
}

func parse(data []string) []Instruction {
	instructions := []Instruction{}

	for _, row := range data {
		before, after, _ := strings.Cut(row, " ")
		argument, _ := strconv.Atoi(after)

		instructions = append(instructions, Instruction{
			Operation: before,
			Argument:  argument,
		})
	}

	return instructions
}

func run(instructions []Instruction) (int, bool) {
	accumulator := 0
	visited := make(map[int]bool)

	i := 0

	for i >= 0 && i < len(instructions) && !visited[i] {
		visited[i] = true

		instruction := instructions[i]

		switch instruction.Operation {
		case "acc":
			accumulator += instruction.Argument
			i++
		case "jmp":
			i += instruction.Argument
		case "nop":
			i++
		}
	}

	return accumulator, i == len(instructions)
}

func part1(instructions []Instruction) int {
	accumulator, _ := run(instructions)
	return accumulator
}

func part2(instructions []Instruction) int {
	for i, instruction := range instructions {
		switch instruction.Operation {
		case "jmp":
			instructions[i].Operation = "nop"
		case "nop":
			instructions[i].Operation = "jmp"
		default:
			continue
		}

		accumulator, terminated := run(instructions)

		if terminated {
			return accumulator
		}

		// Restore the original instruction before trying the next one.
		instructions[i] = instruction
	}

	return 0
}
