package main

import (
	"fmt"
	"io"
	"os"
	"slices"
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

type hand struct {
	cards [5]int
	bid int
}

/* Returns an integer such that higher ints represent more powerful hand classes. */
func class(hand hand) int {
	frequencies := make(map[int] int)

	for _, c := range hand.cards {
		frequencies[c]++
	}

	counts := make([]int, 0)
	jokers := 0

	for k, v := range frequencies {
		if k == 1 {  // Joker
			jokers = v
		} else {
			counts = append(counts, v)
		}
	}

	slices.Sort(counts)
	
	/* Jokers always morph to join the last card type. */ 
	if len(counts) != 0 {
		counts[len(counts) - 1] += jokers
	} else {
		counts = []int{5}
	}

	switch {
	case slices.Equal(counts, []int{5}): 
		return 7
	case slices.Equal(counts, []int{1, 4}): 
		return 6
	case slices.Equal(counts, []int{2, 3}): 
		return 5
	case slices.Equal(counts, []int{1, 1, 3}): 
		return 4
	case slices.Equal(counts, []int{1, 2, 2}): 
		return 3
	case slices.Equal(counts, []int{1, 1, 1, 2}): 
		return 2
	default: 
		return 1
	}
}

func compare(hand1 hand, hand2 hand) int {
	class1 := class(hand1)
	class2 := class(hand2)

	if class1 != class2 {
		return class1 - class2
	}

	/* Otherwise default to card by card comparison */
	for i, c1 := range hand1.cards {
		c2 := hand2.cards[i]

		if c1 != c2 {
			return c1 - c2
		}
	}

	return 0
}

func main() {
    data := string(check(io.ReadAll(os.Stdin)))

	lines := strings.Split(strings.Trim(data, "\n"), "\n")

	hands := make([]hand, 0)
	for _, line := range lines {
		line = strings.Trim(line, "\n")
		pieces := strings.Split(line, " ")

		var hand hand

		for i, ch := range pieces[0] {
			switch ch {
			case 'J': hand.cards[i] = 1 // (*)
			case '2': hand.cards[i] = 2
			case '3': hand.cards[i] = 3
			case '4': hand.cards[i] = 4
			case '5': hand.cards[i] = 5
			case '6': hand.cards[i] = 6
			case '7': hand.cards[i] = 7
			case '8': hand.cards[i] = 8
			case '9': hand.cards[i] = 9
			case 'T': hand.cards[i] = 10
			case 'Q': hand.cards[i] = 11
			case 'K': hand.cards[i] = 12
			case 'A': hand.cards[i] = 13
			}
		}

		hand.bid = check(strconv.Atoi(pieces[1]))

		hands = append(hands, hand)
	}

	/* We sort ascending */
	slices.SortFunc(hands, compare)

	sum := 0
	for i, hand := range hands {
		sum += hand.bid * (i + 1)
	}

	fmt.Println(sum)
}

/* Cool puzzle. 
 *
 * Go was pretty chill. I'm missing the convenience of comparing slices directly, 
 * but at least there was a method to do it although it turns out this method, and
 * the rest of the slices package, were added only 4 months ago! How did people manage
 * without it - Go seems to have no qualms about making every programmer rewrite the
 * same for loop, or maybe import a really really lightweight library. Actually, I 
 * also used the sort function in that module - it seems like the old way involved 
 * wrapping your data in a particular type.
 *
 * I'm a little annoyed that the Go comparison functions return an int, instead of 
 * some sort of cmp::{LT, GT, EQ} thing. I find the int version makes code way more
 * confusing, though I'll admit that the ability to subtract and get a valid comparison 
 * result can be nice. 
 *
 * I really wanted a method to convert a map into a slice of its values. I couldn't 
 * find it. 
 *
 * Honestly, if Go just had a bigger standard library, I think I'd be a lot happier 
 * right now. The /language/ is fine, I'm annoyed by the standard library, and a 
 * little bit by the documentation. */