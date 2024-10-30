package main

func minimumTotal(triangle [][]int) int {
	for i := len(triangle) - 1; i >= 0; i-- {
		if len(triangle[i]) == 1 {
			return triangle[i][0]
		}

		for j := range triangle[i] {
			if triangle[i][j] < triangle[i][j+1] {
				triangle[i-1][j] += triangle[i][j]
			} else {
				triangle[i-1][j] += triangle[i][j+1]
			}
			if j+1 == len(triangle[i])-1 {
				break
			}
		}
	}
	return 0
}
