package main

func jump(nums []int) int {
	var counter, sliderIndex int
	if len(nums) == 0 || len(nums) == 1 {
		return 0
	}
	for {
		counter++
		maxIndex := sliderIndex + nums[sliderIndex]
		if maxIndex >= len(nums)-1 {
			return counter
		}
		maxVal := nums[maxIndex]
		for i := sliderIndex + nums[sliderIndex]; i > sliderIndex; i-- {
			if nums[i] == 0 {
				continue
			}
			val := nums[i] - (sliderIndex + nums[sliderIndex] - i)
			if val > maxVal {
				maxIndex = i
				maxVal = val
			}
		}
		sliderIndex = maxIndex
	}
}
