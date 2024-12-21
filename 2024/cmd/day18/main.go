package main

import (
	"container/list"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/hvieira512/aoc2024/cmd/utils"
)

const (
	GridSize = 70
	Filename = "input"
	Bytes    = 1024
)

func partTwo(grid [][2]int) [2]int {
	isConnected := func(n int) bool {
		binaryGrid := make([][]int, GridSize+1)
		for i := range binaryGrid {
			binaryGrid[i] = make([]int, GridSize+1)
		}

		for i := 0; i < n; i++ {
			c, r := grid[i][0], grid[i][1]
			binaryGrid[r][c] = 1
		}

		q := list.New()
		q.PushBack([2]int{0, 0})
		seen := map[[2]int]bool{{0, 0}: true}

		for q.Len() > 0 {
			elem := q.Front()
			q.Remove(elem)
			curr := elem.Value.([2]int)
			r, c := curr[0], curr[1]

			directions := [][2]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}

			for _, dir := range directions {
				nr, nc := r+dir[0], c+dir[1]

				if nr < 0 || nc < 0 || nr > GridSize || nc > GridSize {
					continue
				}

				if binaryGrid[nr][nc] == 1 || seen[[2]int{nr, nc}] {
					continue
				}

				if nr == GridSize && nc == GridSize {
					return true
				}

				seen[[2]int{nr, nc}] = true
				q.PushBack([2]int{nr, nc})
			}
		}
		return false
	}

	lo, hi := 0, len(grid)-1

	for lo < hi {
		mi := (lo + hi) / 2
		if isConnected(mi + 1) {
			lo = mi + 1
		} else {
			hi = mi
		}
	}

	return [2]int{grid[lo][1], grid[lo][0]}
}

func partOne(grid [][2]int) int {
	grid = grid[:Bytes]

	queue := list.New()
	queue.PushBack([3]int{0, 0, 0})

	visited := map[[2]int]bool{{0, 0}: true}

	for queue.Len() > 0 {
		elem := queue.Front()
		queue.Remove(elem)
		curr := elem.Value.([3]int)
		r, c, d := curr[0], curr[1], curr[2]

		for _, dir := range [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
			nr, nc := r+dir[0], c+dir[1]

			if nr < 0 || nc < 0 || nr > GridSize || nc > GridSize ||
				slices.Contains(grid, [2]int{nr, nc}) ||
				visited[[2]int{nr, nc}] {
				continue
			}

			if nr == GridSize && nc == GridSize {
				return d + 1
			}

			visited[[2]int{nr, nc}] = true
			queue.PushBack([3]int{nr, nc, d + 1})
		}
	}

	return -1
}

func parseInput(lines []string) [][2]int {
	grid := [][2]int{}

	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		grid = append(grid, [2]int{y, x})
	}

	return grid
}

func main() {
	lines, _ := utils.Strings("cmd/day18/" + Filename + ".txt")
	utils.RenderDayHeader(18)
	grid := parseInput(lines)

	fmt.Printf("Part One: %v\n", partOne(grid))
	fmt.Printf("Part Two: %v\n", partTwo(grid))
}
