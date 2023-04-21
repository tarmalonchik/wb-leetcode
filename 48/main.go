package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}
	printMx(matrix)
	rotate(matrix)
	printMx(matrix)
}

func rotate(matrix [][]int) {
	if len(matrix) == 0 || len(matrix) == 1 {
		return
	}
	for j := range matrix {
		for i := 0 + j; i < len(matrix)-j; i++ {
			if i == len(matrix)-1-j {
				break
			}
			first := matrix[0][i]
			matrix[0][i] = matrix[len(matrix)-1-i][0]
			matrix[len(matrix)-1-i][0] = matrix[len(matrix)-1][len(matrix)-1-i]
			matrix[len(matrix)-1][len(matrix)-1-i] = matrix[i][len(matrix)-1]
			matrix[i][len(matrix)-1] = first
		}
		if j+1 >= len(matrix)/2 {
			return
		}
	}
}

func printMx(in [][]int) {
	for i := range in {
		fmt.Println()
		for j := range in[i] {
			fmt.Print(in[i][j])
			if j != len(in[i])-1 {
				fmt.Print("-")
			}
		}
	}
	fmt.Println()
}
