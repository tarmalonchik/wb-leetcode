package main

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	length := -1

	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i]+length >= 0 {
			if i == 0 {
				return true
			}
			length = -1
			continue
		} else {
			if i == 0 {
				return false
			}
			length--
		}
	}
	return false
}
