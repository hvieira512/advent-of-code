package main

import (
	"fmt"
	"strconv"
	"strings"

	u "github.com/hvieira512/aoc2025/cmd/utils"
)

func invalid(s string) bool {
	n := len(s)
	if n%2 != 0 {
		return false
	}

	half := n / 2
	return s[:half] == s[half:]
}

func secondInvalid(s string) bool {
    n := len(s)

    for size := 1; size <= n/2; size++ {
        if n%size != 0 {
            continue
        }

        block := s[:size]
        valid := true

        for i := 0; i < n; i += size {
            if s[i:i+size] != block {
                valid = false
                break
            }
        }

        if valid {
            return true
        }
    }

    return false
}


func solve(data []string, isInvalid func(string) bool) int {
	total := 0

	for _, row := range data {
		row = strings.TrimSpace(row)
		ranges := strings.SplitSeq(row, ",")

		for part := range ranges {
			part = strings.TrimSpace(part)
			ids := strings.Split(part, "-")

			firstId, _ := strconv.Atoi(ids[0])
			lastId, _ := strconv.Atoi(ids[1])

			for i := firstId; i <= lastId; i++ {
				if isInvalid(strconv.Itoa(i)) {
					total += i
				}
			}
		}
	}

	return total
}



func main() {
	data, _ := u.Strings("cmd/day2/input.txt")
	fmt.Printf("Part 1: %v\n", solve(data, invalid))
	fmt.Printf("Part 2: %v\n", solve(data, secondInvalid))
}
