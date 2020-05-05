package main

func plusOne(digits []int) []int {
	result := make([]int, len(digits))
	value := 1
	for a := len(digits) - 1; a >= 0; a-- {
		if value == 1 {
			digits[a] += value
			value = digits[a] / 10
			digits[a] = digits[a] % 10
		}
		result[a] = digits[a]
	}
	if result[0] == 0 {
		a := []int{1}
		return append(a,result ...)
	}
	return result
}

