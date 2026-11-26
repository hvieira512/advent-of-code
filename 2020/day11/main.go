package main

import (
	"fmt"
	"reflect"

	"github.com/hvieira512/advent-of-code-2020/utils"
)

type Grid [][]byte

const (
	Floor    byte = '.'
	Empty    byte = 'L'
	Occupied byte = '#'
)

var directions = [][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func main() {
	grid := parse(utils.ReadLines("example.txt"))

	fmt.Println("Part 1:", part1(grid))
}

func parse(data []string) Grid {
	grid := make(Grid, len(data))

	for i, line := range data {
		grid[i] = []byte(line)
	}

	return grid
}

func part1(grid Grid) int {
	for {
		next := grid.run()

		if reflect.DeepEqual(grid, next) {
			return next.countOccupied()
		}

		grid = next
	}
}

func (g Grid) run() Grid {
	next := make(Grid, len(g))

	for row := range g {
		next[row] = make([]byte, len(g[row]))

		for col, seat := range g[row] {
			occupied := g.countAdjacentOccupied(row, col)

			switch seat {
			case Empty:
				if occupied == 0 {
					seat = Occupied
				}
			case Occupied:
				if occupied >= 4 {
					seat = Empty
				}
			}

			next[row][col] = seat
		}
	}

	return next
}

func (g Grid) countAdjacentOccupied(row, col int) int {
	count := 0

	for _, dir := range directions {
		r, c := row+dir[0], col+dir[1]

		if r < 0 || r >= len(g) || c < 0 || c >= len(g[r]) {
			continue
		}

		if g[r][c] == Occupied {
			count++
		}
	}

	return count
}

func (g Grid) countOccupied() int {
	count := 0

	for _, row := range g {
		for _, seat := range row {
			if seat == Occupied {
				count++
			}
		}
	}

	return count
}
