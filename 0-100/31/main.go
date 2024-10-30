package main

func nextPermutation(nums []int) {
	if swapped(nums) {
		return
	}
	reorder(nums)
}

func swapped(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	rPos := len(nums) - 1
	lPos := len(nums) - 2

	for {
		if nums[rPos] > nums[lPos] {
			swap(nums, lPos, rPos+binarySearchGreater(nums[rPos:], nums[lPos]))
			reorder(nums[rPos:])
			return true
		}
		if lPos == 0 {
			return false
		}
		rPos--
		lPos--
	}
}

func reorder(nums []int) {
	first := 0
	last := len(nums) - 1
	for {
		swap(nums, first, last)
		first++
		last--
		if last <= first {
			return
		}
	}
}

func swap(arr []int, index1, index2 int) {
	loc := arr[index1]
	arr[index1] = arr[index2]
	arr[index2] = loc
}

func binarySearchGreater(arr []int, num int) int {
	pos1 := 0
	pos2 := len(arr) - 1
	for {
		if (pos2-pos1)%2 == 1 {
			if arr[pos2] > num {
				return pos2
			} else {
				pos2--
				continue
			}
		}
		if (pos2 - pos1) == 0 {
			return pos1
		}
		if arr[(pos2+pos1)/2] > num {
			pos1 = (pos2 + pos1) / 2
		} else {
			pos2 = (pos2 + pos1) / 2
		}
	}
}
