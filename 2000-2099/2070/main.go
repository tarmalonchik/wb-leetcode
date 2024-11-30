package main

import (
	"sort"
)

func maximumBeauty(items [][]int, queries []int) []int {
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})

	mp := make(mapper, len(queries))
	for i := range queries {
		mp.add(queries[i], i)
	}

	sort.Slice(queries, func(i, j int) bool {
		return queries[i] < queries[j]
	})

	maxBeauty := items[0][1]
	for i := range items {
		if items[i][1] < maxBeauty {
			items[i][1] = maxBeauty
		} else {
			maxBeauty = items[i][1]
		}
	}

	out := make([]int, len(queries))

	beautyIdx := 0
	queriesIdx := 0
	maxBeauty = 0
	recalculateBeauty := true
	prevMaxBeauty := 0
	for {
		if recalculateBeauty {
			prevMaxBeauty = maxBeauty
			beautyIdx, maxBeauty = findMaxBeauty(items, beautyIdx)
		}
		recalculateBeauty = false

		if queriesIdx == len(queries) {
			break
		}

		if items[beautyIdx][0] > queries[queriesIdx] {
			out[mp.get(queries[queriesIdx])] = prevMaxBeauty
			queriesIdx++
			continue
		} else if items[beautyIdx][0] == queries[queriesIdx] {
			out[mp.get(queries[queriesIdx])] = maxBeauty
			queriesIdx++
			continue
		} else {
			if beautyIdx == len(items)-1 {
				out[mp.get(queries[queriesIdx])] = maxBeauty
				queriesIdx++
				continue
			} else {
				recalculateBeauty = true
				beautyIdx++
			}
		}
	}
	return out
}

type mapper map[int][]int

func (m *mapper) add(number, pos int) {
	val, ok := (*m)[number]
	if !ok {
		(*m)[number] = []int{pos}
	} else {
		val = append(val, pos)
		(*m)[number] = val
	}
}

func (m *mapper) get(number int) int {
	val, ok := (*m)[number]
	if !ok {
		return -1
	}
	if len(val) == 0 {
		return -1
	}
	out := val[len(val)-1]
	val = val[:len(val)-1]
	(*m)[number] = val
	return out
}

func findMaxBeauty(items [][]int, idx int) (outIdx, maxBeauty int) {
	outIdx = idx
	maxBeauty = items[idx][1]

	for {
		if outIdx+1 < len(items) && items[outIdx+1][0] == items[outIdx][0] {
			if items[outIdx+1][1] > maxBeauty {
				maxBeauty = items[outIdx+1][1]
			}
			outIdx++
			continue
		}
		break
	}
	return outIdx, maxBeauty
}
