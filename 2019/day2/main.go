package main

import (
	"fmt"
	"strconv"

	"github.com/hvieira512/advent-of-code-2019/utils"
)

func getValues(data []string) (values []int) {
	for _, value := range data {
		x, _ := strconv.Atoi(value)
		values = append(values, x)
	}

	return
}

// opcode - either 1, 2, or 99
func getPart1(numbers []int) (result int) {
	numbers[1] = 12
	numbers[2] = 2

	for i := 0; i < len(numbers); i += 4 {
		// opcode to stop the program or trying to access out-of-bounds index
		if numbers[i] == 99 || numbers[i+3] >= len(numbers) {
			break
		}

		switch numbers[i] {
		case 1:
			numbers[numbers[i+3]] = numbers[numbers[i+1]] + numbers[numbers[i+2]]
		case 2:
			numbers[numbers[i+3]] = numbers[numbers[i+1]] * numbers[numbers[i+2]]
		}
	}
	return numbers[0]
}

func main() {
	fmt.Println("---- Day 2 ----")

	data, _ := utils.ReadCSVFile("input.txt")
	numbers := getValues(data)

	resultPart1 := getPart1(numbers)
	fmt.Printf("Result Part 1: %d\n", resultPart1)
}
