package main

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	out := make([][]string, 0)
	mp := make(map[string]int, len(strs))
	for i := range strs {
		sorted := sortString(strs[i])
		val, ok := mp[sorted]
		if !ok {
			out = append(out, []string{})
			out[len(out)-1] = append(out[len(out)-1], strs[i])
			mp[sorted] = len(out) - 1
			continue
		}
		out[val] = append(out[val], strs[i])
	}
	return out
}

func sortString(in string) (out string) {
	slice := make([]byte, len(in))
	for i := range in {
		slice[i] = in[i]
	}
	sort.Slice(slice, func(a int, b int) bool { return slice[a] < slice[b] })
	return string(slice)
}
