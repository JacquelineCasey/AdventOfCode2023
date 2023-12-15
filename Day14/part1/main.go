package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

/* Rust style function to throw away an error, assuming it ran fine. */
func check[T any](val T, e error) T {
    if e != nil {
        panic(e)
    }

	return val
}

func tilt(grid [][]rune) {
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == 'O' {
				for r1 := r-1; r1 >= 0; r1-- {
					if grid[r1][c] == '.' {
						grid[r1][c] = 'O'
						grid[r1 + 1][c] = '.'
					} else {
						break
					}
				}
			}
		}
	}
}

func load(grid [][]rune) int {
	val := 0

	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == 'O' {
				val += len(grid) - r
			}
		}
	}

	return val
}

func main() {
    data := string(check(io.ReadAll(os.Stdin)))

	lines := strings.Split(strings.Trim(data, "\n"), "\n")

	grid := make([][]rune, 0)
	for _, line := range lines {
		row := make([]rune, 0)
		for _, ch := range line {
			row = append(row, ch)
		}

		grid = append(grid, row)
	}

	tilt(grid)

	fmt.Println(load(grid))
}
