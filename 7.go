package main

func reverse(x int) int {
	result := 0
	for {
		if x/10 == 0{
			return result*10 + x%10
		}
		result *= 10
		result += x % 10
		x = x/10
	}
}
