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
func find_reflection_row(grid [][]bool, ignore int) int {
	// The line of reflection is just under row r

	for r := 0; r < len(grid) - 1; r++ {
		if r == ignore {
			continue;
		}

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

/* Returns the summary of the new line. Provide the location of the old line. 
 * Use -1 to represent that it is not that row. */
func find_changed_line(grid [][]bool, row int, col int) int {
	t_grid := transpose(grid)

	for r := range grid {
		for c := range grid[r] {
			grid[r][c] = !grid[r][c]
			t_grid[c][r] = !t_grid[c][r]

			new_row := find_reflection_row(grid, row)
			new_col := find_reflection_row(t_grid, col)

			if new_row != -1 {
				return (new_row + 1) * 100
			}
			if new_col != -1 {
				return new_col + 1
			}

			grid[r][c] = !grid[r][c]
			t_grid[c][r] = !t_grid[c][r]
		}
	}

	panic("Expected a new reflection line to be found")
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

		row := find_reflection_row(grid, -1)
		t_grid := transpose(grid)

		col := find_reflection_row(t_grid, -1)

		sum += find_changed_line(grid, row, col)
	}

	fmt.Println(sum)
}

/* Brute force just kinda works.
 * 
 * This is another time where I used the new slices module in the stdlib. I would be
 * remarkably annoyed if this was not around, and it was not around even 4 months ago. */