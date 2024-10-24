package main

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
				continue
			}
		} else {
			currentNumber = nums[i]
			duplicatesCount = 1
		}

		if i != tail {
			nums[tail] = nums[i]

		}
		tail++
	}
	return tail
}
