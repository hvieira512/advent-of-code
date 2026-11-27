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
	grid := parse(utils.ReadLines("input.txt"))

	fmt.Println("Part 1:", solve(grid, Grid.countAdjacentOccupied, 4))
	fmt.Println("Part 2:", solve(grid, Grid.countVisibleOccupied, 5))
}

func parse(data []string) Grid {
	grid := make(Grid, len(data))

	for i, line := range data {
		grid[i] = []byte(line)
	}

	return grid
}

func solve(grid Grid, count func(Grid, int, int) int, tolerance int) int {
	for {
		next := grid.run(count, tolerance)

		if reflect.DeepEqual(grid, next) {
			return next.countOccupied()
		}

		grid = next
	}
}

func (g Grid) run(count func(Grid, int, int) int, tolerance int) Grid {
	next := make(Grid, len(g))

	for row := range g {
		next[row] = make([]byte, len(g[row]))

		for col, seat := range g[row] {
			occupied := count(g, row, col)

			switch seat {
			case Empty:
				if occupied == 0 {
					seat = Occupied
				}
			case Occupied:
				if occupied >= tolerance {
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

func (g Grid) countVisibleOccupied(row, col int) int {
	count := 0

	for _, dir := range directions {
		r, c := row+dir[0], col+dir[1]

		for r >= 0 && r < len(g) && c >= 0 && c < len(g[r]) && g[r][c] == Floor {
			r, c = r+dir[0], c+dir[1]
		}

		if r >= 0 && r < len(g) && c >= 0 && c < len(g[r]) && g[r][c] == Occupied {
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
