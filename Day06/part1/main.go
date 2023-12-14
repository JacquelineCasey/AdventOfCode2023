package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/* Rust style function to throw away an error, assuming it ran fine. */
func check[T any](val T, e error) T {
    if e != nil {
        panic(e)
    }

	return val
}

func main() {
    data := string(check(io.ReadAll(os.Stdin)))

	lines := strings.Split(strings.Trim(data, "\n"), "\n")

	times := make([]int, 0)
	line_1, _ := strings.CutPrefix(lines[0], "Time:")
	line_1 = strings.Trim(line_1, "\n")
	for _, str := range strings.Split(line_1, " ") {
		if str != "" {
			times = append(times, check(strconv.Atoi(str)))
		} 
	}

	distances := make([]int, 0)
	line_2, _ := strings.CutPrefix(lines[1], "Distance:")
	line_2 = strings.Trim(line_2, "\n")
	for _, str := range strings.Split(line_2, " ") {
		if str != "" {
			distances = append(distances, check(strconv.Atoi(str)))
		} 
	}

	fmt.Println(times, distances)

	product := 1

	for i, time := range times {
		dist := distances[i]
		ways_to_win := 0

		for t := 0; t <= time; t++ {
			my_dist := t * (time - t)

			if my_dist > dist {
				ways_to_win += 1
			}
		}

		product *= ways_to_win
	}

	fmt.Println(product)
}
