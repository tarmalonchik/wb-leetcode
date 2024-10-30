package main

func hammingWeight(n int) int {
	var resp int
	for {
		if n == 1 {
			resp++
			return resp
		}
		if n%2 == 1 {
			resp++
		}
		n = n / 2
	}
}
