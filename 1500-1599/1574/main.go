package main

func valid(arr pointerManager, windowSize int, point int) bool {
	right := false
	if len(arr)-point <= point {
		right = true
	}

	pos1 := 0
	pos2 := 0

	if !right {
		pos1 = point - windowSize
		pos2 = point
	} else {
		pos1 = point
		pos2 = point + windowSize
	}

	if !right {
		for {
			if arr.getPrevIdx(pos1) == -1 {
				return true
			}

			if arr[arr.getPrevIdx(pos1)] <= arr[arr.getNextIdx(pos2)] {
				return true
			}

			if pos1 == point {
				return false
			}

			pos1++
			pos2++
		}
	} else {
		for {
			if arr.getNextIdx(pos2) == -1 {
				return true
			}
			if arr[arr.getPrevIdx(pos1)] <= arr[arr.getNextIdx(pos2)] {
				return true
			}

			if pos2 == point {
				return false
			}

			pos1--
			pos2--
		}
	}
}

type pointerManager []int

func (p *pointerManager) getPrevIdx(idx int) int {
	if idx == 0 {
		return -1
	}
	return idx - 1
}

func (p *pointerManager) getNextIdx(idx int) int {
	if idx == len(*p) {
		return -1
	}
	return idx
}

func findLengthOfShortestSubarray(arr []int) int {
	if len(arr) == 0 || len(arr) == 1 {
		return 0
	}
	point1 := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] >= arr[i-1] {
			point1 = i
		} else {
			break
		}
	}
	if point1 == len(arr)-1 {
		return 0
	}
	point1++

	point2 := len(arr) - 1
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i+1] >= arr[i] {
			point2 = i
		} else {
			break
		}
	}

	newArr := append(arr[:point1], arr[point2:]...)

	out := len(arr) - len(newArr)

	minSize := 0
	maxSize := point1

	if len(newArr)-point1 < maxSize {
		maxSize = len(newArr) - point1
	}

	for {
		if maxSize-minSize <= 1 {
			if valid(newArr, minSize, point1) {
				return out + minSize
			}
			if valid(newArr, maxSize, point1) {
				return out + maxSize
			}
		}

		center := minSize + (maxSize-minSize)/2
		if valid(newArr, center, point1) {
			maxSize = center
		} else {
			minSize = center
		}
	}
}
