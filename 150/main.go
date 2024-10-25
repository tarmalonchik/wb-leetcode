package main

func canCompleteCircuit(gas []int, cost []int) int {
	vals, positions := groupPositiveAndNegatives(gas, cost)
	return findTheBestPosition(nexter(vals), positions)
}

func findTheBestPosition(in nexter, positions []int) int {
	if len(in) == 0 {
		return -1
	}

	pos := 0
	maxVal := 0
	for {
		maxVal += in[pos]
		prevPos := pos
		_, pos = in.Next(pos)
		if prevPos >= pos {
			if maxVal < 0 {
				return -1
			}
			break
		}
	}

	startPos := 0
	pos = startPos
	maxVal = 0
	for {
		maxVal += in[pos]
		if maxVal < 0 {
			_, startPos = in.Next(startPos)
			pos = startPos
			maxVal = 0
		} else {
			_, pos = in.Next(pos)
			if pos == startPos {
				if maxVal >= 0 {
					return positions[pos]
				} else {
					return -1
				}
			}
		}

	}
}

func groupPositiveAndNegatives(gas []int, cost []int) (vals grouped, positions []int) {
	if len(gas) == 0 {
		return nil, nil
	}

	vals = make(grouped, 0, len(gas))
	positions = make([]int, 0, len(gas))

	for i := range gas {
		hadSwitch := vals.addItem(gas[i] - cost[i])
		if hadSwitch {
			if (gas[i] - cost[i]) >= 0 {
				positions = append(positions, i)
			} else {
				positions = append(positions, -1)
			}
		}
	}

	if len(vals) == 1 {
		return vals, positions
	}

	if vals[0] <= 0 && vals[len(vals)-1] <= 0 {
		vals[0] += vals[len(vals)-1]
		return vals[:len(vals)-1], positions[:len(positions)-1]
	}
	if vals[0] > 0 && vals[len(vals)-1] > 0 {
		vals[len(vals)-1] += vals[0]
		return vals[1:], positions[1:]
	}

	return vals, positions
}

type nexter []int

func (n *nexter) Next(currentPos int) (val, pos int) {
	if len(*n) == 0 {
		return 0, -1
	}
	if currentPos >= len(*n) {
		return 0, -1
	}
	if currentPos == len(*n)-1 {
		return (*n)[0], 0
	}
	return (*n)[currentPos+1], currentPos + 1
}

type grouped []int

func (g *grouped) addItem(num int) (hadSwitch bool) {
	if len(*g) == 0 {
		*g = append(*g, num)
		return true
	}
	if (*g)[len(*g)-1] <= 0 {
		if num <= 0 {
			(*g)[len(*g)-1] += num
			return false
		} else {
			*g = append(*g, num)
			return true
		}
	} else {
		if num <= 0 {
			*g = append(*g, num)
			return true
		} else {
			(*g)[len(*g)-1] += num
			return false
		}
	}
}
