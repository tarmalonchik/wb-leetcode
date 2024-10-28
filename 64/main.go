package main

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	baseLine := 0
	baseColumn := 0
	for {
		currentLine := baseLine
		currentColumn := baseColumn

		for {
			if currentLine < 0 || currentColumn > len(grid[0])-1 {
				break
			}
			if currentLine > 0 && currentColumn > 0 {
				if grid[currentLine][currentColumn-1] > grid[currentLine-1][currentColumn] {
					grid[currentLine][currentColumn] += grid[currentLine-1][currentColumn]
				} else {
					grid[currentLine][currentColumn] += grid[currentLine][currentColumn-1]
				}
				// up and left
			} else if currentLine > 0 {
				grid[currentLine][currentColumn] += grid[currentLine-1][currentColumn]
				// up
			} else if currentColumn > 0 {
				grid[currentLine][currentColumn] += grid[currentLine][currentColumn-1]
				// left
			}
			currentLine--
			currentColumn++
		}

		if baseLine == len(grid)-1 && baseColumn == len(grid[0])-1 {
			return grid[baseLine][baseColumn]
		}

		if baseLine == len(grid)-1 {
			baseColumn++
		} else {
			baseLine++
		}
	}
}
