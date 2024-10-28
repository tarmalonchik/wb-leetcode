package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}

	baseLine := 0
	baseColumn := 0
	for {
		currentLine := baseLine
		currentColumn := baseColumn

		for {
			if currentLine == 0 && currentColumn == 0 {
				if obstacleGrid[0][0] == 1 {
					return 0
				}
				obstacleGrid[0][0] = 1
				currentLine--
				currentColumn++
				continue
			}

			if currentLine < 0 || currentColumn > len(obstacleGrid[0])-1 {
				break
			}

			if obstacleGrid[currentLine][currentColumn] == 1 {
				obstacleGrid[currentLine][currentColumn] = 0
				currentLine--
				currentColumn++
				continue
			}

			if currentLine > 0 && currentColumn > 0 {
				obstacleGrid[currentLine][currentColumn] = obstacleGrid[currentLine][currentColumn-1] + obstacleGrid[currentLine-1][currentColumn]
			} else if currentLine > 0 {
				obstacleGrid[currentLine][currentColumn] = obstacleGrid[currentLine-1][currentColumn]
				// up
			} else if currentColumn > 0 {
				obstacleGrid[currentLine][currentColumn] = obstacleGrid[currentLine][currentColumn-1]
				// left
			}
			currentLine--
			currentColumn++
		}

		if baseLine == len(obstacleGrid)-1 && baseColumn == len(obstacleGrid[0])-1 {
			return obstacleGrid[baseLine][baseColumn]
		}

		if baseLine == len(obstacleGrid)-1 {
			baseColumn++
		} else {
			baseLine++
		}
	}
}
