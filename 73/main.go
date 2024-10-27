package main

func setZeroes(matrix [][]int) {
	rowZeroNum := -1 // using rowZero as storage

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				if rowZeroNum == -1 {
					rowZeroNum = i
					break
				}

				for m := range matrix[i] {
					if matrix[i][m] == 0 {
						matrix[rowZeroNum][m] = 0
					} else {
						matrix[i][m] = 0
					}
				}
				break
			}
		}
	}

	if rowZeroNum == -1 {
		return
	}

	for i := range matrix[rowZeroNum] {
		if matrix[rowZeroNum][i] == 0 {
			for j := 0; j < len(matrix); j++ {
				if j == rowZeroNum {
					continue
				}
				matrix[j][i] = 0
			}
		}
	}
	for i := range matrix[rowZeroNum] {
		matrix[rowZeroNum][i] = 0
	}
}
