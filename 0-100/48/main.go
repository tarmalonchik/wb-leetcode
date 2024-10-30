package main

func rotate(matrix [][]int) {
	if len(matrix) == 0 || len(matrix) == 1 {
		return
	}
	for j := range matrix {
		for i := 0 + j; i < len(matrix)-j; i++ {
			if i == len(matrix)-1-j {
				break
			}
			first := matrix[j][i]
			matrix[j][i] = matrix[len(matrix)-1-i][j]
			matrix[len(matrix)-1-i][j] = matrix[len(matrix)-1-j][len(matrix)-1-i]
			matrix[len(matrix)-1-j][len(matrix)-1-i] = matrix[i][len(matrix)-1-j]
			matrix[i][len(matrix)-1-j] = first
		}
		if j+1 >= len(matrix)/2 {
			return
		}
	}
}
