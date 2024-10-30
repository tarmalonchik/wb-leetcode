package main

func maxArea(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	point1 := 0
	point2 := len(height) - 1
	max := (point2 - point1) * minHeight(height[point1], height[point2])
	if point1 == point2 {
		return max
	}
	for {
		if height[point1] >= height[point2] {
			point2--
		} else {
			point1++
		}
		newMax := (point2 - point1) * minHeight(height[point1], height[point2])
		if newMax > max {
			max = newMax
		}

		if point1 == point2 {
			break
		}
	}
	return max
}

func minHeight(a, b int) int {
	if a < b {
		return a
	}
	return b
}
