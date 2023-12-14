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

func sum_slice(slice []int) int {
	sum := 0
	for _, i := range slice {
		sum += i
	}

	return sum
}

/* Returns a slice showing every configuration of balls in bins */
func balls_in_bins_configurations(balls int, bins int) [][]int {
	/* It is always possible to place 0 balls in any number of bins, even 0. */
	if balls == 0 {
		return [][]int{make([]int, bins)}
	}
	
	/* If there is 1 bin, we must place all balls there. */
	if bins == 1 {
		return [][]int{{balls}}
	}

	/* Otherwise, we get to choose how many balls go in the last bin, and how many 
	 * go in later bins. */

	configs := make([][]int, 0)

	for first_balls := 0; first_balls <= balls; first_balls++ {
		for _, config := range balls_in_bins_configurations(balls - first_balls, bins - 1) {
			config = append(config, first_balls)
			configs = append(configs, config)
		}
	}

	return configs
}

func compatible(report string, test string) bool {
	if len(report) != len(test) {
		panic("Bad Lengths")
	}

	for i := range report {
		if report[i] != '?' {
			if report[i] != test[i] {
				return false
			}
		}
	}

	return true
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

		/* It reduces to enumerating the number of ways to put k balls in n bins.
		 * The balls are workings springs we don't know about (we know there is at
		 * least 1 between each run). The bins are the locations between runs, and
		 * at the beginning and end. */
		balls := len(springs) - sum_slice(runs) - (len(runs) - 1)
		bins := len(runs) + 1

		for _, config := range balls_in_bins_configurations(balls, bins) {
			test_str := ""

			for i := range runs {
				test_str += strings.Repeat(".", config[i])
				if i != 0 {
					test_str += "."  // Known
				}

				test_str += strings.Repeat("#", runs[i])
			}

			test_str += strings.Repeat(".", config[len(config) - 1])

			if compatible(springs, test_str) {
				sum++
			}
		}
	}

	fmt.Println(sum)
}
