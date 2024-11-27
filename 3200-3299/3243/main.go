package main

// Dijkstra’s algorithm
// https://ru.wikipedia.org/wiki/Алгоритм_Дейкстры
func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	out := make([]int, 0, n)
	dataHolder := &DataHolder{
		mp: make(map[int]*GraphTop, n),
	}
	dataHolder.fill(n)
	dataHolder.process(n)

	var val int
	for i := range queries {
		if val != 1 {
			dataHolder.add(n, queries[i][0], queries[i][1])
			val = dataHolder.mp[n-1].ShortestPath
		}
		out = append(out, dataHolder.mp[n-1].ShortestPath)
	}
	return out
}

type DataHolder struct {
	mp      map[int]*GraphTop
	rootTop *GraphTop
}

func (d *DataHolder) process(n int) {
	d.rootTop.ShortestPath = 0
	d.rootTop.process(n)
}

func (g *GraphTop) process(n int) {
	if g.Value == n-1 {
		if *g.TotalMinVal == -1 {
			*g.TotalMinVal = g.ShortestPath
		} else {
			if *g.TotalMinVal < g.ShortestPath {
				g.ShortestPath = *g.TotalMinVal
			}
		}

		if g.ShortestPath == 1 {
			return
		}
	}

	if *g.TotalMinVal >= 0 && g.ShortestPath >= *g.TotalMinVal {
		return
	}

	for i := range g.NextTops {
		if g.NextTops[i].ShortestPath == -1 {
			g.NextTops[i].ShortestPath = g.ShortestPath + 1
			g.NextTops[i].process(n)
			continue
		}
		if g.NextTops[i].ShortestPath > g.ShortestPath+1 {
			g.NextTops[i].ShortestPath = g.ShortestPath + 1
			g.NextTops[i].process(n)
		}
	}
}

func (d *DataHolder) fill(n int) {
	var currentTop *GraphTop
	var totalMinVal int

	totalMinVal = -1

	for i := 0; i < n; i++ {
		if currentTop == nil {
			currentTop = &GraphTop{
				ShortestPath: -1,
				Value:        i,
				TotalMinVal:  &totalMinVal,
			}
			d.rootTop = currentTop
			d.mp[i] = currentTop
			continue
		}
		currentTop.NextTops = []*GraphTop{
			{
				ShortestPath: -1,
				Value:        i,
				TotalMinVal:  &totalMinVal,
			},
		}
		currentTop = currentTop.NextTops[0]
		d.mp[i] = currentTop
	}
}

func (d *DataHolder) add(n, from, to int) {
	fromNode, ok := d.mp[from]
	if !ok {
		panic("invalid from input")
	}
	toNode, ok := d.mp[to]
	if !ok {
		panic("invalid from input")
	}

	fromNode.NextTops = append(fromNode.NextTops, toNode)
	if toNode.ShortestPath-fromNode.ShortestPath > 1 {
		toNode.ShortestPath = fromNode.ShortestPath + 1
		toNode.process(n)
	}
}

type GraphTop struct {
	Value        int
	ShortestPath int
	NextTops     []*GraphTop
	TotalMinVal  *int
}
