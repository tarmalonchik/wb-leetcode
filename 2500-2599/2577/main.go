package main

// Dijkstra’s algorithm
// https://ru.wikipedia.org/wiki/Алгоритм_Дейкстры
func minimumTime(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	if grid[0][1] > 1 && grid[1][0] > 1 {
		return -1
	}

	dataHolder := &DataHolder{
		grid:            grid,
		totalMinSeconds: -1,
		highPriority:    &Deque{},
		lowPriority:     &Deque{},
		distanceMp:      make(distanceMap, len(grid[0])*len(grid)),
	}

	dataHolder.distanceMp.set(0, 0, 0)
	dataHolder.highPriority.PushFront(0, 0)
	dataHolder.process()

	val := dataHolder.distanceMp.getShortest(len(grid[0])-1, len(grid)-1)
	if val == -1 {
		panic("invalid data")
	}
	return val
}

type DataHolder struct {
	grid            [][]int
	highPriority    *Deque
	lowPriority     *Deque
	distanceMp      distanceMap
	totalMinSeconds int
}

func (d *DataHolder) getMaxX() int {
	return len(d.grid[0]) - 1
}

func (d *DataHolder) getMaxY() int {
	return len(d.grid) - 1
}

func (d *DataHolder) process() {
	for {
		nextPos1, nextPos2, valid := d.highPriority.GetFront()
		if !valid {
			nextPos1, nextPos2, valid = d.lowPriority.GetFront()
			if !valid {
				return
			}
		}
		d.processOneItem(nextPos1, nextPos2)
	}
}

func (d *DataHolder) processOneItem(pos1, pos2 int) {
	topBestTime := d.distanceMp.getShortest(pos1, pos2)

	if pos1 == len(d.grid[0])-1 && pos2 == len(d.grid)-1 {
		if d.totalMinSeconds == -1 {
			d.totalMinSeconds = topBestTime
		} else if d.totalMinSeconds > topBestTime {
			d.totalMinSeconds = topBestTime
		}
	}

	if d.totalMinSeconds >= 0 && topBestTime >= d.totalMinSeconds {
		return
	}

	if childPos1, childPos2, ok := d.GetLeft(pos1, pos2); ok {
		if ok := d.processParentChildRelation(pos1, pos2, childPos1, childPos2); ok {
			if d.grid[childPos2][childPos1] <= topBestTime+1 {
				d.highPriority.PushTail(childPos1, childPos2)
			} else {
				d.lowPriority.PushTail(childPos1, childPos2)
			}
		}
	}

	if childPos1, childPos2, ok := d.GetUp(pos1, pos2); ok {
		if ok := d.processParentChildRelation(pos1, pos2, childPos1, childPos2); ok {
			if d.grid[childPos2][childPos1] <= topBestTime+1 {
				d.highPriority.PushTail(childPos1, childPos2)
			} else {
				d.lowPriority.PushTail(childPos1, childPos2)
			}
		}
	}

	if childPos1, childPos2, ok := d.GetDown(pos1, pos2); ok {
		if ok := d.processParentChildRelation(pos1, pos2, childPos1, childPos2); ok {
			if d.grid[childPos2][childPos1] <= topBestTime+1 {
				d.highPriority.PushFront(childPos1, childPos2)
			} else {
				d.lowPriority.PushFront(childPos1, childPos2)
			}
		}
	}

	if childPos1, childPos2, ok := d.GetRight(pos1, pos2); ok {
		if ok := d.processParentChildRelation(pos1, pos2, childPos1, childPos2); ok {
			if d.grid[childPos2][childPos1] <= topBestTime+1 {
				d.highPriority.PushFront(childPos1, childPos2)
			} else {
				d.lowPriority.PushFront(childPos1, childPos2)
			}
		}
	}
}

func (d *DataHolder) processParentChildRelation(topPos1, topPos2, childPos1, childPos2 int) bool {
	childTime := d.grid[childPos2][childPos1]
	childBestTime := d.distanceMp.getShortest(childPos1, childPos2)
	topTime := d.distanceMp.getShortest(topPos1, topPos2)
	newChildTime := 0

	if topTime+1 < childTime {
		newChildTime = topTime
		offset := childTime - (topTime + 1)
		if offset%2 != 0 {
			offset += 1
		}
		newChildTime += offset
		newChildTime++
	} else {
		newChildTime = topTime + 1
	}

	if childBestTime == -1 {
		d.distanceMp.set(childPos1, childPos2, newChildTime)
		if d.totalMinSeconds >= 0 && newChildTime >= d.totalMinSeconds {
			return false
		}
		return true
	}

	if childBestTime > newChildTime {
		d.distanceMp.set(childPos1, childPos2, newChildTime)
		if d.totalMinSeconds >= 0 && newChildTime >= d.totalMinSeconds {
			return false
		}
		return true
	}

	return false
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

func (m *distanceMap) set(pos1, pos2 int, time int) {
	(*m)[Position{
		Pos1: pos1,
		Pos2: pos2,
	}] = time
}
