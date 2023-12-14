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

type span struct {
	start int
	size int
} 

/* Returns the intersection span, and a list of spans from span1 not included in span2.
 * If there is no intersection, returns {-1, -1} and an empty slice */
func intersect(span1 span, span2 span) (span, []span) {
	end_1 := span1.start + span1.size - 1
	end_2 := span2.start + span2.size - 1

	if span1.start > end_2 || span2.start > end_1 {
		return span{-1, -1}, make([]span, 0)
	}

	intersect_start := max(span1.start, span2.start)
	intersect_end := min(end_1, end_2)

	intersection := span{intersect_start, intersect_end - intersect_start + 1}

	missing := make([]span, 0)

	if (span1.start < intersect_start) {
		missing = append(missing, span{span1.start, intersect_start - span1.start})
	}

	if (end_1 > intersect_end) {
		missing = append(missing, span{intersect_end + 1, end_1 - intersect_end})
	}

	return intersection, missing
}

func main() {
    data := string(check(io.ReadAll(os.Stdin)))

	chunks := strings.Split(strings.Trim(data, "\n"), "\n\n")

	inputs := make([]span, 0)

	_, start_line, _ := strings.Cut(chunks[0], ":")
	start_line = strings.Trim(start_line, " \n")

	strs := strings.Split(start_line, " ")
	for i, str1 := range strs {
		if i % 2 == 0 {
			str1 = strings.Trim(str1, " \n")
			
			if str1 == "" {
				continue;
			}

			str2 := strings.Trim(strs[i+1], " \n")

			if str2 != "" {
				inputs = append(inputs, span{check(strconv.Atoi(str1)), check(strconv.Atoi(str2))})
			}
		}
	}

	chunks = chunks[1:]

	for _, chunk := range chunks {
		outputs := make([]span, 0)

		lines := strings.Split(strings.Trim(chunk, "\n"), "\n")
		lines = lines[1:]

		for _, line := range lines {
			pieces := strings.Split(line, " ")
			dst_start := check(strconv.Atoi(pieces[0]))
			src_start := check(strconv.Atoi(pieces[1]))
			size := check(strconv.Atoi(pieces[2]))

			for i := 0; i < len(inputs); i++ {
				input := inputs[i]
				intersecting_span, missed_spans := intersect(input, span{src_start, size})

				if (intersecting_span.size != -1) {
					inputs[i] = inputs[len(inputs) - 1]
					inputs = inputs[:len(inputs) - 1]
					i -= 1  // We will need to iterate spans[i] again 

					/* Splat works with slice */
					inputs = append(inputs, missed_spans...) 

					intersecting_span.start += dst_start - src_start
					outputs = append(outputs, intersecting_span)
				}
			}
		}

		inputs = append(inputs, outputs...)
	}

	min_val := math.MaxInt32

	for _, span := range inputs {
		min_val = min(span.start, min_val)
	}

	fmt.Println(min_val)
}

/* Kinda a fun puzzle, liked the twist. 
 *
 * At least one win for go here, I like that it has the (...) splay operator, so 
 * concat is expressed as simply appending a splat. While I would have liked concat
 * just as well for that (and honestly maybe operator overloading would be best here,
 * though I understand some langauge designers are perhaps rightfully squeemish),
 * I know I will end up liking splat in some other contexts. 
 *
 * slice = append(slice, val) is still a pretty stupid pattern. Python shortens this
 * to slice.append(val), which I think is ideal, though you can go further in Python
 * to slice += val (which is short, but honestly a little weird). Go's version is
 * _very_ wordy.
 *
 * I am also generally opposed to attaching methods to namespaces instead of objects.
 * I don't understand why I am doing strings.Split(str, ...) instead of str.Split(...).
 * I'm kinda ok with it for some global functions though, in particular I think I prefer
 * len(arr) over arr.len, I think it makes index arithmetic read a bit better. */
