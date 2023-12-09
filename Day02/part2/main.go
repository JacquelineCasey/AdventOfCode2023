package main

/* This syntax is distrubing IMO. No commas?! */
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

// We'll use the multiple returns idea
func read_subset(subset string) (red int, green int, blue int) {
	pieces := strings.Split(subset, ",")
	for _, piece := range pieces {
		piece = strings.Trim(piece, " ")

		subpieces := strings.Split(piece, " ")
		
		// Unfortunately switches are not expressions. Go is not very expression oriented,
		// which I think is a shame. That being said, it does keep the language simpler,
		// which might have been a design goal.
		
		val := int(check(strconv.ParseInt(subpieces[0], 10, 32)))

		switch subpieces[1] {
		case "red":
			red = val
		case "green":
			green = val
		case "blue":
			blue = val
		default:
			panic("Unknown Color")
		}
	}

	return /* Returns the variables declared at the top */
}

func main() {
    data := string(check(io.ReadAll(os.Stdin)))

	lines := strings.Split(strings.Trim(data, "\n"), "\n")

	power_sum := 0
	for i, line := range lines {
		id := i + 1
		line, _ = strings.CutPrefix(line, "Game " + fmt.Sprint(id) + ": ")
		substrs := strings.Split(line, ";")
		
		red, green, blue := 0, 0, 0
		for _, substr := range substrs {
			new_red, new_green, new_blue := read_subset(strings.Trim(substr, " "))

			red = max(red, new_red)
			green = max(green, new_green)
			blue = max(blue, new_blue)
		}

		power_sum += red * green * blue
	}

	fmt.Println(power_sum);
}

/* Already got into more of a groove with go. I think I would have dissed go's multiple
 * returns idea (it's just a tuple), but I did try named returns and I think it could
 * make things look a little nicer. It's weird that you still need a naked return at the
 * bottom though.
 *
 * Go's switch is better than C's, but I really wanted it to be an expression and
 * not a statement, so it fails to deliver there.
 *
 * I've am no longer stumbling (much) on := and =, and I think overall I like it.
 * Shorter than `let _ =` in the other languages, but still protects against the
 * pythony / javascripty errors where you assign a value to a typo'd variable and
 * break your code without warning. My only question - does Go care about constness?
 * A google reveals that it has a `const` keyword that is the counterpart of `var`,
 * however, := does var by default, so if you like the short form then you have to
 * actively think about const-ing things, which in some ways is convenient and in
 * other ways is dangerous. I like that Rust is immutable by default, and I like that
 * Zig is even handed at least (though const-ness overall is confusing in Zig, like
 * a const array can still have its values modified).
 *
 * I like that you can just + strings together. It says something about the languages
 * that I code with that I miss that feature (Rust forces you to think about slices,
 * Zig maybe can + things but probably can't). However, you can't + on an int, you
 * have to convert, and you can't just use string() like in python, so it's not perfect. 
 *
 * Strings overall are reasonably convenient. Split is great, and solves 80% of parsing. 
 * I remember Zig had split but it was weird, returning an iterator, so you had to be
 * ready for weirdness. 
 *
 * So yeah, I'm already becoming way more productive in this language. I think this 
 * is the sign of a good learning curve, but also Go probably just doesn't have much 
 * to learn besides syntax and where functions are. Easyness to learn is a very admirable 
 * feature for a language like Go. */