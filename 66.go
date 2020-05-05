package main

func plusOne(digits []int) []int {
	value := 1
	for a := len(digits) - 1; a >= 0; a-- {
		if value == 1 {
			digits[a] ++
			value = digits[a] / 10
			digits[a] = digits[a] % 10
		}else{
			return digits
		}
	}
	a := []int{1}
	return append(a,digits ...)
}

