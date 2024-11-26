package main

func findChampion(n int, edges [][]int) int {
	if len(edges) == 0 {
		if n == 1 {
			return 0
		}
		return -1
	}

	mp := make(map[int]int, n)

	for i := 0; i < n; i++ {
		mp[i] = i
	}

	for i := range edges {
		delete(mp, edges[i][1])
	}
	if len(mp) != 1 {
		return -1
	}

	for key, _ := range mp {
		return key
	}
	return -1
}
