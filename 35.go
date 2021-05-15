package main

func searchInsert(nums []int, target int) int {
	for i := range nums{
		if target <= nums[i]{
			return i
		}
	}
	return len(nums)
}