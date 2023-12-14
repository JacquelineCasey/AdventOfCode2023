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

type node struct {
	L string
	R string
}

func main() {
    data := string(check(io.ReadAll(os.Stdin)))

	lines := strings.Split(strings.Trim(data, "\n"), "\n")

	directions := strings.Trim(lines[0], " \n")

	nodes := make(map[string] node)

	for _, str := range lines[2:] {
		nodes[str[0:3]] = node{str[7:10], str[12:15]}
	}

	steps := 0
	location := "AAA"

	for location != "ZZZ" {
		if directions[steps % len(directions)] == 'L' {
			location = nodes[location].L
		} else {
			location = nodes[location].R
		}
		
		steps++
	}

	fmt.Println(steps)
}
