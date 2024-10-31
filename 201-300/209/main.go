package main

func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	pos1 := 0
	pos2 := 0
	sum := nums[0]
	minSize := 0

	for {
		if sum >= target {
			if minSize == 0 || pos2-pos1+1 < minSize {
				minSize = pos2 - pos1 + 1
			}
		}

		if pos1 == pos2 {
			if pos2 == 0 {
				pos2++
				if pos2 == len(nums) {
					return minSize
				}
				sum += nums[pos2]
				continue
			}
			return minSize
		}

		if pos2 == len(nums)-1 {
			sum -= nums[pos1]
			pos1++
			continue
		}
		if pos2-pos1 == 1 {
			pos2++
			sum += nums[pos2]
			continue
		}

		if sum-nums[pos1] >= target {
			sum -= nums[pos1]
			pos1++
			continue
		} else {
			pos2++
			sum += nums[pos2]
			continue
		}
	}
}
