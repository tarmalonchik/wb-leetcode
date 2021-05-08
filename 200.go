package main

type area struct {
	current [2]int
	up      *area
	right   *area
	down    *area
	left    *area
}

func calculate(mp map[[2]int]struct{}) {
	var needToServe []*area
	current := &area{}
	for key, _ := range mp {
		current.current = key
		break
	}
	needToServe = append(needToServe, current)
	for {
		if len(needToServe) == 0 {
			break
		}
		current := needToServe[len(needToServe)-1]
		needToServe = needToServe[:len(needToServe)-1]
		if current.up == nil {
			up := [2]int{current.current[0], current.current[1] + 1}
			if _, ok := mp[up]; ok {
				current.up = &area{
					current: up,
					down:    current,
				}
				needToServe = append(needToServe, current.up)
			}
		}
		if current.right == nil {
			right := [2]int{current.current[0] + 1, current.current[1]}
			if _, ok := mp[right]; ok {
				current.right = &area{
					current: right,
					left:    current,
				}
				needToServe = append(needToServe, current.right)
			}
		}
		if current.down == nil {
			down := [2]int{current.current[0], current.current[1] - 1}
			if _, ok := mp[down]; ok {
				current.down = &area{
					current: down,
					up:      current,
				}
				needToServe = append(needToServe, current.down)
			}
		}
		if current.left == nil {
			left := [2]int{current.current[0] - 1, current.current[1]}
			if _, ok := mp[left]; ok {
				current.left = &area{
					current: left,
					right:   current,
				}
				needToServe = append(needToServe, current.left)
			}
		}
		delete(mp, current.current)
	}
}

func numIslands(grid [][]byte) int {
	if grid == nil {
		return 0
	}
	mp := map[[2]int]struct{}{}
	var emptyStruct struct{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 49 {
				mp[[2]int{i, j}] = emptyStruct
			}
		}
	}
	count := 0
	for {
		if len(mp) == 0 {
			return count
		}
		calculate(mp)
		count++
	}
}
