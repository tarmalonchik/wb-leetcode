package main

func trap(height []int) int {
	volume := 0
	if len(height) < 3 {
		return 0
	}
	left := 0
	right := len(height) - 1

	leftMax := height[left]
	rightMax := height[right]
	for {
		if height[left] < height[right] {
			left++
			if height[left] < leftMax {
				volume += leftMax - height[left]
			} else {
				leftMax = height[left]
			}
		} else {
			right--
			if height[right] < rightMax {
				volume += rightMax - height[right]
			} else {
				rightMax = height[right]
			}
		}
		if right-left <= 1 {
			break
		}
	}
	return volume
}
