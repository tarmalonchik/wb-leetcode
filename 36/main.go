package main

import (
	"fmt"
)

var input = [][]byte{
	{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
}

func main() {
	fmt.Println(isValidSudoku(input))
}

func isValidSudoku(board [][]byte) bool {
	return checkRows(board) && checkColumns(board) && checkBoxes(board)
}

func checkRows(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		mp := make(map[byte]interface{}, 10)
		for j := 0; j < 9; j++ {
			if board[i][j] == 46 {
				continue
			}
			_, ok := mp[board[i][j]]
			if ok {
				return false
			} else {
				mp[board[i][j]] = nil
			}
		}
	}
	return true
}

func checkColumns(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		mp := make(map[byte]interface{}, 10)
		for j := 0; j < 9; j++ {
			if board[j][i] == 46 {
				continue
			}
			_, ok := mp[board[j][i]]
			if ok {
				return false
			} else {
				mp[board[j][i]] = nil
			}
		}
	}
	return true
}

func checkBoxes(board [][]byte) bool {
	var mp map[byte]interface{}

	for i := 0; i < 81; i++ {
		num, isNewBox := getBoxesNext(board, i)
		if isNewBox {
			mp = make(map[byte]interface{}, 9)
		}
		if num == 46 {
			continue
		}
		_, ok := mp[num]
		if ok {
			return false
		} else {
			mp[num] = nil
		}
	}
	return true
}

func getBoxesNext(board [][]byte, position int) (value byte, newBox bool) {
	boxNum := position / 9
	inBoxNum := position % 9

	var row int
	var column int

	row += 3 * (boxNum / 3)
	row += inBoxNum / 3

	column += 3 * (boxNum % 3)
	column += inBoxNum % 3

	return board[row][column], inBoxNum == 0
}
