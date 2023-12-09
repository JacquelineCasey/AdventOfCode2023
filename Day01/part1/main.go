package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

/* Rust style function to throw away an error, assuming it ran fine. */
func unwrap[T any](val T, e error) T {
    if e != nil {
        panic(e)
    }

	return val
}

func get_code(line string) int {
	first_digit := -1;
	last_digit := -1;

	for _, ch := range line {
		if unicode.IsDigit(ch) {
			if first_digit == -1 {
				first_digit = int(ch) - '0'
			}
			last_digit = int(ch) - '0'
		}
	}

	return first_digit * 10 + last_digit
}

func main() {
    data := string(unwrap(io.ReadAll(os.Stdin)))

	lines := strings.Split(strings.Trim(data, "\n"), "\n")
	_ = lines

	code_sum := 0
	for _, line := range lines {
		code_sum += get_code(line)
	}

	fmt.Println(code_sum)
}