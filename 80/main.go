package main

import (
	"fmt"
)

func main() {
	var some = []int{1, 1, 1, 2, 2, 3}

	fmt.Println(removeDuplicates(some))
	fmt.Println(some)
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	tail := 0
	duplicatesCount := 0
	currentNumber := nums[0]

	for i := range nums {
		if nums[i] == currentNumber {
			duplicatesCount++
			if duplicatesCount > 2 {
				//tail++
				//duplicatesCount = 0
				continue
			}
		} else {
			currentNumber = nums[i]
			duplicatesCount = 1
			//duplicatesCount++
		}

		if i != tail {
			nums[tail] = nums[i]

		}
		tail++
	}
	return tail
}
