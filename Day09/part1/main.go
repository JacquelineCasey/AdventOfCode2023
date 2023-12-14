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

	return nums[len(nums) - 1] + delta
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
