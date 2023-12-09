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

	for i, ch := range line {
		var val = -1
		if unicode.IsDigit(ch) {
			val = int(ch) - '0';
		}

		if (line[i:min(i+3, len(line))] == "one") {
			val = 1
		}
		if (line[i:min(i+3, len(line))] == "two") {
			val = 2
		}
		if (line[i:min(i+5, len(line))] == "three") {
			val = 3
		}
		if (line[i:min(i+4, len(line))] == "four") {
			val = 4
		}
		if (line[i:min(i+4, len(line))] == "five") {
			val = 5
		}
		if (line[i:min(i+3, len(line))] == "six") {
			val = 6
		}
		if (line[i:min(i+5, len(line))] == "seven") {
			val = 7
		}
		if (line[i:min(i+5, len(line))] == "eight") {
			val = 8
		}
		if (line[i:min(i+4, len(line))] == "nine") {
			val = 9
		}
		
		if val != -1 && first_digit == -1 {
			first_digit = val
		}
		if val != -1 {
			last_digit = val
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

/* Go First Impressions: 
 * 
 * Kinda annoying, to be honest. Spent too much time fiddling with the package
 * stuff, and now I have these extra layers of nesting to contend with, and so on.
 *
 * I was able to figure out a simple generic really quickly, so thats nice. We'll
 * see to what extent I will use generics in this project.
 *
 * Split did not behave in the expected way, I thought it would automatically remove
 * the last newline. To be honest, the way go does it makes sense, I just had to be
 * ready to add a Trim.
 *
 * I don't like how all of the string methods are in strings instead of being actual
 * methods on string types. I also got thrown off by the fact that go uses runes instead
 * of ascii characters, even though they did seem to coincide at the end of the day.
 * Actually I think this idea is admirable, it just slowed me down today.
 *
 * Every range for loop being an enumerate is certainly a choice, though its possible
 * Zig does that too? Hard to say.
 *
 * I was upset to learn that you couldn't overslice like in Python (and Rust?), 
 * you have to use min(...) to make sure you don't go out of bounds.
 *
 * I am weary of optional (discouraged, even) semicolons due to Javascript, but I
 * suspect go actually got it right, and I think I like how it looks. Though I tend
 * to think such efforts are a waste of time.
 *
 * The string() and int() half method half type things remind me of Python. For sure
 * pythons influence is strongly felt here. 
 *
 * Finally I'll note that I kinda wrote crappy code here with the 9 seperate checks,
 * but it was what I could get done fastest. Probably a map or a list is the nicer
 * approach, but I would need to learn them really quickly. I imagine I will want
 * to learn them soon. 
 *
 * Slices, other than the index thing, are nice, and slicing slices and comparing
 * to strings worked seamlessly. I know in one of {Rust, Zig} I tried to do it and
 * it worked seamedly, so I am glad it is fine here. And C++ forces you to do substr
 * calls I imagine.
 *
 * Not having to worry about memory or borrow checking is nice too, I guess. */