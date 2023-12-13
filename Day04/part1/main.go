package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

/* Rust style function to throw away an error, assuming it ran fine.
 * I am kinda shocked this is not in the language. Or maybe it is and I just haven't
 * found it. */
func check[T any](val T, e error) T {
    if e != nil {
        panic(e)
    }

	return val
}

func main() {
    data := string(check(io.ReadAll(os.Stdin)))

	lines := strings.Split(strings.Trim(data, "\n"), "\n")

	sum := 0

	for _, line := range lines {
		line = strings.Trim(line, "\n")
		_, line, _ = strings.Cut(line, ":")

		left, right, _ := strings.Cut(line, "|")

		var left_nums []int;
		for _, str := range strings.Split(strings.Trim(left, " "), " ") {
			if str != "" {
				left_nums = append(left_nums, check(strconv.Atoi(str)))
			}
		}

		var right_nums []int;
		for _, str := range strings.Split(strings.Trim(right, " "), " ") {
			if str != "" {
				right_nums = append(right_nums, check(strconv.Atoi(str)))
			}
		}
		
		matches := 0
		for _, val := range right_nums {
			for _, target := range left_nums {
				if val == target {
					matches += 1
					break
				}
			}
		}

		if matches != 0 {
			/* You know I bet if the language had generics back then, we would not need to cast... */
			sum += int(math.Pow(2.0, float64(matches - 1)))
		}
	}

	fmt.Println(sum)
}
