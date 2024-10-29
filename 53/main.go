package main

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	centerFirst := len(nums) / 2
	centerLast := centerFirst
	centerMax := nums[centerFirst]
	centerSum := centerMax
	leftBalance := 0
	rightBalance := 0

	for {
		if centerFirst == 0 && centerLast == len(nums)-1 {
			break
		} else if centerFirst == 0 {
			centerLast++
			centerSum += nums[centerLast]

			rightBalance += nums[centerLast]
			if rightBalance > 0 {
				rightBalance = 0
			}

			if centerSum-rightBalance-leftBalance > centerMax {
				centerMax = centerSum - rightBalance - leftBalance
			}
			continue
		} else if centerLast == len(nums)-1 {
			centerFirst--
			centerSum += nums[centerFirst]

			leftBalance += nums[centerFirst]
			if leftBalance >= 0 {
				leftBalance = 0
			}

			if centerSum-rightBalance-leftBalance > centerMax {
				centerMax = centerSum - rightBalance - leftBalance
			}
			continue
		}

		if nums[centerFirst-1] >= nums[centerLast+1] {
			centerFirst--
			centerSum += nums[centerFirst]

			leftBalance += nums[centerFirst]
			if leftBalance >= 0 {
				leftBalance = 0
			}

			if centerSum-rightBalance-leftBalance > centerMax {
				centerMax = centerSum - rightBalance - leftBalance
			}
			continue
		} else {
			centerLast++
			centerSum += nums[centerLast]

			rightBalance += nums[centerLast]
			if rightBalance > 0 {
				rightBalance = 0
			}

			if centerSum-rightBalance-leftBalance > centerMax {
				centerMax = centerSum - rightBalance - leftBalance
			}
			continue
		}
	}

	leftMax := maxSubArray(nums[len(nums)/2:])
	rightMax := maxSubArray(nums[:len(nums)/2])

	if centerMax >= leftMax && centerMax >= rightMax {
		return centerMax
	}

	if rightMax >= leftMax {
		return rightMax
	}

	return leftMax
}

// Kadan algo O(n)
// https://en.wikipedia.org/wiki/Maximum_subarray_problem
func maxSubArrayKadan(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	newArr := make([]int, len(nums))

	for i := range nums {
		if i-1 < 0 {
			newArr[i] = nums[i]
			continue
		}
		if nums[i]+newArr[i-1] > nums[i] {
			newArr[i] = nums[i] + newArr[i-1]
		} else {
			newArr[i] = nums[i]
		}
	}

	maxVal := newArr[0]
	for i := 1; i < len(newArr); i++ {
		if newArr[i] > maxVal {
			maxVal = newArr[i]
		}
	}
	return maxVal
}
