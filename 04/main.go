package main

import "fmt"

func main() {
	flowerbed := []int{1, 0, 0, 0, 1}
	n := 2
	fmt.Println(canPlaceFlowers(flowerbed, n))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	length := len(flowerbed)
	for i := range flowerbed {
		isLeftBusy := false
		if i > 0 && flowerbed[i-1] == 1 {
			isLeftBusy = true
		}
		isRightBusy := false
		if i < length-1 && flowerbed[i+1] == 1 {
			isRightBusy = true
		}

		if !isRightBusy && !isLeftBusy && flowerbed[i] == 0 {
			flowerbed[i] = 1
			count++
		}
	}
	return count >= n
}
