package main

import (
	"sort"
)

func maximumBeauty(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)

	maxLen := 1
	distances := make([]point, 0, len(nums))
	distances = append(distances, point{value: nums[0], position: 0})

	numsPos := 0
	distancesPos := 0
	for {
		if numsPos == len(nums) {
			newMax := numsPos - distances[distancesPos].position
			if newMax > maxLen {
				maxLen = newMax
			}
			break
		}
		if nums[numsPos] == distances[distancesPos].value {
			numsPos++
			continue
		}
		if distances[distancesPos].value+2*k >= nums[numsPos] {
			newMax := numsPos - distances[distancesPos].position + 1
			if newMax > maxLen {
				maxLen = newMax
			}
			distances = append(distances, point{
				value:    nums[numsPos],
				position: numsPos,
			})
			numsPos++
		} else {
			newMax := numsPos - distances[distancesPos].position
			if newMax > maxLen {
				maxLen = newMax
			}
			distances = append(distances, point{
				value:    nums[numsPos],
				position: numsPos,
			})
			distancesPos++
		}
		if numsPos == len(nums) {
			continue
		}
	}
	return maxLen
}

type point struct {
	position int
	value    int
}
