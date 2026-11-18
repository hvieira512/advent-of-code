package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/hvieira512/advent-of-code-2019/utils"
)

const (
	OpAdd  = 1
	OpMul  = 2
	OpHalt = 99
	Step   = 4
)

func parse(row string) []int {
	data := strings.Split(row, ",")
	numbers := make([]int, 0, len(data))

	for _, aux := range data {
		value, _ := strconv.Atoi(aux)
		numbers = append(numbers, value)
	}

	return numbers
}

func run(data []int) int {
	i := 0
	for i < len(data) {
		if data[i] == OpHalt {
			return data[0]
		}
		a, b, c := data[i+1], data[i+2], data[i+3]
		switch data[i] {
		case OpAdd:
			data[c] = data[a] + data[b]
		case OpMul:
			data[c] = data[a] * data[b]
		}
		i += Step
	}

	return data[0]
}

func part1(data []int) int {
	mem := slices.Clone(data)
	mem[1], mem[2] = 12, 2
	return run(mem)
}

func part2(data []int) int {
	for noun := range 100 {
		for verb := range 100 {
			mem := slices.Clone(data)
			mem[1], mem[2] = noun, verb
			if run(mem) == 19690720 {
				return 100*noun + verb
			}
		}
	}
	panic("no noun/verb pair produces 19690720")
}

func main() {
	row := utils.ReadLines("input.txt")[0]
	data := parse(row)

	fmt.Println("Part 1: ", part1(data))
	fmt.Println("Part 2: ", part2(data))
}
