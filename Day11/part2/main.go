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
			for i := 0; i < 999; i++ {
				output = append(output, row)
			}
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

	/* We know the input is less than 1000 by 1000, so we expand by 1000 instead of 
	 * a million to keep sizes managable. Then, whenever we see an x or y distance 
	 * greater than 1000, we know it crosses an expansion row / column, and can update 
	 * the sizes accordingly. */

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
			dist := abs(galaxies[i].row - galaxies[j].row) + abs(galaxies[i].col - galaxies[j].col)
			
			/* True distance */
			dist = (dist % 1000) + (dist / 1000 * 1000000)

			sum += dist
		}
	}

	fmt.Println(sum)
}

/* I love my hacky solution to this. It definitely does not work for bigger inputs,
 * but it doesn't have too, and speed to complete the puzzle is a metric I care about.
 *
 * Basically, instead of expanding by 1000000 we expand by 1000. This number is large
 * enough so that the thousands place in the distance between two galaxies tells you
 * how many expansion lines were crossed, so we can easily compute the actual distance
 * by moving the thousands place to the millions place! 
 *
 * Go was fine. I am befuddled by the lack of abs() for ints, I wonder how many Go 
 * programs have been polluted by people writing their own copy of the function I 
 * wrote, possibly copied in multiple files. I wonder how many people just case to 
 * the float version unnecessarily, or inline the absolute value logic adding another 
 * layer of nesting to their code. I wonder how many people add another dependency. 
 * I wonder if that dependency will ever disapear, like `leftPad()` in javascript.
 * The Go people have a really high bar for things that should be included in the 
 * standard library, thinking that someone will have to maintain those functions until 
 * the end of time. But I want convenience, dammit! The maintanence burden has just been
 * moved to someone else anyhow. */
