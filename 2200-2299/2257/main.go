package main

const (
	guard      = uint8(2)
	wall       = uint8(1)
	occupiedX  = uint8(3)
	occupiedY  = uint8(4)
	occupiedXY = uint8(5)
)

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	matrix := make([][]uint8, m)
	for i := range matrix {
		matrix[i] = make([]uint8, n)
	}

	for i := range guards {
		matrix[guards[i][0]][guards[i][1]] = guard
	}
	for i := range walls {
		matrix[walls[i][0]][walls[i][1]] = wall
	}

	if m == 1 {
		newGuards := [][]int{}
		prevWasGuard := false
		for i := range matrix[0] {
			if matrix[0][i] == guard {
				if prevWasGuard {
					continue
				} else {
					newGuards = append(newGuards, []int{0, i})
					prevWasGuard = true
				}
			} else if matrix[0][i] == wall {
				prevWasGuard = false
			}
		}
		guards = newGuards
	}

	if n == 1 {
		newGuards := [][]int{}
		prevWasGuard := false
		for i := 0; i < m; i++ {
			if matrix[i][0] == guard {
				if prevWasGuard {
					continue
				} else {
					newGuards = append(newGuards, []int{0, i})
					prevWasGuard = true
				}
			} else if matrix[i][0] == wall {
				prevWasGuard = false
			}
		}
	}

	for i := range guards {
		markOutSingleGuard(matrix, guards[i])
	}

	var result int

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				result++
			}
		}
	}

	return result
}

func markOutSingleGuard(matrix [][]uint8, guardPos []int) {
	// up
	for i := guardPos[0] - 1; i >= 0; i-- {
		done := false

		switch matrix[i][guardPos[1]] {
		case 0:
			matrix[i][guardPos[1]] = occupiedY
		case wall:
			done = true
			break
		case guard:
			continue
		case occupiedY:
			done = true
			break
		case occupiedX:
			matrix[i][guardPos[1]] = occupiedXY
		case occupiedXY:
			done = true
			break
		}

		if done {
			break
		}
	}

	// down
	for i := guardPos[0] + 1; i < len(matrix); i++ {
		done := false

		switch matrix[i][guardPos[1]] {
		case 0:
			matrix[i][guardPos[1]] = occupiedY
		case wall:
			done = true
			break
		case guard:
			continue
		case occupiedY:
			done = true
			break
		case occupiedX:
			matrix[i][guardPos[1]] = occupiedXY
		case occupiedXY:
			done = true
			break
		}

		if done {
			break
		}
	}

	// left
	for i := guardPos[1] - 1; i >= 0; i-- {
		done := false

		switch matrix[guardPos[0]][i] {
		case 0:
			matrix[guardPos[0]][i] = occupiedX
		case wall:
			done = true
			break
		case guard:
			continue
		case occupiedX:
			done = true
			break
		case occupiedY:
			matrix[guardPos[0]][i] = occupiedXY
		case occupiedXY:
			done = true
			break
		}

		if done {
			break
		}
	}

	// right
	for i := guardPos[1] + 1; i < len(matrix[0]); i++ {
		done := false

		switch matrix[guardPos[0]][i] {
		case 0:
			matrix[guardPos[0]][i] = occupiedX
		case wall:
			done = true
			break
		case guard:
			continue
		case occupiedX:
			done = true
			break
		case occupiedY:
			matrix[guardPos[0]][i] = occupiedXY
		case occupiedXY:
			done = true
			break
		}

		if done {
			break
		}
	}
}
