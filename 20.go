package main

import "fmt"

func main() {
	fmt.Println(isValid("[()]"))
}

// 40 41 91 93 123 125

func isValid(s string) bool {
	var (
		stack []uint8
	)
	pairs := map[uint8]uint8{
		'{': '}',
		'[': ']',
		'(': ')',
	}

	for i := range s {
		if len(stack) == 0 {
			stack = append(stack, s[i])
		} else {
			if _, ok := pairs[s[i]]; ok {
				stack = append(stack, s[i])
			} else {
				if pairs[stack[len(stack)-1]] == s[i] {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
