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

/* Both ends are inclusive */
type spanned_num struct {
	num int
	row int
	start_col int 
	end_col int
}

/* Returns gear ratio of a position is a gear (provided it is a star), or returns
 * 0. */
func get_gear_ratio(pos coord, numbers []spanned_num) int {
	first_num := -1
	second_num := -1

	for _, span := range numbers {
		if span.row - 1 <= pos.row && pos.row <= span.row + 1 &&
			span.start_col - 1 <= pos.col && pos.col <= span.end_col + 1 {
			
			if first_num == -1 {
				first_num = span.num
			} else if second_num == -1 {
				second_num = span.num
			} else {
				// 3 nums found. Not a gear.
				return 0
			}
		}
	}

	if second_num != -1 {
		return first_num * second_num
	}

	return 0
}

func main() {
    data := string(check(io.ReadAll(os.Stdin)))

	lines := strings.Split(strings.Trim(data, "\n"), "\n")

	symbols := make(map[coord]rune)
	var numbers []spanned_num

	for row, line := range lines {
		start_col := 0
		num := 0
		for col, ch := range strings.Trim(line, "\n") {
			if ch >= '0' && ch <= '9' {
				num *= 10
				num += int(ch - '0')
			} else {
				if start_col != col {
					/* Why do we need a return here??? I mean, I get why, but does
					 * Go really want us to use a slice instead of a vector or array list? 
					 * Wait, those generic containers don't exist! You have to use slices! */
					numbers = append(numbers, spanned_num{num, row, start_col, col - 1})
					num = 0
				}

				start_col = col + 1

				if ch != '.' {
					symbols[coord{row, col}] = ch
				} 
			}
		}

		if num != 0 {
			numbers = append(numbers, spanned_num{num, row, start_col, len(line) - 1})
		}
	}

	sum := 0
	for pos, sym := range symbols {
		if (sym == '*') {
			sum += get_gear_ratio(pos, numbers)
		}
	}

	fmt.Println(sum);
}

/* Notes:
 * No ternary if in Go.
 *
 * Go enforces the `else if` on the same line thing due to the auto semicolon.
 * - I guess they did screw up auto semicolon. I also have always hated how this looks (cramped).
 *
 * I am finding it really hard to find info on Go features online, e.g. anonymous structs.
 *
 * Seemingly there are not tuples in this language.
 * Or sets. Zig did that too to be fair, but they at least had a void type, so a
 * set was a map from X to void. But the idiom in Go seems to be map from set to 
 * bool, and I can't figure out why you'd want to do that. 
 *
 * It is really upsetting that rune prints out as a number. 
 * 
 * Go had builtin slices and maps before it got generics, which means that you don't 
 * have generic versions of those things, you have to use the slightly awkward versions 
 * instead. Slices are remarkably awkward, if you append to one you have to use the
 * return value because it *might* need to create a new array (why not just use a 
 * layer of indirection? 
 *
 * In short, this puzzle took me from kinda liking Go despite its quirks to really
 * really disliking these weird choices and missing features. The contrast with Zig 
 * is quite surprising, because Zig also has a "less is more" philosophy, yet it feels 
 * like Zig has many more conveniences like tuples and good generics, even though Zig 
 * also forces you to think about manual memory management. */
