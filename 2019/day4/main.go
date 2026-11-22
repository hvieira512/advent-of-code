package main

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/hvieira512/advent-of-code-2019/utils"
)

// runs returns the length of each group of repeated digits,
// or nil if the digits ever decrease.
func runs(num int) []int {
	s := strconv.Itoa(num)
	var out []int
	for i := range len(s) {
		switch {
		case i > 0 && s[i] < s[i-1]:
			return nil
		case i > 0 && s[i] == s[i-1]:
			out[len(out)-1]++
		default:
			out = append(out, 1)
		}
	}
	return out
}

func count(ini, fini int, ok func(run int) bool) int {
	n := 0
	for i := ini; i <= fini; i++ {
		if slices.ContainsFunc(runs(i), ok) {
			n++
		}
	}
	return n
}

func main() {
	var ini, fini int
	fmt.Sscanf(utils.ReadLines("input.txt")[0], "%d-%d", &ini, &fini)

	fmt.Println("Part 1:", count(ini, fini, func(r int) bool { return r >= 2 }))
	fmt.Println("Part 2:", count(ini, fini, func(r int) bool { return r == 2 }))
}
