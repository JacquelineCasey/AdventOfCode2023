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

func hash(str string) int {
	hash := 0

	for _, ch := range str {
		hash += int(ch)
		hash *= 17
		hash %= 256
	}

	return hash
}

func main() {
    data := string(check(io.ReadAll(os.Stdin)))
	line := strings.Trim(data, "\n")

	sum := 0

	for _, piece := range strings.Split(line, ",") {
		sum += hash(piece)
	}

	fmt.Println(sum)
}
