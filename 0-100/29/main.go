package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(divide(-2147483648, -1))
}

func divide(dividend int, divisor int) int {
	var oneMinus bool
	var twoMinus bool

	if dividend < 0 {
		dividend = -dividend
		oneMinus = true
	}
	if divisor < 0 {
		divisor = -divisor
		twoMinus = true
	}

	var resp int
	num := divisor
	for {
		if num > dividend {
			if oneMinus && !twoMinus {
				resp = -resp
			}
			if twoMinus && !oneMinus {
				resp = -resp
			}
			if resp > math.MaxInt32 {
				resp = math.MaxInt32
			}
			if resp < math.MinInt32 {
				resp = math.MinInt32
			}
			return resp
		}
		num += divisor
		resp++
	}
}
