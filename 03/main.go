package main

import (
	"fmt"
	"slices"
)

func main() {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3

	fmt.Println(kidsWithCandies(candies, extraCandies))
	fmt.Println(candies)
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	result := []bool{}
	candiesCopy := make([]int, len(candies))

	copy(candiesCopy, candies)

	for i := range candies {
		candiesCopy[i] = candiesCopy[i] + extraCandies
	}

	max := slices.Max(candies)
	for _, candy := range candiesCopy {
		if candy >= max {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}

	return result
}
