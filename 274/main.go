package main

import (
	"sort"
)

func hIndex(citations []int) int {
	var out int

	sort.Ints(citations)
	for i := range citations {
		if citations[i] == 0 {
			continue
		}
		if citations[i] < len(citations)-i-1 {
			continue
		} else {
			if len(citations)-i > citations[i] {
				return citations[i]
			}
			return len(citations) - i
		}
	}
	return out
}
