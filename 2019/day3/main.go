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

func walk(row string) map[Point]bool {
	directions := map[byte]Point{
		'R': {1, 0},
		'L': {-1, 0},
		'U': {0, 1},
		'D': {0, -1},
	}
	visited := map[Point]bool{}
	pos := Point{0, 0}
	instructions := strings.SplitSeq(row, ",")
	for instruction := range instructions {
		d := directions[instruction[0]]
		steps, _ := strconv.Atoi(instruction[1:])
		for range steps {
			pos.X += d.X
			pos.Y += d.Y
			visited[pos] = true
		}
	}

	return visited
}

func part1(lines []string) int {
	wire1, wire2 := walk(lines[0]), walk(lines[1])

	best := math.MaxInt
	for p := range wire1 {
		if wire2[p] {
			best = min(best, utils.Abs(p.X)+utils.Abs(p.Y))
		}
	}

	return best
}

func main() {
	data := utils.ReadLines("input.txt")

	fmt.Println("Part 1:", part1(data))
}
