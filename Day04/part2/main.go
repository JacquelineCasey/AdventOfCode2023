package main

import (
	"fmt"
	"io"
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

	cards := make([]int, len(lines))

	for i, line := range lines {
		cards[i] += 1  // Original card

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

		for j := i + 1; j <= min(i + matches, len(lines) - 1); j++ {
			cards[j] += cards[i]
		}
	}

	/* No sum function... */
	sum := 0
	for _, val := range cards {
		sum += val
	}

	fmt.Println(sum)

}

/* Not a hard puzzle. I thought you'd have to go backwards, but you don't.
 *
 * Go's `make` function is really weird to me.
 *
 * Go doesn't have a [sum over slice or list] function, you end up writing it yourself. 
 * Go seems to be trying to evoke Python with some of its choices, particularly some 
 * syntax things, but it doesn't have even the most basic convenience functions... 
 * which is fine, I couldn't find sum() in Zig either, but it is becoming clear that
 * Go is focused far more on simplicity than on convenience.
 *
 * All told, I'm lest frustrated than in Day03, but still not all that happy with Go. 
 * I think maybe it will get better as I acclimate more. */