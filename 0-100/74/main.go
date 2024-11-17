package main

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	colIdx := findInColumn(matrix, 0, len(matrix)-1, target)
	if colIdx == -1 {
		return false
	}
	return findInRow(matrix[colIdx], 0, len(matrix[0])-1, target)
}

func findInColumn(matrix [][]int, pos1, pos2, target int) (columNum int) {
	if pos1 == pos2 {
		if matrix[pos1][0] <= target && matrix[pos1][len(matrix[0])-1] >= target {
			return pos1
		}
		return -1
	}
	if pos2-pos1 == 1 {
		if matrix[pos1][0] <= target && matrix[pos1][len(matrix[0])-1] >= target {
			return pos1
		}
		if matrix[pos2][0] <= target && matrix[pos2][len(matrix[0])-1] >= target {
			return pos2
		}
		return -1
	}

	center := pos1 + (pos2-pos1)/2

	if matrix[center][0] > target {
		return findInColumn(matrix, pos1, center, target)
	}
	if matrix[center][len(matrix[0])-1] < target {
		return findInColumn(matrix, center, pos2, target)
	}
	return center
}

func findInRow(row []int, pos1, pos2, target int) bool {
	if pos1 == pos2 {
		if row[pos1] == target {
			return true
		}
		return false
	}
	if pos2-pos1 == 1 {
		if row[pos1] == target {
			return true
		}
		if row[pos2] == target {
			return true
		}
		return false
	}

	center := pos1 + (pos2-pos1)/2
	if row[center] > target {
		return findInRow(row, pos1, center, target)
	}
	if row[center] < target {
		return findInRow(row, center, pos2, target)
	}
	return row[center] == target
}
