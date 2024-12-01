package main

import (
	"sort"
)

func countFairPairs(nums []int, lower int, upper int) int64 {
	out := int64(0)

	if len(nums) <= 1 {
		return out
	}

	sort.Ints(nums)

	lPos := 0
	rPos := len(nums) - 1

	for {
		if rPos == lPos {
			return 0
		}
		sum := nums[lPos] + nums[rPos]

		if sum >= lower && sum <= upper {
			break
		} else if sum < lower {
			lPos++
		} else if sum > upper {
			rPos--
		}
	}

	center := rPos
	for {
		if center-1 == lPos {
			break
		}

		sum := nums[lPos] + nums[center-1]
		if sum < lower {
			break
		}
		center--
	}

	increase := true

	for {
		if increase {
			out += int64((rPos - center) + 1)
		}
		lPos++
		if center == lPos {
			center++
			if center > rPos {
				return out
			}
		}

		if nums[lPos]+nums[center] >= lower {
			for {
				if center-1 <= lPos {
					break
				}
				sum := nums[lPos] + nums[center-1]
				if sum >= lower {
					center--
				} else {
					break
				}
			}
		}

		if nums[lPos]+nums[rPos] > upper {
			for {
				if rPos-1 < center {
					increase = false
					break
				}
				rPos--
				if rPos <= lPos {
					return out
				}
				sum := nums[lPos] + nums[rPos]
				if sum <= upper {
					increase = true
					break
				}
			}
		}
	}
	return out
}
