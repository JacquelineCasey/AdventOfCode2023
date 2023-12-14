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

type memo_entry struct {
	spring_count int
	run_count int
}

func no_dots(str string) bool {
	for _, ch := range str {
		if ch == '.' {
			return false
		}
	}

	return true
}

/* A memoized DP approach. */
func count_arrangments(springs string, runs []int, memo_map map[memo_entry] int) int {
	val, ok := memo_map[memo_entry{len(springs), len(runs)}]
	if ok {
		return val
	}

	val = 0

	if springs == "" {
		if len(runs) == 0 {
			val = 1
		} 
	} else {
		if springs[0] == '.' || springs[0] == '?' {
			val += count_arrangments(springs[1:], runs, memo_map)
		}
		if springs[0] == '#' || springs[0] == '?' {
			if len(runs) > 0 && len(springs) >= runs[0] && no_dots(springs[:runs[0]]) {
				skip_point := runs[0]
				
				/* If there is more string left, we have to skip an additional char
				 * because it has to be a working spring. We also ensure that it could 
				 * be a working spring. */
				if len(springs) > skip_point {
					skip_point++;

					if springs[skip_point - 1] != '#' {
						val += count_arrangments(springs[skip_point:], runs[1:], memo_map)
					}
				} else {
					val += count_arrangments(springs[skip_point:], runs[1:], memo_map)
				}
			}
		}
	}

	memo_map[memo_entry{len(springs), len(runs)}] = val

	return val
}

func duplicate_spring(input string) string {
	output := input
	
	for i := 0; i < 4; i++ {
		output += "?" + input
	}

	return output
}

func main() {
    data := string(check(io.ReadAll(os.Stdin)))

	lines := strings.Split(strings.Trim(data, "\n"), "\n")

	sum := 0
	for _, line := range lines {
		pieces := strings.Split(line, " ")
		springs := pieces[0]
		runs := make([]int, 0)

		for _, str := range strings.Split(pieces[1], ",") {
			runs = append(runs, check(strconv.Atoi(str)))
		}

		springs = duplicate_spring(springs)
		new_runs := make([]int, 0)

		for i := 0; i < 5; i++ {
			new_runs = append(new_runs, runs...)
		}

		runs = new_runs

		memo_map := make(map[memo_entry] int)

		sum += count_arrangments(springs, runs, memo_map)
	}

	fmt.Println(sum)
}

/* Very lovely puzzle. I want to tuck this one away for DP week I think. 
 *
 * There's definitely some trickiness around getting it exactly right. I think it
 * is definitely more forgiving as a programming problem than as a theory problem. 
 *
 * Go was fine here. I do find myself writing a lot of functions that could be handled 
 * with an `all()` or an `any()`, but to be honest I think I do that every year even
 * when those tools are available, so who knows. 
 *
 * Go's ability to reslice slices without fuss is pretty handy here. */