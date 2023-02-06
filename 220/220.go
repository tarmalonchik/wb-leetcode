package main

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if len(nums) < 2 {
		return false
	}
	for kVar := k; kVar > 0; kVar-- {
		if kVar >= len(nums) && len(nums) > 1 {
			kVar = len(nums) - 1
		}
		first := 0
		last := first + kVar
		for last < len(nums) {
			result := nums[last] - nums[first]
			if result < 0 {
				result = -result
			}
			if result <= t {
				return true
			}
			first++
			last++
		}
	}
	return false
}
