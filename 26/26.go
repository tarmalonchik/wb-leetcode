package main

func removeDuplicates(nums []int) int {
	var (
		firstPos int
		lastPos  int
	)

	if len(nums) < 2 {
		return len(nums)
	}

	for {
		nums[firstPos] = nums[lastPos]

		lastPos++
		if nums[lastPos] > nums[lastPos-1] {
			if lastPos == 5 {
			}
			firstPos++
			nums[firstPos] = nums[lastPos]
		}

		if lastPos == len(nums)-1 {
			nums = nums[:firstPos+1]
			return len(nums)
		}
	}
}
