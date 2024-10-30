package main

func singleNumber(nums []int) int {
	var positiveCounter = make([]int, 64)
	var negativeCounter = make([]int, 64)
	for i := range nums {
		if nums[i] >= 0 {
			addNum(positiveCounter, nums[i])
		} else {
			addNum(negativeCounter, abs(nums[i]))
		}
	}

	positive := getNumFromCounter(positiveCounter)
	negative := getNumFromCounter(negativeCounter)
	if positive != 0 {
		return positive
	}
	return -negative
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func getNumFromCounter(counter []int) (resp int) {
	for i := len(counter) - 1; i >= 0; i-- {
		if counter[i] == 0 {
			continue
		}
		resp = setBit(resp, uint(len(counter)-1-i))
	}
	return resp
}

func addNum(counter []int, num int) {
	for i := len(counter) - 1; i >= 0; i-- {
		if hasBit(num, uint(len(counter)-1-i)) {
			counter[i]++
		}
		if counter[i] == 3 {
			counter[i] = 0
		}
	}
}

func setBit(n int, pos uint) int {
	return n | (1 << pos)

}

func hasBit(n int, pos uint) bool {
	val := n & (1 << pos)
	return val > 0
}
