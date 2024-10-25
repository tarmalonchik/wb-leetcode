package main

func rotate(nums []int, k int) {
	if k > len(nums) {
		k = k % len(nums)
	}
	if len(nums) == 1 || k == len(nums) || k == 0 {
		return
	}

	for i := range nums {
		swap(nums, i, len(nums)-k+i)
		if i == k-1 {
			rotate(nums[k:], k)
			return
		}
		if len(nums)-k+i+1 == len(nums) {
			rotate(nums[i+1:], k-i-1)
			return
		}
	}
}

func swap(in []int, pos1, pos2 int) {
	swapItem := in[pos1]
	in[pos1] = in[pos2]
	in[pos2] = swapItem
}
