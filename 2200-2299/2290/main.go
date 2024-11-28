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

	dataHolder := &DataHolder{}
	dataHolder.fill(grid)
	dataHolder.process(dataHolder.graph)

	val := dataHolder.graph[len(grid)-1][len(grid[0])-1]
	if val.ShortestPath == -1 {
		return *val.TotalMinVal
	}
	return val.ShortestPath
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

type DataHolder struct {
	graph [][]*GraphTop
}

func (d *DataHolder) process(grid [][]*GraphTop) {
	if grid[0][0].Distance {
		d.graph[0][0].ShortestPath = 1
	} else {
		d.graph[0][0].ShortestPath = 0
	}
	d.graph[0][0].process(grid, len(grid[0])-1, len(grid)-1)
}

func countNegativeCaseDistance(pos1, pos2, maxPos1, maxPos2 int) int {
	return (maxPos2 - pos2) + (maxPos1 - pos1)
}

func (g *GraphTop) process(grid [][]*GraphTop, maxX, maxY int) {
	minNegative := countNegativeCaseDistance(g.Pos1, g.Pos2, maxX, maxY) + g.ShortestPath

	if minNegative < *g.TotalMinVal {
		*g.TotalMinVal = minNegative
	}

	if g.Next == nil {
		g.Next = &Deque{}

		left := g.GetLeft(grid, maxX, maxY)
		if left != nil {
			if left.Distance {
				g.Next.PushTail(left)
			} else {
				g.Next.PushFront(left)
			}
		}

		up := g.GetUp(grid, maxX, maxY)
		if up != nil {
			if up.Distance {
				g.Next.PushTail(up)
			} else {
				g.Next.PushFront(up)
			}
		}

		down := g.GetDown(grid, maxX, maxY)
		if down != nil {
			if down.Distance {
				g.Next.PushTail(down)
			} else {
				g.Next.PushFront(down)
			}
		}

		right := g.GetRight(grid, maxX, maxY)
		if right != nil {
			if right.Distance {
				g.Next.PushTail(right)
			} else {
				g.Next.PushFront(right)
			}
		}
	}

	for nextTop := g.Next.GetFront(); nextTop != nil; nextTop = g.Next.GetFront() {
		nextTopDistance := 0
		if grid[nextTop.Pos2][nextTop.Pos1].Distance {
			nextTopDistance = 1
		}

		if nextTop.ShortestPath == -1 {
			nextTop.ShortestPath = g.ShortestPath + nextTopDistance
			if nextTop.ShortestPath >= *g.TotalMinVal {
				continue
			}
			nextTop.process(grid, maxX, maxY)
			continue
		}

		if nextTop.ShortestPath > g.ShortestPath+nextTopDistance {
			nextTop.ShortestPath = g.ShortestPath + nextTopDistance
			if nextTop.ShortestPath >= *g.TotalMinVal {
				continue
			}
			nextTop.process(grid, maxX, maxY)
		}
	}
}

func (g *GraphTop) GetRight(grid [][]*GraphTop, maxX, maxY int) *GraphTop {
	if g.Pos1 < maxX {
		return grid[g.Pos2][g.Pos1+1]
	}
	return nil
}

func (g *GraphTop) GetLeft(grid [][]*GraphTop, maxX, maxY int) *GraphTop {
	if g.Pos1 > 0 {
		return grid[g.Pos2][g.Pos1-1]
	}
	return nil
}

func (g *GraphTop) GetDown(grid [][]*GraphTop, maxX, maxY int) *GraphTop {
	if g.Pos2 < maxY {
		return grid[g.Pos2+1][g.Pos1]
	}
	return nil
}

func (g *GraphTop) GetUp(grid [][]*GraphTop, maxX, maxY int) *GraphTop {
	if g.Pos2 > 0 {
		return grid[g.Pos2-1][g.Pos1]
	}
	return nil
}

func (d *DataHolder) fill(grid [][]int) {
	totalMinVal := countNegativeCaseDistance(0, 0, len(grid[0])-1, len(grid)-1)

	d.graph = make([][]*GraphTop, len(grid))
	d.graph[0] = make([]*GraphTop, len(grid[0]))
	for yPos := range grid {
		if yPos+1 <= len(grid)-1 {
			d.graph[yPos+1] = make([]*GraphTop, len(grid[0]))
		}

		for xPos := range grid[yPos] {
			if d.graph[yPos][xPos] == nil {
				d.graph[yPos][xPos] = &GraphTop{
					Distance:     grid[yPos][xPos] == 1,
					ShortestPath: -1,
					TotalMinVal:  &totalMinVal,
					Pos1:         xPos,
					Pos2:         yPos,
				}
			}
		}
	}
}

type GraphTop struct {
	Pos1         int
	Pos2         int
	Distance     bool
	ShortestPath int
	TotalMinVal  *int
	Next         *Deque
}

type Deque struct {
	Head       *Node
	Tail       *Node
	CurrentPos *Node
}

func (d *Deque) GetFront() *GraphTop {
	out := d.CurrentPos
	if out == nil {
		d.CurrentPos = d.Head
	} else {
		d.CurrentPos = d.CurrentPos.Next
	}
	if out != nil {
		return out.Val
	}
	return nil
}

func (d *Deque) PushFront(top *GraphTop) {
	if d.Head == nil && d.Tail == nil {
		d.Head = &Node{
			Val: top,
		}
		d.CurrentPos = d.Head
		d.Tail = d.Head
		return
	}
	d.Head = &Node{
		Val:  top,
		Next: d.Head,
	}
	d.CurrentPos = d.Head
}

func (d *Deque) PushTail(top *GraphTop) {
	if d.Head == nil && d.Tail == nil {
		d.Head = &Node{
			Val: top,
		}
		d.CurrentPos = d.Head
		d.Tail = d.Head
		return
	}
	d.Tail.Next = &Node{
		Val: top,
	}
	d.Tail = d.Tail.Next
}

type Node struct {
	Val  *GraphTop
	Next *Node
}
