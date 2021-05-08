package main

func romanToInt(s string) int {
	var (
		response int
		previous = 1001
	)
	symbols := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	for i := range s {
		if previous < symbols[s[i]] {
			response -= 2 * previous

		}
		response += symbols[s[i]]
		previous = symbols[s[i]]
	}
	return response
}
