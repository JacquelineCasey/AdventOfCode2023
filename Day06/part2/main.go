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

	time_str := ""
	line_1, _ := strings.CutPrefix(lines[0], "Time:")
	line_1 = strings.Trim(line_1, "\n")
	for _, str := range strings.Split(line_1, " ") {
		if str != "" {
			time_str += str
		} 
	}

	time := check(strconv.Atoi(time_str))

	distance_str := ""
	line_2, _ := strings.CutPrefix(lines[1], "Distance:")
	line_2 = strings.Trim(line_2, "\n")
	for _, str := range strings.Split(line_2, " ") {
		if str != "" {
			distance_str += str
		} 
	}

	distance := check(strconv.Atoi(distance_str))

	ways_to_win := 0
	for t := 0; t <= time; t++ {
		if t * (time - t) > distance {
			ways_to_win++
		}
	}

	fmt.Println(ways_to_win)
}

/* Can't accuse Go of being slow. I thought this would at least take a second, but
 * it went pretty fast. Which is kinda a shame, I think the puzzle would definitely 
 * be better if I had to do binary search. */