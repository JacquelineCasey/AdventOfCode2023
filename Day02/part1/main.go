package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/* Rust style function to throw away an error, assuming it ran fine. */
func check[T any](val T, e error) T {  // maybe check is a better name than unwrap
    if e != nil {
        panic(e)
    }

	return val
}

func subset_possible(subset string) bool {
	pieces := strings.Split(subset, ",")
	for _, piece := range pieces {
		piece = strings.Trim(piece, " ")

		subpieces := strings.Split(piece, " ")
		
		// Unfortunately switches are not expressions. Go is not very expression oriented,
		// which I think is a shame. That being said, it does keep the language simpler,
		// which might have been a design goal.

		max := 0
		switch subpieces[1] {
		case "red":
			max = 12
		case "green":
			max = 13
		case "blue":
			max = 14
		default:
			panic("Unknown Color")
		}
		
		// The fact that go has both `int` and `int64 / 32 / 16 / ...` in this day
		// and age is viscerally upsetting. You have to convert here.

		if int(check(strconv.ParseInt(subpieces[0], 10, 32))) > max {
			return false
		}
	}

	return true
}

func main() {
    data := string(check(io.ReadAll(os.Stdin)))

	lines := strings.Split(strings.Trim(data, "\n"), "\n")
	_ = lines  /* Hmm, I had an unused variable error, so I did this to supress, and now 
	            * I have this leftover because I forgot to clean up. */

	id_sum := 0
	for i, line := range lines {
		id := i + 1
		line, _ = strings.CutPrefix(line, "Game " + fmt.Sprint(id) + ": ")
		substrs := strings.Split(line, ";");
		
		possible := true
		for _, substr := range substrs {
			substr = strings.Trim(substr, " ");

			// I'd love break with value here instead of a flag.
			possible = possible && subset_possible(substr)			
		}

		if possible {
			id_sum += id
		}
	}

	fmt.Println(id_sum);
}
