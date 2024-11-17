package main

func generateMatrix(n int) [][]int {
	out := make([][]int, n)
	for i := range out {
		out[i] = make([]int, n)
	}
	spiralOrder(out)
	return out
}

const right = 0
const down = 1
const left = 2
const up = 3

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	count := len(matrix) * len(matrix[0])
	line := 0
	col := 0
	offset := 0
	direction := right
	if len(matrix[0]) == 1 {
		direction = down
	}
	out := make([]int, count)

	idx := 1

	for i := 0; i < count; i++ {
		matrix[col][line] = idx
		line, col, direction, offset = getNext(matrix, line, col, direction, offset)
		idx++
	}
	return out
}

func getNext(matrix [][]int, linePos, colPos, dir, offset int) (lineOut, colOut, dirOut, offsetOut int) {
	if dir == right && linePos < len(matrix[0])-1-offset {
		if linePos == len(matrix[0])-2-offset {
			return linePos + 1, colPos, down, offset
		}
		return linePos + 1, colPos, right, offset
	}
	if dir == down && colPos < len(matrix)-1-offset {
		if colPos == len(matrix)-2-offset {
			return linePos, colPos + 1, left, offset
		}
		return linePos, colPos + 1, down, offset
	}
	if dir == left && linePos > offset {
		if linePos == offset+1 {
			return linePos - 1, colPos, up, offset
		}
		return linePos - 1, colPos, left, offset
	}
	if dir == up && colPos > offset+1 {
		if colPos == offset+2 {
			return linePos, colPos - 1, right, offset + 1
		}
		return linePos, colPos - 1, up, offset
	}

	return 0, 0, 0, 0
}
