package main

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	idx := findRotatedIndex(nums)
	if idx == -1 {
		return searchWithPos(nums, 0, len(nums)-1, target)
	}

	if target >= nums[0] && target <= nums[idx] {
		return searchWithPos(nums, 0, idx, target)
	}
	return searchWithPos(nums, idx+1, len(nums)-1, target)
}

func searchWithPos(nums []int, pos1, pos2, target int) int {
	if pos2-pos1 == 0 {
		if nums[pos1] == target {
			return pos1
		}
		return -1
	} else if pos2-pos1 == 1 {
		if nums[pos1] == target {
			return pos1
		}
		if nums[pos2] == target {
			return pos2
		}
		return -1
	}
	center := pos1 + ((pos2 - pos1) / 2)

	if target < nums[center] {
		return searchWithPos(nums, pos1, center, target)
	}
	return searchWithPos(nums, center, pos2, target)
}

func findRotatedIndex(nums []int) int {
	if len(nums) <= 1 {
		return -1
	}

	if nums[0] < nums[len(nums)-1] {
		return -1
	}

	return recursive(nums, 0, len(nums)-1)
}

func recursive(nums []int, pos1, pos2 int) int {
	if pos2 == pos1 {
		return -1
	}
	if pos2-pos1 == 1 {
		if nums[pos1] > nums[pos2] {
			return pos1
		}
	}

	center := pos1 + (pos2-pos1)/2
	if nums[pos1] > nums[center] {
		return recursive(nums, pos1, center)
	}
	return recursive(nums, center, pos2)
}
