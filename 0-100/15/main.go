package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	var out [][]int
	sort.Ints(nums)
	for i := range nums {
		if i > 0 {
			if nums[i] == nums[i-1] {
				continue
			}
		}
		resp := twoSum(nums[i+1:], -nums[i])
		if resp == nil {
			continue
		}

		for j := range resp {
			resp[j][0] = nums[i+1:][resp[j][0]-1]
			resp[j][1] = nums[i+1:][resp[j][1]-1]
			out = append(out, append(resp[j], nums[i]))
		}

	}
	return out
}

func twoSum(numbers []int, target int) (out [][]int) {
	if len(numbers) == 0 {
		return nil
	}
	pos1 := 0
	pos2 := len(numbers) - 1
	for {
		if pos1 == pos2 {
			return out
		}
		if pos1 > 0 && numbers[pos1] == numbers[pos1-1] {
			pos1++
			continue
		}
		if pos2 < len(numbers)-1 && numbers[pos2] == numbers[pos2+1] {
			pos2--
			continue
		}
		sum := numbers[pos1] + numbers[pos2]
		if sum == target {
			out = append(out, []int{pos1 + 1, pos2 + 1})
			pos1++
			continue
		}
		if sum > target {
			pos2--
		} else {
			pos1++
		}
	}
}
