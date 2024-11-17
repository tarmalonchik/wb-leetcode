package main

func sortColors(nums []int) {
	zeros := 0
	ones := 0
	twos := 0
	for i := range nums {
		if nums[i] == 0 {
			zeros++
		} else if nums[i] == 1 {
			ones++
		} else if nums[i] == 2 {
			twos++
		}
	}

	for i := range nums {
		if zeros != 0 {
			zeros--
			nums[i] = 0
		} else if ones != 0 {
			ones--
			nums[i] = 1
		} else if twos != 0 {
			twos--
			nums[i] = 2
		}
	}
}
