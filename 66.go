package main

func plusOne(digits []int) []int {
	result := make([]int, len(digits) + 1)
	value := 1
	for a := len(digits) - 1; a >= 0; a-- {
		if value == 1 {
			digits[a] += value
			value = digits[a] / 10
			digits[a] = digits[a] % 10
		}
		result[a+1] = digits[a]
	}
	if result[1] == 0 {
		result[0] = 1
		return result
	}
	return result[1:]
}

