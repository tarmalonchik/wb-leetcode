package main

// https://en.wikipedia.org/wiki/Boyer–Moore_majority_vote_algorithm

func majorityElement(nums []int) int {
	currentSym := 0
	sum := 0
	for i := range nums {
		if sum == 0 {
			currentSym = nums[i]
			sum++
			continue
		}
		if currentSym == nums[i] {
			sum++
		} else {
			sum--
		}
	}
	return currentSym
}
