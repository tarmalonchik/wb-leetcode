package main

func maxMatrixSum(matrix [][]int) int64 {
	negativeNumbers := 0
	maxSum := int64(0)
	minNumber := abs(matrix[0][0])
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] < 0 {
				negativeNumbers++
			}
			absNum := abs(matrix[i][j])
			maxSum += absNum
			if absNum < minNumber {
				minNumber = absNum
			}
		}
	}
	if negativeNumbers%2 == 0 {
		return maxSum
	}
	return maxSum - 2*minNumber
}

func abs(a int) int64 {
	if a < 0 {
		return int64(-a)
	}
	return int64(a)
}
