package main

func reverse(x int) int {
	result := 0
	for {
		if x/10 == 0 {
			answer := result*10 + x
			if answer > 2147483647 || answer < -2147483648 {
				return 0
			}
			return answer
		}
		result *= 10
		result += x % 10
		x = x / 10
	}
}
