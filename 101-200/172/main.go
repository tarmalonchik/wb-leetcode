package main

func trailingZeroes(n int) int {
	var resp int
	for {
		n = n / 5
		if n == 0 {
			return resp
		}
		resp += n
	}
}
