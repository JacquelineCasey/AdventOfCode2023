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

	/* Print the largest value */
	max_val := 0
	for _, val := range visited {
		max_val = max(max_val, val)
	}

	fmt.Println(max_val)
}
