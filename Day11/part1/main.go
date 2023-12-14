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

func all_false(row []bool) bool {
	for _, b := range row {
		if b {
			return false
		}
	}
	return true
}

func expand_rows(grid [][]bool) [][]bool {
	output := make([][]bool, 0)

	for _, row := range grid {
		output = append(output, row)

		if all_false(row) {
			output = append(output, row)
		}
	}

	return output
}

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

type coord struct {
	row int
	col int
}

/* Somehow the stdlib only has this for floats... */
func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func main() {
    data := string(check(io.ReadAll(os.Stdin)))

	lines := strings.Split(strings.Trim(data, "\n"), "\n")

	grid := make([][]bool, len(lines))
	for i, line := range lines {
		grid[i] = make([]bool, len(line))

		for j, ch := range line {
			grid[i][j] = ch == '#'
		}
	}

	grid = expand_rows(grid)
	grid = transpose(grid)
	grid = expand_rows(grid)
	grid = transpose(grid)

	galaxies := make([]coord, 0)

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] {
				galaxies = append(galaxies, coord{i, j})
			}
		}
	}

	sum := 0
	for i := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			sum += abs(galaxies[i].row - galaxies[j].row) + abs(galaxies[i].col - galaxies[j].col)
		}
	}

	fmt.Println(sum)
}
