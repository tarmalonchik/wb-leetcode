package main

func isHappy(n int) bool {
	mp := make(map[int]interface{})
	if n == 0 {
		return false
	}
	for {
		if n == 1 {
			return true
		}
		n = getNext(n)
		_, ok := mp[n]
		if ok {
			return false
		}
		mp[n] = nil
	}
}

func getNext(n int) int {
	var digits []int
	for {
		rem := n % 10
		n = n / 10
		digits = append(digits, rem)
		if n == 0 {
			break
		}
	}
	if len(digits) == 0 {
		return 0
	}

	out := digits[0] * digits[0]
	for i := 1; i < len(digits); i++ {
		out += digits[i] * digits[i]
	}
	return out
}
