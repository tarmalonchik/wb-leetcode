package main

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return []int{0, 0}
		}
		return []int{-1, -1}
	}
	if nums[0] > target {
		return []int{-1, -1}
	}
	if nums[len(nums)-1] < target {
		return []int{-1, -1}
	}
	founded := binarySearch(nums, target)
	if founded == -1 {
		return []int{-1, -1}
	}
	lPos := founded
	rPos := founded

	for {
		counter := 0
		if lPos != 0 && nums[lPos-1] == target {
			lPos--
			counter++
		}
		if rPos != len(nums)-1 && nums[rPos+1] == target {
			rPos++
			counter++
		}
		if counter == 0 {
			return []int{lPos, rPos}
		}
	}

}

func binarySearch(nums []int, target int) int {
	lPos := 0
	rPos := len(nums) - 1
	for {
		if nums[rPos] == target {
			return rPos
		}
		if nums[lPos] == target {
			return lPos
		}
		if rPos-lPos == 1 {
			return -1
		}
		newPos := lPos + (rPos-lPos)/2
		if nums[newPos] >= target {
			rPos = newPos
		} else {
			lPos = newPos
		}
	}
}
