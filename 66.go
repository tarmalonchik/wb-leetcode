package main

func plusOne(digits []int) []int {
	for a := len(digits) - 1; a >= 0; a-- {
		digits[a]++
		if digits[a] < 10 {
			return digits
		}
		digits[a] = 0
	}
	a := []int{1}

	return append(a, digits...)
}
