package main

func permute(nums []int) [][]int {
	var resp = make([][]int, findFactorial(len(nums)))

	mask := make([]int, 0, len(nums))
	for i := range nums {
		mask = append(mask, i+1)
	}

	for i := 0; i < findFactorial(len(nums)); i++ {
		resp[i] = make([]int, len(nums))
		copy(resp[i], nums)
		nextPermutation(mask, nums)
	}
	return resp
}

func nextPermutation(mask, slave []int) {
	if swapped(mask, slave) {
		return
	}
	reorder(mask, slave)
}

func swapped(mask, slave []int) bool {
	if len(mask) <= 1 {
		return true
	}
	rPos := len(mask) - 1
	lPos := len(mask) - 2

	for {
		if mask[rPos] > mask[lPos] {
			pos := binarySearchGreater(mask[rPos:], mask[lPos])
			swap(mask, lPos, rPos+pos)
			swap(slave, lPos, rPos+pos)
			reorder(mask[rPos:], slave[rPos:])
			return true
		}
		if lPos == 0 {
			return false
		}
		rPos--
		lPos--
	}
}

func reorder(nums, slave []int) {
	first := 0
	last := len(nums) - 1
	for {
		swap(nums, first, last)
		swap(slave, first, last)
		first++
		last--
		if last <= first {
			return
		}
	}
}

func swap(arr []int, index1, index2 int) {
	loc := arr[index1]
	arr[index1] = arr[index2]
	arr[index2] = loc
}

func binarySearchGreater(arr []int, num int) int {
	pos1 := 0
	pos2 := len(arr) - 1
	for {
		if (pos2-pos1)%2 == 1 {
			if arr[pos2] > num {
				return pos2
			} else {
				pos2--
				continue
			}
		}
		if (pos2 - pos1) == 0 {
			return pos1
		}
		if arr[(pos2+pos1)/2] > num {
			pos1 = (pos2 + pos1) / 2
		} else {
			pos2 = (pos2 + pos1) / 2
		}
	}
}

func findFactorial(num int) (resp int) {
	if num <= 0 {
		return 0
	}
	resp = 1

	for i := 1; i <= num; i++ {
		resp *= i
	}
	return resp
}
