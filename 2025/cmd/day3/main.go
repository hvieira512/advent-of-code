package main

import (
	"fmt"
	"strings"

	u "github.com/hvieira512/aoc2025/cmd/utils"
)


func getBankJoltageK(bank string, K int) int64 {
	n := len(bank)

	digits := []byte(bank)
	stack := make([]byte, 0, K)

	for i := range n {
		for len(stack) > 0 && stack[len(stack)-1] < digits[i] && len(stack)-1+(n-i) >= K {
			stack = stack[:len(stack)-1]
		}

		if len(stack) < K {
			stack = append(stack, digits[i])
		}
	}

	var result int64 = 0
	for i := range K {
		result = result*10 + int64(stack[i]-'0')
	}

	return result
}


func solve(data []string, getBankVoltage func(string) int64) int64 {
	var total int64 = 0

	for _, bank := range data {
		bank = strings.TrimSpace(bank)
		total += getBankVoltage(bank)
	}

	return total
}


func main() {
	data, _ := u.Strings("cmd/day3/input.txt")

	fmt.Printf(
		"Part 1: %v\n",
		solve(data, func(bank string) int64 {
			return getBankJoltageK(bank, 2)
		}),
	)

	fmt.Printf(
		"Part 2: %v\n",
		solve(data, func(bank string) int64 {
			return getBankJoltageK(bank, 12)
		}),
	)
}
