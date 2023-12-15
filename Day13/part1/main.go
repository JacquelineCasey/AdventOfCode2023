package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

/* Rust style function to throw away an error, assuming it ran fine. */
func check[T any](val T, e error) T {
    if e != nil {
        panic(e)
    }

	return val
}

/* A returning star from Day 11 */
func transpose(grid[][]bool) [][]bool {
	output := make([][]bool, len(grid[0]))
	for i := range output {
		output[i] = make([]bool, len(grid))
	}

	for i := range grid {
		for j := range grid[0] {
			output[j][i] = grid[i][j]
		}
	}

	return output
}

/* -1 if no such row exists. */
func find_reflection_row(grid [][]bool) int {
	// The line of reflection is just under row r

	for r := 0; r < len(grid) - 1; r++ {
		reflect := true

		for r1 := 0; r1 <= r; r1++ {
			/* Calculate the reflection of this row. */

			r2 := r - r1 + r + 1

			if r2 >= len(grid) {
				continue
			}

			if !slices.Equal(grid[r1], grid[r2]) {
				reflect = false
				break
			}
		}

		if reflect {
			return r
		}
	}

	return -1
}

func main() {
    data := string(check(io.ReadAll(os.Stdin)))

	chunks := strings.Split(strings.Trim(data, "\n"), "\n\n")

	sum := 0

	for _, chunk := range chunks {
		grid := make([][]bool, 0)

		for _, line := range strings.Split(chunk, "\n") {
			row := make([]bool, 0)
			for _, ch := range line {
				row = append(row, ch == '#')
			} 

			grid = append(grid, row)
		}

		row := find_reflection_row(grid)
		grid = transpose(grid)

		col := find_reflection_row(grid)

		if row == -1 {
			sum += col + 1
		} else {
			sum += (row + 1) * 100
		}
	}

	fmt.Println(sum)
}
