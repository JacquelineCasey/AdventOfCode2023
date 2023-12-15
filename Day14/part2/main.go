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

// We add a direction parameter (or two)
func tilt(grid [][]rune, dr int, dc int) {
	r := 0
	if dr == 1 {
		r = len(grid) - 1
	}

	for r < len(grid) && r >= 0 {
		c := 0
		if dc == 1 {
			c = len(grid[r]) - 1
		}

		for c < len(grid[r]) && c >= 0 {
			if grid[r][c] == 'O' {
				r1 := r
				c1 := c
				for {
					r1 += dr
					c1 += dc

					if r1 < 0 || r1 >= len(grid) || c1 < 0 || c1 >= len(grid[0]) {
						break
					}

					if grid[r1][c1] != '.' {
						break
					}

					grid[r1][c1] = 'O'
					grid[r1 - dr][c1 - dc] = '.'
				}
			}

			if dc == 1 {
				c--
			} else {
				c++
			}
		}

		if dr == 1 {
			r--
		} else {
			r++
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

func cycle(grid [][]rune) {
	tilt(grid, -1, 0)
	tilt(grid, 0, -1)
	tilt(grid, 1, 0)
	tilt(grid, 0, 1)
}

func stringify(grid [][]rune) string {
	val := ""

	for r := range grid {
		for c := range grid[r] {
			val += string(grid[r][c])
		}

		val += "\n"
	}

	return val
}

func gridify(str string) [][]rune {
	lines := strings.Split(strings.Trim(str, "\n"), "\n")

	grid := make([][]rune, 0)
	for _, line := range lines {
		row := make([]rune, 0)
		for _, ch := range line {
			row = append(row, ch)
		}

		grid = append(grid, row)
	}

	return grid
}

func main() {
    data := string(check(io.ReadAll(os.Stdin)))

	grid := gridify(data)

	/* Tells you after how many cycles the state appeared. */
	seen := make(map[string] int)

	cycle_start := -1
	cycle_len := -1
	cycle_end := -1

	for i := 0; true; i++ {
		cycle(grid)

		str := stringify(grid)
		cycle_start_, ok := seen[str]
		cycle_start = cycle_start_  // Annoying, can't mix = and := above
			
		if ok {
			cycle_end = i + 1
			cycle_len = cycle_end - cycle_start
			
			break
		}

		seen[str] = i + 1			
	}

	displacement := (1000000000 - cycle_start) % cycle_len
	final_time := + cycle_start + displacement
	
	final_grid_str := ""
	for k, v := range seen {
		if v == final_time {
			final_grid_str = k
			break
		}
	}

	fmt.Println(load(gridify(final_grid_str)))
}

/* Go's runes are weird little creatures, and its really annoying that they don't 
 * show up as characters when you print them (you have to cast to string).
 *
 * This is a nice puzzle, but I feel like we already had our cycle analysis puzzle 
 * with the graph theory and the ghosts. Furthermore, I felt like that was a harder
 * puzzle? For a while here I worried about how to store the board states before I
 * settled on a fairly naive feeling string approach that does ultimately work.
 *
 * Rewriting the tilt algorithm was annoying, and I wish I developed a cleaner approach 
 * but ultimately I did not. Still, even two years ago I would have written four
 * copies of this function, and I think this is a lot better.
 *
 * Go's lack of ternary op really hurts there. It could clean things up a lot.
 *
 * There is also a very annoying line where I want to mix := and = destructuring, but
 * you can't. I actually don't know if Rust, C++, Zig let you do this, but python definitely 
 * does since it has no real notion of declaration. Probably all the other langauges
 * have this deficiency, so I can't fault Go for it much, though it feels like := should
 * work maybe. */