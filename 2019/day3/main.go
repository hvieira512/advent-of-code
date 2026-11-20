package main

import (
	"fmt"
	"github.com/hvieira512/advent-of-code-2019/utils"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

// walk returns, for every point the wire touches, the steps taken to first reach it.
func walk(row string) map[Point]int {
	directions := map[byte]Point{
		'R': {1, 0},
		'L': {-1, 0},
		'U': {0, 1},
		'D': {0, -1},
	}
	visited := map[Point]int{}
	pos := Point{0, 0}
	total := 0
	instructions := strings.SplitSeq(row, ",")
	for instruction := range instructions {
		d := directions[instruction[0]]
		steps, _ := strconv.Atoi(instruction[1:])
		for range steps {
			pos.X += d.X
			pos.Y += d.Y
			total++
			if _, seen := visited[pos]; !seen {
				visited[pos] = total
			}
		}
	}

	return visited
}

func part1(lines []string) int {
	wire1, wire2 := walk(lines[0]), walk(lines[1])

	best := math.MaxInt
	for p := range wire1 {
		if _, ok := wire2[p]; ok {
			best = min(best, utils.Abs(p.X)+utils.Abs(p.Y))
		}
	}

	return best
}

func part2(lines []string) int {
	wire1, wire2 := walk(lines[0]), walk(lines[1])

	best := math.MaxInt
	for p, steps1 := range wire1 {
		if steps2, ok := wire2[p]; ok {
			best = min(best, steps1+steps2)
		}
	}

	return best
}

func main() {
	data := utils.ReadLines("input.txt")

	fmt.Println("Part 1:", part1(data))
	fmt.Println("Part 2:", part2(data))
}
