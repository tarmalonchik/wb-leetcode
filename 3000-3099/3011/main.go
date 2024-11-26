package main

func canSortArray(nums []int) bool {
	if len(nums) == 0 || len(nums) == 1 {
		return true
	}
	startPos := 0

	prevMax := -1

	for {
		if startPos >= len(nums) {
			return true
		}

		var newMin, newMax int

		newMin, newMax, startPos = processSubArray(nums, startPos)
		if prevMax == -1 {
			prevMax = newMax
			continue
		}
		if prevMax > newMin {
			return false
		}
		prevMax = newMax
	}
}

func processSubArray(nums []int, startPos int) (min, max, newPos int) {
	prevNum := nums[startPos]
	prevSetBitsCount := getSetBitsCount(prevNum)
	min = prevNum
	max = prevNum

	for {
		startPos++
		if startPos == len(nums) {
			break
		}
		newSetBitCount := getSetBitsCount(nums[startPos])
		if prevSetBitsCount != newSetBitCount {
			break
		}

		if nums[startPos] > max {
			max = nums[startPos]
		}
		if nums[startPos] < min {
			min = nums[startPos]
		}
	}

	return min, max, startPos
}

func getSetBitsCount(in int) (count int) {
	for {
		if in == 0 {
			return count
		}
		if in%2 == 1 {
			count++
		}
		in = in / 2
	}
}
