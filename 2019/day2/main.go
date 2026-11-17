package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hvieira512/advent-of-code-2019/utils"
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

const (
	OpAdd  = 1
	OpMul  = 2
	OpHalt = 99
	Step   = 4
)

func part1(data []int) int {
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

func main() {
	row := utils.ReadLines("input.txt")[0]
	data := parse(row)

	fmt.Println("Part 1: ", part1(data))
}
