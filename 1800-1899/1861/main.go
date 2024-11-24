package main

const stone = '#'
const obstacle = '*'
const empty = '.'

func rotateTheBox(box [][]byte) [][]byte {
	if len(box) == 0 || len(box[0]) == 0 {
		return nil
	}

	out := make([][]byte, len(box[0]))
	for i := range out {
		out[i] = make([]byte, len(box))
	}

	for col := range out[0] {
		stonePos := []int{len(out) - 1, col}
		for row := len(out) - 1; row >= 0; row-- {
			pos1Origin, pos2Origin := rotateAdapter(len(out[0]), row, col)
			if box[pos1Origin][pos2Origin] == stone {
				out[row][col] = empty
				out[stonePos[0]][stonePos[1]] = stone
				stonePos[0]--
			} else if box[pos1Origin][pos2Origin] == obstacle {
				out[row][col] = obstacle
				stonePos[0] = row - 1
			} else {
				out[row][col] = empty
			}
		}
	}

	return out
}

func rotateAdapter(nRotated, pos1Rotated, pos2Rotated int) (pos1Original, pos2Original int) {
	return nRotated - 1 - pos2Rotated, pos1Rotated
}
