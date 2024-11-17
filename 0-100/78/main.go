package main

func subsets(nums []int) (out [][]int) {
	out = subsetsRecursion(nums)
	out = append(out, []int{})
	return out
}

func subsetsRecursion(nums []int) (out [][]int) {
	out = [][]int{
		{nums[0]},
	}

	if len(nums) == 1 {
		return out
	}

	subData := subsetsRecursion(nums[1:])
	for i := range subData {
		out = append(out, subData[i])
		withPrefix := make([]int, 1, len(subData[i])+1)
		withPrefix[0] = nums[0]
		withPrefix = append(withPrefix, subData[i]...)
		out = append(out, withPrefix)
	}
	return out
}
