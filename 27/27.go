package main

func main() {
	nums := []int{3, 3, 3}
	removeElement(nums, 3)
}

func removeElement(nums []int, val int) int {
	var (
		firstPos int
		lastPos  int
	)
	if len(nums) == 0 {
		return 0
	}
	lastPos = len(nums) - 1
	for {
		if lastPos < firstPos {
			nums = nums[:firstPos]
			return len(nums)
		}
		if nums[lastPos] == val {
			lastPos--
			continue
		}
		if nums[firstPos] == val {
			nums[firstPos] = nums[lastPos]
			lastPos--
			firstPos++
		} else {
			firstPos++
		}
	}
}
