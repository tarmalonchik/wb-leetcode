package main

func isPalindrome(s string) bool {
	firstIdx := 0
	lastIdx := len(s) - 1
	for {
		if firstIdx >= lastIdx {
			return true
		}
		if convert(s[firstIdx]) == '.' {
			firstIdx++
			continue
		}
		if convert(s[lastIdx]) == '.' {
			lastIdx--
			continue
		}
		if convert(s[firstIdx]) != convert(s[lastIdx]) {
			return false
		}
		firstIdx++
		lastIdx--
	}
}

func convert(in byte) byte {
	if (in >= 97 && in <= 122) || (in >= 48 && in <= 57) {
		return in
	}
	if in >= 65 && in <= 90 {
		return in + 32
	}
	return '.'
}
