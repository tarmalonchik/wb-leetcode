package main

func getMaximumXor(nums []int, maximumBit int) []int {
	currentXor := 0

	xorArr := make([]int, len(nums))
	for i := range nums {
		currentXor = currentXor ^ nums[i]
		xorArr[len(nums)-1-i] = currentXor
	}

	for i := range xorArr {
		xorArr[i] = int(countMaxXor(uint32(xorArr[i]), maximumBit))
	}
	return xorArr
}

func countMaxXor(num uint32, maxBit int) (out uint32) {
	out = ^num
	outMask := uint32(1 << (maxBit))
	outMask--
	return out & outMask
}
