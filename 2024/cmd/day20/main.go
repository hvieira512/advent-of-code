package main

import (
	"fmt"

	"github.com/hvieira512/aoc2024/cmd/utils"
)

const Filename = "example"

func partOne(track [][]rune) int {
	rows, cols := len(track), len(track[0])

	getStart := func() (int, int) {
		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				if track[r][c] == 'S' {
					return r, c
				}
			}
		}
		return -1, -1
	}
	r, c := getStart()

	distances := make([][]int, rows)
	for i := range distances {
		distances[i] = make([]int, cols)
		for j := range distances[i] {
			distances[i][j] = -1
		}
	}

	distances[r][c] = 0

	for track[r][c] != 'E' {
		for _, dir := range [][2]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}} {
			nr, nc := r+dir[0], c+dir[1]

			if nr < 0 || nc < 0 || nr >= rows || nc >= cols {
				continue
			}

			if track[nr][nc] == '#' || distances[nr][nc] != -1 {
				continue
			}

			distances[nr][nc] = distances[r][c] + 1
			r, c = nr, nc
			break
		}
	}

	for r := 0; r < len(distances); r++ {
		for c := 0; c < len(distances[0]); c++ {
			if distances[r][c] != -1 {
				fmt.Printf("%v\t", distances[r][c])
			} else {
				fmt.Printf(".\t")
			}
		}
		fmt.Println()
	}

	count := 0
	return count
}

func main() {
	utils.RenderDayHeader(20)
	track, _ := utils.Runes("cmd/day20/" + Filename + ".txt")

	fmt.Printf("Part One: %v\n", partOne(track))
}

