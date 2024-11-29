package main

// Dijkstra’s algorithm
// https://ru.wikipedia.org/wiki/Алгоритм_Дейкстры
func minimumObstacles(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	if len(grid[0]) == 1 || len(grid) == 1 {
		return counter(grid)
	}

	totalMinVal := countNegativeCaseDistance(0, 0, len(grid[0])-1, len(grid))
	dataHolder := &DataHolder{
		grid:          grid,
		totalMinVal:   &totalMinVal,
		firstPriority: &Deque{},
		lastPriority:  &Deque{},
		distanceMp:    make(distanceMap, len(grid[0])*len(grid)),
	}
	dataHolder.distanceMp.set(0, 0, grid[0][0])
	dataHolder.firstPriority.PushFront(0, 0)
	dataHolder.process()

	val := dataHolder.distanceMp.getShortest(len(grid[0])-1, len(grid)-1)
	if val == -1 {
		return totalMinVal
	}
	return val
}

type DataHolder struct {
	grid          [][]int
	firstPriority *Deque
	lastPriority  *Deque
	distanceMp    distanceMap
	totalMinVal   *int
}

func (d *DataHolder) getMaxX() int {
	return len(d.grid[0]) - 1
}

func (d *DataHolder) getMaxY() int {
	return len(d.grid) - 1
}

func (d *DataHolder) process() {
	for {
		nextPos1, nextPos2, valid := d.firstPriority.GetFront()
		if !valid {
			nextPos1, nextPos2, valid = d.lastPriority.GetFront()
			if !valid {
				return
			}
		}
		d.processOneItem(nextPos1, nextPos2)
		if *d.totalMinVal == 0 {
			return
		}
	}
}

func (d *DataHolder) processOneItem(pos1, pos2 int) {
	topShortestPath := d.distanceMp.getShortest(pos1, pos2)

	if topShortestPath >= *d.totalMinVal {
		return
	}

	maxX := d.getMaxX()
	maxY := d.getMaxY()

	minNegative := countNegativeCaseDistance(pos1, pos2, maxX, maxY) + topShortestPath

	if minNegative < *d.totalMinVal {
		*d.totalMinVal = minNegative
	}

	if childPos1, childPos2, ok := d.GetLeft(pos1, pos2); ok {
		if ok := d.processParentChildRelation(pos1, pos2, childPos1, childPos2); ok {
			if d.grid[childPos2][childPos1] == 0 {
				d.firstPriority.PushTail(childPos1, childPos2)
			} else {
				d.lastPriority.PushTail(childPos1, childPos2)
			}
		}
	}

	if childPos1, childPos2, ok := d.GetUp(pos1, pos2); ok {
		if ok := d.processParentChildRelation(pos1, pos2, childPos1, childPos2); ok {
			if d.grid[childPos2][childPos1] == 0 {
				d.firstPriority.PushTail(childPos1, childPos2)
			} else {
				d.lastPriority.PushTail(childPos1, childPos2)
			}
		}
	}

	if childPos1, childPos2, ok := d.GetDown(pos1, pos2); ok {
		if ok := d.processParentChildRelation(pos1, pos2, childPos1, childPos2); ok {
			if d.grid[childPos2][childPos1] == 0 {
				d.firstPriority.PushTail(childPos1, childPos2)
			} else {
				d.lastPriority.PushTail(childPos1, childPos2)
			}
		}
	}
	if childPos1, childPos2, ok := d.GetRight(pos1, pos2); ok {
		if ok := d.processParentChildRelation(pos1, pos2, childPos1, childPos2); ok {
			if d.grid[childPos2][childPos1] == 0 {
				d.firstPriority.PushTail(childPos1, childPos2)
			} else {
				d.lastPriority.PushTail(childPos1, childPos2)
			}
		}
	}
}

func (d *DataHolder) processParentChildRelation(topPos1, topPos2, childPos1, childPos2 int) bool {
	nextTopDistance := d.grid[childPos2][childPos1]
	childShortestPath := d.distanceMp.getShortest(childPos1, childPos2)
	topShortestPath := d.distanceMp.getShortest(topPos1, topPos2)

	if childShortestPath == -1 {
		childShortestPath = topShortestPath + nextTopDistance
		d.distanceMp.set(childPos1, childPos2, childShortestPath)
		if childShortestPath >= *d.totalMinVal {
			return false
		}
		return true
	}

	if childShortestPath > topShortestPath+nextTopDistance {
		childShortestPath = topShortestPath + nextTopDistance
		d.distanceMp.set(childPos1, childPos2, childShortestPath)
		if childShortestPath >= *d.totalMinVal {
			return false
		}
		return true
	}
	return false
}

func countNegativeCaseDistance(pos1, pos2, maxPos1, maxPos2 int) int {
	return (maxPos2 - pos2) + (maxPos1 - pos1)
}

func (d *DataHolder) GetRight(pos1, pos2 int) (int, int, bool) {
	if pos1 < len(d.grid[0])-1 {
		return pos1 + 1, pos2, true
	}
	return 0, 0, false
}

func (d *DataHolder) GetLeft(pos1, pos2 int) (int, int, bool) {
	if pos1 > 0 {
		return pos1 - 1, pos2, true
	}
	return 0, 0, false
}

func (d *DataHolder) GetDown(pos1, pos2 int) (int, int, bool) {
	if pos2 < len(d.grid)-1 {
		return pos1, pos2 + 1, true
	}
	return 0, 0, false
}

func (d *DataHolder) GetUp(pos1, pos2 int) (int, int, bool) {
	if pos2 > 0 {
		return pos1, pos2 - 1, true
	}
	return 0, 0, false
}

type Deque struct {
	Head *Node
	Tail *Node
}

func (d *Deque) GetFront() (int, int, bool) {
	out := d.Head
	if d.Head != nil {
		d.Head = d.Head.Next
		if d.Head == nil {
			d.Tail = nil
		}
	}
	if out != nil {
		return out.Position.Pos1, out.Position.Pos2, true
	}
	return 0, 0, false
}

func (d *Deque) PushFront(pos1, pos2 int) {
	if d.Head == nil && d.Tail == nil {
		d.Head = &Node{
			Position: Position{
				Pos1: pos1,
				Pos2: pos2,
			},
		}
		d.Tail = d.Head
		return
	}
	d.Head = &Node{
		Position: Position{
			Pos1: pos1,
			Pos2: pos2,
		},
		Next: d.Head,
	}
}

func (d *Deque) PushTail(pos1, pos2 int) {
	if d.Head == nil && d.Tail == nil {
		d.Head = &Node{
			Position: Position{
				Pos1: pos1,
				Pos2: pos2,
			},
		}
		d.Tail = d.Head
		return
	}
	d.Tail.Next = &Node{
		Position: Position{
			Pos1: pos1,
			Pos2: pos2,
		},
	}
	d.Tail = d.Tail.Next
}

type Node struct {
	Position Position
	Next     *Node
}

func counter(grid [][]int) (out int) {
	if len(grid) == 1 {
		for i := range grid[0] {
			if grid[0][i] == 1 {
				out++
			}
		}
		return out
	}

	pos := 0

	for {
		if pos >= len(grid) {
			break
		}
		if grid[pos][0] == 1 {
			out++
		}
		pos++
	}
	return out
}

type Position struct {
	Pos1, Pos2 int
}

type distanceMap map[Position]int

func (m *distanceMap) getShortest(pos1, pos2 int) int {
	val, ok := (*m)[Position{
		Pos1: pos1,
		Pos2: pos2,
	}]
	if ok {
		return val
	}
	return -1
}

func (m *distanceMap) set(pos1, pos2 int, shortestDistance int) {
	(*m)[Position{
		Pos1: pos1,
		Pos2: pos2,
	}] = shortestDistance
}
