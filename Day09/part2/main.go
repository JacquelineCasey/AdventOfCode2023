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

func all_zero(nums []int) bool {
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}

	return true
}

/* Now we just extrapolate backwards. */
func extrapolate(nums []int) int {
	if all_zero(nums) {
		return 0
	}

	differences := make([]int, 0)

	for i := range nums {
		if i == 0 {
			continue
		}

		differences = append(differences, nums[i] - nums[i-1])
	}

	delta := extrapolate(differences)

	return nums[0] - delta  // (*) Only change
}

func main() {
    data := string(check(io.ReadAll(os.Stdin)))

	lines := strings.Split(strings.Trim(data, "\n"), "\n")

	sum := 0
	for _, line := range lines {
		nums := make([]int, 0)

		for _, str := range strings.Split(line, " ") {
			nums = append(nums, check(strconv.Atoi(str)))	
		}

		sum += extrapolate(nums)
	}

	fmt.Println(sum)
}

/* When the puzzle is easy, Go is unitrusive. Perhaps I was somewhat lucky in spotting 
 * the recursive solution to this puzzle, as I think the naive solution involves a 2D 
 * array, which I suspect would be quite cumbersome. */