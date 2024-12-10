package main

func minimumSize(nums []int, maxOperations int) int {
	maxSizeOrigin := 0
	for i := range nums {
		if nums[i] > maxSizeOrigin {
			maxSizeOrigin = nums[i]
		}
	}

	minSize := 1
	maxSize := maxSizeOrigin

	for {
		if maxSize-minSize <= 1 {
			if possible(nums, maxOperations, minSize) {
				return minSize
			}
			if possible(nums, maxOperations, maxSize) {
				return maxSize
			}
			return maxSizeOrigin
		}

		center := minSize + (maxSize-minSize)/2
		if possible(nums, maxOperations, center) {
			maxSize = center
		} else {
			minSize = center
		}
	}
}

func possible(nums []int, maxOperations int, maxSize int) bool {
	for i := range nums {
		if nums[i] <= maxSize {
			continue
		}
		remain := nums[i] % maxSize
		division := nums[i] / maxSize

		division--
		if remain != 0 {
			division++
		}

		maxOperations -= division
		if maxOperations < 0 {
			return false
		}
	}
	return true
}
