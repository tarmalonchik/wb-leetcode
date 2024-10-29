package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxSubArray([]int{0}))
}

// Kadan algo O(n)
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	newArr := make([]int, len(nums))

	for i := range nums {
		if i-1 < 0 {
			newArr[i] = nums[i]
			continue
		}
		if nums[i]+newArr[i-1] > nums[i] {
			newArr[i] = nums[i] + newArr[i-1]
		} else {
			newArr[i] = nums[i]
		}
	}

	maxVal := newArr[0]
	for i := 1; i < len(newArr); i++ {
		if newArr[i] > maxVal {
			maxVal = newArr[i]
		}
	}
	return maxVal
}
