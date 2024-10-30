package main

import (
	"fmt"
)

func main() {
	var x = 1
	var y = 10

	x ^= y
	y ^= x
	x ^= y

	fmt.Println(x)
	fmt.Println(y)
}
