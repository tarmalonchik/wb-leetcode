package main

import (
	"fmt"
)

func main() {
	fmt.Println(maximumSubarraySum([]int{1, 3, 3, 3, 3, 2, 1, 3, 1, 1, 2}, 3))
}

func maximumSubarraySum(nums []int, k int) int64 {
	if len(nums) < k {
		return 0
	}

	sum := int64(0)
	currentSum := int64(0)

	oneTimeMP := make(map[int]interface{})
	multipleTimeMp := make(map[int]int)

	for i := 0; i < k; i++ {
		sum += int64(nums[i])
		addItem(oneTimeMP, multipleTimeMp, nums[i])
	}
	currentSum = sum
	if len(multipleTimeMp) != 0 {
		sum = 0
	}

	pos1 := 0
	pos2 := k - 1

	for {
		if pos2 == len(nums)-1 {
			return sum
		}
		pos2++
		currentSum += int64(nums[pos2])
		addItem(oneTimeMP, multipleTimeMp, nums[pos2])
		currentSum -= int64(nums[pos1])
		removeItem(oneTimeMP, multipleTimeMp, nums[pos1])
		pos1++

		if len(multipleTimeMp) == 0 {
			sum = getMax(currentSum, sum)
		}
	}
}

func removeItem(oneTime map[int]interface{}, multipleTime map[int]int, val int) {
	_, ok := oneTime[val]
	if ok {
		delete(oneTime, val)
	} else {
		number, ok := multipleTime[val]
		if !ok {
			return
		}
		if number == 2 {
			delete(multipleTime, val)
			oneTime[val] = nil
			return
		} else if number == 1 {
			panic("invalid number")
		} else {
			multipleTime[val]--
		}
	}
}

func addItem(oneTime map[int]interface{}, multipleTime map[int]int, val int) {
	_, ok := oneTime[val]
	if ok {
		delete(oneTime, val)
		multipleTime[val] = 2
	} else {
		_, mulOk := multipleTime[val]
		if mulOk {
			multipleTime[val]++
			return
		}
		oneTime[val] = nil
	}
}

func getMax(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}
