package main

func grayCode(n int) []int {
	out := make([]int, pow2(n))
	if n == 0 {
		return nil
	}
	out[0] = 0
	out[1] = 1
	if n == 1 {
		return out[:2]
	}
	for i := 1; i < n; i++ {
		grayCodeCounter(out, i)
	}
	return out
}

func grayCodeCounter(input []int, iteration int) {
	downPos := pow2(iteration) - 1
	upPos := pow2(iteration)

	downFirstValue := input[downPos]
	upFirstValue := input[downPos] + pow2(iteration)

	for {
		offset := downFirstValue - input[downPos]
		input[upPos] = upFirstValue - offset
		downPos--
		upPos++
		if downPos == -1 {
			break
		}
	}
}

func pow2(a int) int {
	if a == 0 {
		return 1
	}
	return 2 << (a - 1)
}
