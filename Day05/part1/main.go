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

	chunks := strings.Split(strings.Trim(data, "\n"), "\n\n")

	inputs := make([]int, 0)

	_, start_line, _ := strings.Cut(chunks[0], ":")
	for _, str := range strings.Split(start_line, " ") {
		str = strings.Trim(str, " \n")

		if str != "" {
			inputs = append(inputs, check(strconv.Atoi(str)))
		}
	}

	chunks = chunks[1:]

	for _, chunk := range chunks {
		outputs := make([]int, 0)
		mapped := make([]bool, len(inputs))

		lines := strings.Split(strings.Trim(chunk, "\n"), "\n")
		lines = lines[1:]

		for _, line := range lines {
			
			pieces := strings.Split(line, " ")
			dst_start := check(strconv.Atoi(pieces[0]))
			src_start := check(strconv.Atoi(pieces[1]))
			size := check(strconv.Atoi(pieces[2]))

			for i, input := range inputs {
				if src_start <= input && input < src_start + size {
					outputs = append(outputs, dst_start + input - src_start)

					mapped[i] = true
				} 
			}
		}

		for i, mapped := range mapped {
			if !mapped {
				outputs = append(outputs, inputs[i])
			}
		}
		
		inputs = outputs
	}

	min_val := math.MaxInt32

	for _, val := range inputs {
		min_val = min(val, min_val)
	}

	fmt.Println(min_val)
}
