package main

import (
	"fmt"
	"io"
	"math"
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

/* github.com/TheAlgorithms/Go/math/gcd */

// Recursive finds and returns the greatest common divisor of a given integer.
func RecursiveGDC(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return RecursiveGDC(b, a%b)
}

/* No LCM in stdlib :(. Copied form online. */
// Lcm returns the lcm of two numbers using the fact that lcm(a,b) * gcd(a,b) = | a * b |
func Lcm(a, b int64) int64 {
	return int64(math.Abs(float64(a*b)) / float64(RecursiveGDC(a, b)))
}

type node struct {
	L string
	R string
}

type history_entry struct {
	direction_loc int
	location string
}

func main() {
    data := string(check(io.ReadAll(os.Stdin)))

	lines := strings.Split(strings.Trim(data, "\n"), "\n")

	directions := strings.Trim(lines[0], " \n")

	nodes := make(map[string] node)

	locations := make([] string, 0)

	for _, str := range lines[2:] {
		nodes[str[0:3]] = node{str[7:10], str[12:15]}

		if str[2] == 'A' {
			locations = append(locations, str[0:3])
		}
	}

	history := make([] map[history_entry] int, len(locations))

	for i, location := range locations {
		history[i] = make(map[history_entry] int)
		history[i][history_entry{0, location}] = 0
	}

	steps := 0
	
	cycled := make(map[int] bool)

	cycle_start_time := make(map[int] int)
	Z_time := make(map[int] int64)
	cycle_end_time := make(map[int] int)

	for len(cycled) != len(locations) {
		next_locations := make([] string, len(locations))

		for i, location := range locations {
			if directions[steps % len(directions)] == 'L' {
				next_locations[i] = nodes[location].L
			} else {
				next_locations[i] = nodes[location].R
			}

			if next_locations[i][2] == 'Z' {
				Z_time[i] = int64(steps + 1)
			}

			seen_time, saw_history := history[i][history_entry{(steps + 1) % len(directions), next_locations[i]}]
			history[i][history_entry{(steps + 1) % len(directions), next_locations[i]}] = steps + 1

			_, saw_cycle := cycled[i]
			if !saw_cycle && saw_history { 
				cycled[i] = true

				cycle_start_time[i] = seen_time
				cycle_end_time[i] = steps + 1
			}
		}

		locations = next_locations
		steps += 1
	}

	cycle_time := make(map[int] int64)
	for i, end_time := range cycle_end_time {
		start_time := cycle_start_time[i]

		cycle_time[i] = int64(end_time - start_time)
	}

	/* We begin skipping around */
	for len(Z_time) != 1 {
		/* Select the min index */
		min_val := int64(math.MaxInt64)
		min_index := -1
		for i, val := range Z_time { 
			if val < min_val { 
				min_index = i
				min_val = val
			}
		}

		Z_time[min_index] += cycle_time[min_index]

		/* Link! */
		for i, z_time_i := range Z_time {
			for j, z_time_j := range Z_time {
				if (z_time_i == z_time_j && i != j) {
					cycle_time_1 := cycle_time[i]
					cycle_time_2 := cycle_time[j]

					new_cycle_time := Lcm(cycle_time_1, cycle_time_2)

					delete(cycle_time, j)
					delete(Z_time, j)

					cycle_time[i] = new_cycle_time

					// fmt.Println(cycle_time)
				}
			}
		}

		// fmt.Println(Z_time)
	}

	for _, v := range Z_time {
		fmt.Println(v)
	}
}

/* Assumptions - each cycle contains only one Z location. Then we can use Chinese
 * Remainder Theorem. */

/* GCD and LCM not being the standard library really pisses me off. Like, what?
 * How on earth is that not in the math module? 
 *
 * Actually, I'm looking at Rust, and it doesn't have gcd or lcm in the stdlib either, 
 * so maybe this is not *that* crazy. They are not thought of as "batteries included" 
 * languages I guess. Zig has it though, which is surprising, but also Zig has awful
 * package management whereas Rust's is divine. I haven't tried Go's yet, opting instead
 * to just copy in the functions. */