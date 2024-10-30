package main

// using XOR operator https://habr.com/ru/companies/vdsina/articles/538298/
func singleNumber(nums []int) int {
	var result int
	for i := range nums {
		result = result ^ nums[i]
	}
	return result
}
