package main

import (
	"fmt"
)

func main() {
	fmt.Println(mySqrt(625))
}

func mySqrt(x int) int {
	if x == 1 {
		return x
	}
	start, end := 0, x
	for {
		if end-start <= 1 {
			return start
		}
		test := (end - start) / 2
		if (start+test)*(start+test) > x {
			end -= test
		} else {
			start += test
		}
	}
}
