package main

func twoSum(numbers []int, target int) []int {
	pos1 := 0
	pos2 := len(numbers) - 1
	for {
		if pos1 == pos2 {
			return nil
		}
		sum := numbers[pos1] + numbers[pos2]
		if sum == target {
			return []int{pos1 + 1, pos2 + 1}
		}
		if sum > target {
			pos2--
		} else {
			pos1++
		}
	}
}
