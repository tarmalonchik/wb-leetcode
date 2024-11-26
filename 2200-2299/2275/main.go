package main

func largestCombination(candidates []int) int {
	maxNum := 0
	maxOut := 0
	for i := range candidates {
		if candidates[i] > maxNum {
			maxNum = candidates[i]
		}
	}

	maxCount := bitCount(maxNum)
	out := make([]int, maxCount)
	for i := uint8(0); i < maxCount; i++ {
		for j := range candidates {
			out[i] += getBit(uint32(candidates[j]), i)
			if out[i] >= maxOut {
				maxOut = out[i]
			}
		}
	}
	return maxOut
}

func getBit(num uint32, pos uint8) int {
	if num&(1<<pos) > 0 {
		return 1
	}
	return 0
}

func bitCount(in int) (len uint8) {
	for {
		if in == 0 {
			return len
		}
		if in%2 == 1 {
			len++
		} else {
			len++
		}
		in = in / 2
	}
}
