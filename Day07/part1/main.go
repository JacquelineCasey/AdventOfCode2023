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

	for _, v := range frequencies {
		counts = append(counts, v)
	}

	slices.Sort(counts)

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
			case '2': hand.cards[i] = 1
			case '3': hand.cards[i] = 2
			case '4': hand.cards[i] = 3
			case '5': hand.cards[i] = 4
			case '6': hand.cards[i] = 5
			case '7': hand.cards[i] = 6
			case '8': hand.cards[i] = 7
			case '9': hand.cards[i] = 8
			case 'T': hand.cards[i] = 9
			case 'J': hand.cards[i] = 10
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
