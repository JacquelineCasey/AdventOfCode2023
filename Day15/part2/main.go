package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
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

type lens struct {
	label string
	focal_length int
}

func main() {
    data := string(check(io.ReadAll(os.Stdin)))
	line := strings.Trim(data, "\n")

	boxes := make([][]lens, 256)
	for i := range boxes {
		boxes[i] = make([]lens, 0)
	}

	for _, instruction := range strings.Split(line, ",") {
		if strings.Contains(instruction, "=") {
			pieces := strings.Split(instruction, "=")

			label := pieces[0]
			value := check(strconv.Atoi(pieces[1]))

			box_idx := hash(label)

			found := false
			
			for i, lens := range boxes[box_idx] {
				if lens.label == label {
					found = true
					boxes[box_idx][i].focal_length = value

					break
				}
			}

			if !found {
				boxes[box_idx] = append(boxes[box_idx], lens{label, value})
			}
		} else if strings.Contains(instruction, "-") {
			label := instruction[:len(instruction) - 1]

			box_idx := hash(label)

			for i, curr_lens := range boxes[box_idx] {
				if curr_lens.label == label {
					new_list := make([]lens, 0)

					new_list = append(new_list, boxes[box_idx][0:i]...)
					new_list = append(new_list, boxes[box_idx][i+1:]...)

					boxes[box_idx] = new_list

					break
				}
			}
		} else {
			panic("Unreachable")
		}
	}

	power := 0

	for i, box := range boxes {
		for j, lens := range box {
			power += (i + 1) * (j + 1) * lens.focal_length
		}
	}

	fmt.Println(power)
}

/* I suppose I have acclimated to the lack of tuples, though I still want them. 
 * To be fair, Go makes it really really easy to declare lightweight structs, somehow 
 * I feel it is easier than in other languages. I guess I code a lot of Rust, where
 * even for the lightest of dumb data holders you find yourself going back again and
 * again to add #[derive(Clone, Copy, Debug, PartialOrd, Ord, PartialEq, Eq)].
 * That is probably a uniquely Rust problem.
 *
 * The lack of remove() method on slices really grinds my gears though. Slicing and
 * splat make this easy enough, but this should obviously be a method its not even funny. 
 *
 * Anyways the puzzle was pretty easy, actually there wasn't much puzzle to it. I really 
 * only had issues involving my lack of familiarity with Go: I figured out how to remove 
 * an element from a slice before, but there are two ways to do it - quick but not order
 * preserving, and slow and order preserving, and I picked the wrong one here. The other
 * issue - I implicitely assumed that Go's iteration gave you a reference to the data, but
 * it does not. Changing the focal_length of the [lens] that comes out of [for _, lens in range box] 
 * does not affect the underlying data. Which is fair, I just wasn't thinking and assumed 
 * the opposite. */