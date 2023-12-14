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

type coord struct {
	row int
	col int
}

func get_adjacent(pos coord) []coord {
	output := make([]coord, 4)

	output[0] = coord{pos.row - 1, pos.col}
	output[1] = coord{pos.row + 1, pos.col}
	output[2] = coord{pos.row, pos.col - 1}
	output[3] = coord{pos.row, pos.col + 1}
	
	return output
}

func dfs(pos coord, seen map[coord] bool, grid [][]int) {
	if seen[pos] {
		return
	}

	seen[pos] = true

	for _, neighbor := range get_adjacent(pos) {
		if neighbor.row < 0 || neighbor.row >= len(grid) || neighbor.col < 0 || neighbor.col >= len(grid[0]) {
			continue
		}

		if grid[neighbor.row][neighbor.col] == 0 {
			dfs(neighbor, seen, grid)
		}
	}
}

func main() {
    data := string(check(io.ReadAll(os.Stdin)))

	lines := strings.Split(strings.Trim(data, "\n"), "\n")

	graph := make(map[coord] []coord)

	var S coord

	for row, line := range lines {
		for col, ch := range line {
			neighbors := make([]coord, 0)

			switch ch {
			case '|':
				neighbors = append(neighbors, coord{row - 1, col})
				neighbors = append(neighbors, coord{row + 1, col})
			case '-':
				neighbors = append(neighbors, coord{row, col - 1})
				neighbors = append(neighbors, coord{row, col + 1})
			case 'L':
				neighbors = append(neighbors, coord{row - 1, col})
				neighbors = append(neighbors, coord{row, col + 1})
			case 'J':
				neighbors = append(neighbors, coord{row - 1, col})
				neighbors = append(neighbors, coord{row, col - 1})
			case '7':
				neighbors = append(neighbors, coord{row, col - 1})
				neighbors = append(neighbors, coord{row + 1, col})
			case 'F':
				neighbors = append(neighbors, coord{row, col + 1})
				neighbors = append(neighbors, coord{row + 1, col})
			case '.':  // pass (no fallthrough)
			case 'S':
				S.row = row
				S.col = col
			default:
				panic("Unreachable on valid input")
			}

			graph[coord{row, col}] = neighbors
		}
	}

	/* We have to link S to its neighbors by looking at all things that touch S. */
	for node, neighbors := range graph {
		for _, neighbor := range neighbors {
			if neighbor == S {
				graph[S] = append(graph[S], node)
			}
		}
	}


	/* And now for a BFS part. */

	visited := make(map[coord] int)
	visited[S] = 0

	boundary := make(map[coord] bool)
	boundary[S] = true

	for len(boundary) != 0 {
		new_boundary := make(map[coord] bool)

		for node := range boundary {
			for _, neighbor := range graph[node] {
				_, neighbor_visited := visited[neighbor]
				if !neighbor_visited {
					visited[neighbor] = visited[node] + 1
					new_boundary[neighbor] = true
				}
			}
		}

		boundary = new_boundary
	}

	/* At this point, visited tells you all the tiles that have a pipe. We'd like
	 * to expand the resolution so that there is always a gap between adjacent but
	 * not connected pipes. Then graph search can fill all the area reachable from
	 * the outside. */

	/* We exapand everything out by a factor of two. The new tiles that correspond 
	 * to old tiles are exactly those that have both coordinates even. */

	width := len(lines[0]) * 2 - 1
	height := len(lines) * 2 - 1

	grid := make([][]int, height)
	for i := range grid {
		grid[i] = make([]int, width)
	}

	for node := range visited {
		grid[node.row * 2][node.col * 2] = 1  /* 0 represents air, 1 represents pipe */

		/* Make sure to draw the pipes correctly */
		for _, neighbor := range graph[node] {
			delta_r := neighbor.row - node.row
			delta_c := neighbor.col - node.col

			grid[node.row * 2 + delta_r][node.col * 2 + delta_c] = 1
		}
	}

	/* Now we do graph search from the outside of the map going in. DFS works well
	 * enough. By inspection we know that the outer loop is all connected, so we 
	 * need only search from the top left. */

	seen := make(map[coord] bool)

	dfs(coord{0, 0}, seen, grid)

	/* Now we simply count the enclosed tiles with even coordinates */
	count := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if i % 2 == 0 && j % 2 == 0 {
				if grid[i][j] != 1 && !seen[coord{i, j}] {
					count++
				}
			}
		}
	}

	fmt.Println(count)

	// /* Print out the map */
	// for r, row := range grid {
	// 	for c, tile := range row {
	// 		if tile == 1 {
	// 			fmt.Print("#")
	// 		} else if seen[coord{r, c}] {
	// 			fmt.Print(" ")
	// 		} else {
	// 			fmt.Print(".")
	// 		}
	// 	}

	// 	fmt.Println()
	// }
}

/* Go was once again unintrusive.
 * Also, this was a great puzzle! I loved this one. You have to do graph search 
 * on two types of graphs! The expanding idea is pretty clever.
 *
 * PS: I wonder if there is some sort of deranged approach involving something like 
 * Green's theorem, which can relate the path of the boundary to the enclosed area. */